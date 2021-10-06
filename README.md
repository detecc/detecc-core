# Detecc-core

Collection of core plugins🔌 and middleware for the Detecc platform. Contains the plugins for both
⚡[Detecctor](https://github.com/detecc/detecctor) and ⚡[Deteccted](https://github.com/detecc/deteccted).

## 🔌Plugin list

1. Hardware status plugin 🌡️

### Note

The plugins are still under development and are subject of change. The list of core plugins might grow with added
functionality. Each plugin will be available as an `.so` file, ready to be used. Plugins might depend on other plugins
or middleware. The source code for the plugins is located in the `src` folder.

## Middleware list

is empty as of now.

## Compiling the plugins

Compile the source code using:

```bash
go build --buildmode=plugin . 
```

### Contributions

All contributions are welcome. If you feel like your plugin should be in the Core plugins, open an issue or create a
pull request with:

1. Compiled plugin
2. Source code and tests for the plugin
3. New directory under docs with the proper documentation
