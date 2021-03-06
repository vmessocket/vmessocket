package internet

import (
	"github.com/vmessocket/vmessocket/common/serial"
	"github.com/vmessocket/vmessocket/features"
)

const unknownProtocol = "unknown"

var (
	globalTransportConfigCreatorCache = make(map[string]ConfigCreator)
	globalTransportSettings           []*TransportConfig
)

type ConfigCreator func() interface{}

func ApplyGlobalTransportSettings(settings []*TransportConfig) error {
	features.PrintDeprecatedFeatureWarning("global transport settings")
	globalTransportSettings = settings
	return nil
}

func CreateTransportConfig(name string) (interface{}, error) {
	creator, ok := globalTransportConfigCreatorCache[name]
	if !ok {
		return nil, newError("unknown transport protocol: ", name)
	}
	return creator(), nil
}

func RegisterProtocolConfigCreator(name string, creator ConfigCreator) error {
	if _, found := globalTransportConfigCreatorCache[name]; found {
		return newError("protocol ", name, " is already registered").AtError()
	}
	globalTransportConfigCreatorCache[name] = creator
	return nil
}

func transportProtocolToString(protocol TransportProtocol) string {
	switch protocol {
	case TransportProtocol_TCP:
		return "tcp"
	case TransportProtocol_UDP:
		return "udp"
	case TransportProtocol_HTTP:
		return "http"
	case TransportProtocol_WebSocket:
		return "websocket"
	default:
		return unknownProtocol
	}
}

func (c *StreamConfig) GetEffectiveProtocol() string {
	if c == nil {
		return "tcp"
	}
	if len(c.ProtocolName) > 0 {
		return c.ProtocolName
	}
	return transportProtocolToString(c.Protocol)
}

func (c *StreamConfig) GetEffectiveSecuritySettings() (interface{}, error) {
	for _, settings := range c.SecuritySettings {
		if settings.Type == c.SecurityType {
			return settings.GetInstance()
		}
	}
	return serial.GetInstance(c.SecurityType)
}

func (c *StreamConfig) GetEffectiveTransportSettings() (interface{}, error) {
	protocol := c.GetEffectiveProtocol()
	return c.GetTransportSettingsFor(protocol)
}

func (c *StreamConfig) GetTransportSettingsFor(protocol string) (interface{}, error) {
	if c != nil {
		for _, settings := range c.TransportSettings {
			if settings.GetUnifiedProtocolName() == protocol {
				return settings.GetTypedSettings()
			}
		}
	}
	for _, settings := range globalTransportSettings {
		if settings.GetUnifiedProtocolName() == protocol {
			return settings.GetTypedSettings()
		}
	}
	return CreateTransportConfig(protocol)
}

func (c *TransportConfig) GetTypedSettings() (interface{}, error) {
	return c.Settings.GetInstance()
}

func (c *TransportConfig) GetUnifiedProtocolName() string {
	if len(c.ProtocolName) > 0 {
		return c.ProtocolName
	}
	return transportProtocolToString(c.Protocol)
}

func (c *StreamConfig) HasSecuritySettings() bool {
	return len(c.SecurityType) > 0
}
