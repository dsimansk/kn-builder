# knb - kn builder

CLI tool to enhance plugin build experience for [Knative Client](https://github.com/knative/client).


#### Build
```bash
make build
```

#### Install 
```bash
go get -u github.com/dsimansk/knb
```

#### Usage

##### Create custom `kn` distribution

The `knb` can be used to generate enhanced `kn` source files with inlined plugins. 

Create configuration file `.kn.yaml` in a root directory of `knative/client` that should specify at least `name, module, version` coordinates of the plugin. 
The `plugin distro` will generated the required go files and add dependency to `go.mod`. Optionally module replacements can be added.


Example of `.kn.yaml`
```yaml
plugins:
  - name: kn-plugin-source-kafka
    module: knative.dev/kn-plugin-source-kafka
    version: v0.19.0
    replace:
      - module: foo.bar
        version: v0.0.1
```


Execute command
```bash
knb plugin distro
```


Build `kn`
```bash
./hack/build.sh
```

##### Enable plugin inline feature 

The `knb` can be used to generate required go files to inline any `kn` plugin.

```bash
knb plugin init --name kn-source-kafka --cmd source,kafka --description "Some plugin"
```


##### List of commands

Plugin level commands

```
Manage kn plugins.

Usage:
  knb plugin [command]

Available Commands:
  distro      Generate required files to build `kn` with inline plugins.
  init        Generate required resource to inline plugin.
  register    Register plugin to kn.

Flags:
  -h, --help   help for plugin

Use "knb plugin [command] --help" for more information about a command.

```

