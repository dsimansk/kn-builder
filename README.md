# kn-builder

CLI tool to enhance plugin development experience for [Knative Client](https://github.com/knative/client).


#### Build
```bash
make build
```

#### Usage

Top level commands
```bash
Manage and build kn plugins.

Usage:
  kn-builder [command]

Available Commands:
  fetch       Fetch kn repository.
  help        Help about any command
  plugin      Manage kn plugins.

Flags:
  -h, --help   help for kn-builder

Use "kn-builder [command] --help" for more information about a command.
```

Plugin level commands

```
Manage kn plugins.

Usage:
  kn-builder plugin [command]

Available Commands:
  init        Generate required resource to inline plugin.
  register    Register plugin to kn.

Flags:
  -h, --help   help for plugin

Use "kn-builder plugin [command] --help" for more information about a command.

```

