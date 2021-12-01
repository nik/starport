package pluginsdk

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

// PluginManifest is a test
type PluginManifest struct {
	ID, Name, Author, Version string
}

// Info is the interface that we're exposing as a plugin.
type Info interface {
	GetManifest() PluginManifest
}

// StarportPluginRPC is what the server is using to communicate to the plugin over RPC
type StarportPluginRPC struct {
	client *rpc.Client
}

func (i *StarportPluginRPC) GetManifest() PluginManifest {
	rep := PluginManifest{}
	err := i.client.Call("Plugin.GetManifest", new(interface{}), &rep)
	if err != nil {
		panic(err)
	}
	return rep
}

// This is the implementation of hashicorp plugin stuff
type StarportPlugin struct {
	Impl Info
}

func (s *StarportPlugin) GetManifest(args interface{}, resp *PluginManifest) error {
	*resp = s.Impl.GetManifest()
	return nil
}

func (p *StarportPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &StarportPlugin{Impl: p.Impl}, nil
}

func (StarportPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &StarportPluginRPC{client: c}, nil
}
