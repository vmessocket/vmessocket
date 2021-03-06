package conf

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/vmessocket/vmessocket/app/dispatcher"
	"github.com/vmessocket/vmessocket/app/proxyman"
	"github.com/vmessocket/vmessocket/common/serial"
	"github.com/vmessocket/vmessocket/core"
	"github.com/vmessocket/vmessocket/infra/conf/cfgcommon"
)

var (
	inboundConfigLoader = NewJSONConfigLoader(ConfigCreatorCache{
		"http":  func() interface{} { return new(HTTPServerConfig) },
		"vmess": func() interface{} { return new(VMessInboundConfig) },
	}, "protocol", "settings")
	outboundConfigLoader = NewJSONConfigLoader(ConfigCreatorCache{
		"dns":     func() interface{} { return new(DNSOutboundConfig) },
		"freedom": func() interface{} { return new(FreedomConfig) },
		"http":    func() interface{} { return new(HTTPClientConfig) },
		"vmess":   func() interface{} { return new(VMessOutboundConfig) },
	}, "protocol", "settings")
	ctllog = log.New(os.Stderr, "v2ctl> ", 0)
)

type Config struct {
	Port            uint16                      `json:"port"`
	InboundConfig   *InboundDetourConfig        `json:"inbound"`
	OutboundConfig  *OutboundDetourConfig       `json:"outbound"`
	InboundDetours  []InboundDetourConfig       `json:"inboundDetour"`
	OutboundDetours []OutboundDetourConfig      `json:"outboundDetour"`
	LogConfig       *LogConfig                  `json:"log"`
	InboundConfigs  []InboundDetourConfig       `json:"inbounds"`
	OutboundConfigs []OutboundDetourConfig      `json:"outbounds"`
	Transport       *TransportConfig            `json:"transport"`
	Services        map[string]*json.RawMessage `json:"services"`
}

type InboundDetourConfig struct {
	Protocol       string                `json:"protocol"`
	PortRange      *cfgcommon.PortRange  `json:"port"`
	ListenOn       *cfgcommon.Address    `json:"listen"`
	Settings       *json.RawMessage      `json:"settings"`
	StreamSetting  *StreamConfig         `json:"streamSettings"`
	DomainOverride *cfgcommon.StringList `json:"domainOverride"`
}

type OutboundDetourConfig struct {
	Protocol      string           `json:"protocol"`
	Settings      *json.RawMessage `json:"settings"`
	StreamSetting *StreamConfig    `json:"streamSettings"`
	ProxySettings *ProxyConfig     `json:"proxySettings"`
}

func applyTransportConfig(s *StreamConfig, t *TransportConfig) {
	if s.TCPSettings == nil {
		s.TCPSettings = t.TCPConfig
	}
	if s.WSSettings == nil {
		s.WSSettings = t.WSConfig
	}
}

func toProtocolList(s []string) ([]proxyman.KnownProtocols, error) {
	kp := make([]proxyman.KnownProtocols, 0, 8)
	for _, p := range s {
		switch strings.ToLower(p) {
		case "http":
			kp = append(kp, proxyman.KnownProtocols_HTTP)
		case "https", "tls", "ssl":
			kp = append(kp, proxyman.KnownProtocols_TLS)
		default:
			return nil, newError("Unknown protocol: ", p)
		}
	}
	return kp, nil
}

func (c *Config) Build() (*core.Config, error) {
	config := &core.Config{
		App: []*serial.TypedMessage{
			serial.ToTypedMessage(&dispatcher.Config{}),
			serial.ToTypedMessage(&proxyman.InboundConfig{}),
			serial.ToTypedMessage(&proxyman.OutboundConfig{}),
		},
	}
	var logConfMsg *serial.TypedMessage
	if c.LogConfig != nil {
		logConfMsg = serial.ToTypedMessage(c.LogConfig.Build())
	} else {
		logConfMsg = serial.ToTypedMessage(DefaultLogConfig())
	}
	config.App = append([]*serial.TypedMessage{logConfMsg}, config.App...)
	if msg, err := c.BuildServices(c.Services); err != nil {
		return nil, newError("Cannot load service").Base(err)
	} else {
		config.App = append(config.App, msg...)
	}
	var inbounds []InboundDetourConfig
	if c.InboundConfig != nil {
		inbounds = append(inbounds, *c.InboundConfig)
	}
	if len(c.InboundDetours) > 0 {
		inbounds = append(inbounds, c.InboundDetours...)
	}
	if len(c.InboundConfigs) > 0 {
		inbounds = append(inbounds, c.InboundConfigs...)
	}
	if len(inbounds) > 0 && inbounds[0].PortRange == nil && c.Port > 0 {
		inbounds[0].PortRange = &cfgcommon.PortRange{
			From: uint32(c.Port),
			To:   uint32(c.Port),
		}
	}
	for _, rawInboundConfig := range inbounds {
		if c.Transport != nil {
			if rawInboundConfig.StreamSetting == nil {
				rawInboundConfig.StreamSetting = &StreamConfig{}
			}
			applyTransportConfig(rawInboundConfig.StreamSetting, c.Transport)
		}
		ic, err := rawInboundConfig.Build()
		if err != nil {
			return nil, err
		}
		config.Inbound = append(config.Inbound, ic)
	}
	var outbounds []OutboundDetourConfig
	if c.OutboundConfig != nil {
		outbounds = append(outbounds, *c.OutboundConfig)
	}
	if len(c.OutboundDetours) > 0 {
		outbounds = append(outbounds, c.OutboundDetours...)
	}
	if len(c.OutboundConfigs) > 0 {
		outbounds = append(outbounds, c.OutboundConfigs...)
	}
	for _, rawOutboundConfig := range outbounds {
		if c.Transport != nil {
			if rawOutboundConfig.StreamSetting == nil {
				rawOutboundConfig.StreamSetting = &StreamConfig{}
			}
			applyTransportConfig(rawOutboundConfig.StreamSetting, c.Transport)
		}
		oc, err := rawOutboundConfig.Build()
		if err != nil {
			return nil, err
		}
		config.Outbound = append(config.Outbound, oc)
	}
	return config, nil
}

func (c *InboundDetourConfig) Build() (*core.InboundHandlerConfig, error) {
	receiverSettings := &proxyman.ReceiverConfig{}
	if c.ListenOn == nil {
		if c.PortRange == nil {
			return nil, newError("Listen on AnyIP but no Port(s) set in InboundDetour.")
		}
		receiverSettings.PortRange = c.PortRange.Build()
	} else {
		receiverSettings.Listen = c.ListenOn.Build()
		listenDS := c.ListenOn.Family().IsDomain() && (c.ListenOn.Domain()[0] == '/' || c.ListenOn.Domain()[0] == '@')
		listenIP := c.ListenOn.Family().IsIP() || (c.ListenOn.Family().IsDomain() && c.ListenOn.Domain() == "localhost")
		switch {
		case listenIP:
			if c.PortRange == nil {
				return nil, newError("Listen on specific ip without port in InboundDetour.")
			}
			receiverSettings.PortRange = c.PortRange.Build()
		case listenDS:
			if c.PortRange != nil {
				receiverSettings.PortRange = nil
			}
		default:
			return nil, newError("unable to listen on domain address: ", c.ListenOn.Domain())
		}
	}
	if c.StreamSetting != nil {
		ss, err := c.StreamSetting.Build()
		if err != nil {
			return nil, err
		}
		receiverSettings.StreamSettings = ss
	}
	if c.DomainOverride != nil {
		kp, err := toProtocolList(*c.DomainOverride)
		if err != nil {
			return nil, newError("failed to parse inbound detour config").Base(err)
		}
		receiverSettings.DomainOverride = kp
	}
	settings := []byte("{}")
	if c.Settings != nil {
		settings = ([]byte)(*c.Settings)
	}
	rawConfig, err := inboundConfigLoader.LoadWithID(settings, c.Protocol)
	if err != nil {
		return nil, newError("failed to load inbound detour config.").Base(err)
	}
	ts, err := rawConfig.(Buildable).Build()
	if err != nil {
		return nil, err
	}
	return &core.InboundHandlerConfig{
		ReceiverSettings: serial.ToTypedMessage(receiverSettings),
		ProxySettings:    serial.ToTypedMessage(ts),
	}, nil
}

func (c *OutboundDetourConfig) Build() (*core.OutboundHandlerConfig, error) {
	senderSettings := &proxyman.SenderConfig{}
	if c.StreamSetting != nil {
		ss, err := c.StreamSetting.Build()
		if err != nil {
			return nil, err
		}
		senderSettings.StreamSettings = ss
	}
	settings := []byte("{}")
	if c.Settings != nil {
		settings = ([]byte)(*c.Settings)
	}
	rawConfig, err := outboundConfigLoader.LoadWithID(settings, c.Protocol)
	if err != nil {
		return nil, newError("failed to parse to outbound detour config.").Base(err)
	}
	ts, err := rawConfig.(Buildable).Build()
	if err != nil {
		return nil, err
	}
	return &core.OutboundHandlerConfig{
		SenderSettings: serial.ToTypedMessage(senderSettings),
		ProxySettings:  serial.ToTypedMessage(ts),
	}, nil
}

func (c *Config) Override(o *Config, fn string) {
	if o.LogConfig != nil {
		c.LogConfig = o.LogConfig
	}
	if o.Transport != nil {
		c.Transport = o.Transport
	}
	if o.InboundConfig != nil {
		c.InboundConfig = o.InboundConfig
	}
	if o.OutboundConfig != nil {
		c.OutboundConfig = o.OutboundConfig
	}
	if o.InboundDetours != nil {
		c.InboundDetours = o.InboundDetours
	}
	if o.OutboundDetours != nil {
		c.OutboundDetours = o.OutboundDetours
	}
	if len(o.InboundConfigs) > 0 {
		if len(c.InboundConfigs) > 0 && len(o.InboundConfigs) == 1 {
		} else {
			c.InboundConfigs = o.InboundConfigs
		}
	}
	if len(o.OutboundConfigs) > 0 {
		if len(c.OutboundConfigs) > 0 && len(o.OutboundConfigs) == 1 {
		} else {
			c.OutboundConfigs = o.OutboundConfigs
		}
	}
}
