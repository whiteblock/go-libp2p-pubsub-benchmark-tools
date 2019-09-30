package host

import (
	"context"

	"github.com/agencyenterprise/gossip-host/pkg/cerr"
	"github.com/agencyenterprise/gossip-host/pkg/host/config"

	lhost "github.com/libp2p/go-libp2p-core/host"
	lrouter "github.com/libp2p/go-libp2p-core/routing"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const (
	// ErrUnknownTransportOption is returned when an unknown transport has been specified
	ErrUnknownTransportOption = cerr.Error("unknown transport option")
	// ErrImproperTransportOption is returned when an improper transport has been specified (e.g. 'none' with other options)
	ErrImproperTransportOption = cerr.Error("unknown improper option")
	// ErrUnknownMuxerOption is returned when an unknown muxer has been specified
	ErrUnknownMuxerOption = cerr.Error("unknown muxer option")
	// ErrImproperMuxerOption is returned when an improper muxer option format has been provided. It expects ['name', 'type']
	ErrImproperMuxerOption = cerr.Error("improper muxer option")
	// ErrUnknownSecurityOption is returned when an unknown security option has been specified
	ErrUnknownSecurityOption = cerr.Error("unknown security option")
	// ErrNilRouter is returned when a router is needed but not found
	ErrNilRouter = cerr.Error("nil router")

	pubsubTopic = "/libp2p/test/1.0.0"
)

type publisher struct {
	ps *pubsub.PubSub
}

// Host is the gossipsub host
// note: lhost.Host and lrouter.Routing are interfaces
type Host struct {
	host   lhost.Host
	ctx    context.Context
	conf   config.Config
	router lrouter.Routing
}
