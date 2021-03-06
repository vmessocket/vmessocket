package conf

import (
	"encoding/json"

	"github.com/golang/protobuf/proto"

	"github.com/vmessocket/vmessocket/common/protocol"
	"github.com/vmessocket/vmessocket/common/serial"
	"github.com/vmessocket/vmessocket/infra/conf/cfgcommon"
	"github.com/vmessocket/vmessocket/proxy/http"
)

type HTTPAccount struct {
	Username string `json:"user"`
	Password string `json:"pass"`
}

type HTTPClientConfig struct {
	Servers []*HTTPRemoteConfig `json:"servers"`
}

type HTTPRemoteConfig struct {
	Address *cfgcommon.Address `json:"address"`
	Port    uint16             `json:"port"`
	Users   []json.RawMessage  `json:"users"`
}

type HTTPServerConfig struct {
	Timeout     uint32         `json:"timeout"`
	Accounts    []*HTTPAccount `json:"accounts"`
	Transparent bool           `json:"allowTransparent"`
}

func (v *HTTPAccount) Build() *http.Account {
	return &http.Account{
		Username: v.Username,
		Password: v.Password,
	}
}

func (v *HTTPClientConfig) Build() (proto.Message, error) {
	config := new(http.ClientConfig)
	config.Server = make([]*protocol.ServerEndpoint, len(v.Servers))
	for idx, serverConfig := range v.Servers {
		server := &protocol.ServerEndpoint{
			Address: serverConfig.Address.Build(),
			Port:    uint32(serverConfig.Port),
		}
		for _, rawUser := range serverConfig.Users {
			user := new(protocol.User)
			if err := json.Unmarshal(rawUser, user); err != nil {
				return nil, newError("failed to parse HTTP user").Base(err).AtError()
			}
			account := new(HTTPAccount)
			if err := json.Unmarshal(rawUser, account); err != nil {
				return nil, newError("failed to parse HTTP account").Base(err).AtError()
			}
			user.Account = serial.ToTypedMessage(account.Build())
			server.User = append(server.User, user)
		}
		config.Server[idx] = server
	}
	return config, nil
}

func (c *HTTPServerConfig) Build() (proto.Message, error) {
	config := &http.ServerConfig{
		Timeout:          c.Timeout,
		AllowTransparent: c.Transparent,
	}
	if len(c.Accounts) > 0 {
		config.Accounts = make(map[string]string)
		for _, account := range c.Accounts {
			config.Accounts[account.Username] = account.Password
		}
	}
	return config, nil
}
