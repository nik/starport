### the flow

1. Run `plugin install` 
   1. check if _.plugins/_ already exists
   2. if yes, check which plugins already exist and clone missing ones (nice to have)
   3. CURRENT_MVP: `rm -rf .plugins` and re-clone everything
      1. build binaries with `go build -o binary_name ./plugins/{name}/main.go`
      2. will have folders that look like: <br/>
         `.plugins/{plugin_name}/pluginbinary`
2. PluginHealth
   1. Are plugins loaded?
   2. Are any plugins crashed? Can we reload them?
   3. 