# knb - kn builder

CLI tool to enhance plugin development experience for [Knative Client](https://github.com/knative/client).


#### Build
```bash
make build
```

#### Usage

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

