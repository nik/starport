package plugin

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	hashiplug "github.com/hashicorp/go-plugin"
	"github.com/tendermint/starport/starport/pkg/pluginsdk"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type StarportPlugin struct {
	Plugin   pluginsdk.Info
	client   *hashiplug.Client
	Manifest pluginsdk.PluginManifest
}

var handshakeConfig = hashiplug.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "PLUGIN",
	MagicCookieValue: "z0z0z0z",
}

var pluginMap = map[string]hashiplug.Plugin{
	"starport_plugin": &pluginsdk.StarportPlugin{},
}

var loadedPlugins []StarportPlugin

func GetLoadedPlugins() []StarportPlugin {
	return loadedPlugins
}

func LoadPlugins() {
	log.Println("Loading plugins...")
	files, err := ioutil.ReadDir(pluginFolder)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		LoadPlugin(f.Name())
	}
	log.Printf("%d plugins loaded!\n", len(loadedPlugins))
}

func LoadPlugin(f string) {
	plugin := hostPlugin(f)
	loadedPlugins = append(loadedPlugins, plugin)
}

func hostPlugin(f string) StarportPlugin {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	client := newClient(fmt.Sprintf("%s/%s/%s", pluginFolder, f, f), logger)
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	raw, err := rpcClient.Dispense("starport_plugin")
	if err != nil {
		log.Fatal(err)
	}

	plugin := raw.(pluginsdk.Info)
	return StarportPlugin{
		Plugin: plugin,
		client: client,
		Manifest: plugin.GetManifest(),
	}
}

func newClient(file string, logger hclog.Logger) *hashiplug.Client {
	return hashiplug.NewClient(&hashiplug.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Managed:         true,
		Cmd:             exec.Command(file),
		Logger:          logger,
	})
}