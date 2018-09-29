# netboxxx

Tool to automatically generate configs from IP prefixes


## Install

Get precompiled binaries from [netboxxx.kradalby.no](https://netboxxx.kradalby.no) or install with Go:

```
go get -u github.com/kradalby/netboxxx
```

## Usage

```
$ netboxxx
Using config file: $HOME/.netboxxx.yaml
Usage:
  netboxxx [flags]

Flags:
  -k, --apikey string     netbox API key (required)
      --config string     config file (default is $HOME/.netboxxx.yaml)
  -h, --help              help for netboxxx
  -n, --host string       netbox host (required)
  -t, --template string   Jinja template file (required)

```
