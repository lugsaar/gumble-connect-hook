# Gumble Connect Hook
a simple script to invoke HTTP requests when someone connects to your Mumble server.
written because we needed a translation layer to invoke code on a D1mini.

## Configuration
see example configuration file `gumble-connect-hook.yaml.example`

### Mumble Configuration
* `host` mumble host. includes port (`host:port`)
* `username` username to use in mumble. connection with certificate is not supported

### Hook Configuration
* `method` HTTP method
* `url` URL of the Hook endpoint

### Currently Supported HTTP Methods
* GET
* POST

## Used Libraries
* [layeh/gumble](https://github.com/layeh/gumble) `gumble gumbleutil opus`
* [go-yaml/yaml](https://github.com/go-yaml/yaml) `yaml.v3`
