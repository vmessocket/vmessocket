package session

import (
	"context"
	"math/rand"

	"github.com/vmessocket/vmessocket/common/errors"
	"github.com/vmessocket/vmessocket/common/net"
	"github.com/vmessocket/vmessocket/common/protocol"
)

type ID uint32

func NewID() ID {
	for {
		id := ID(rand.Uint32())
		if id != 0 {
			return id
		}
	}
}

func ExportIDToError(ctx context.Context) errors.ExportOption {
	id := IDFromContext(ctx)
	return func(h *errors.ExportOptionHolder) {
		h.SessionID = uint32(id)
	}
}

type Inbound struct {
	Source  net.Destination
	Gateway net.Destination
	Tag     string
	User    *protocol.MemoryUser
}

type Outbound struct {
	Target  net.Destination
	Gateway net.Address
}

type SniffingRequest struct {
	OverrideDestinationForProtocol []string
	Enabled                        bool
	MetadataOnly                   bool
}

type Content struct {
	Protocol        string
	SniffingRequest SniffingRequest
	Attributes      map[string]string
	SkipDNSResolve  bool
}

type Sockopt struct {
	Mark uint32
}

func (c *Content) SetAttribute(name string, value string) {
	if c.Attributes == nil {
		c.Attributes = make(map[string]string)
	}
	c.Attributes[name] = value
}

func (c *Content) Attribute(name string) string {
	if c.Attributes == nil {
		return ""
	}
	return c.Attributes[name]
}
