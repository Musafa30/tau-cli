# tau

`tau` is a command line tool for interacting with a Taubyte-based Cloud Network. Similar to the [web console](https://console.taubyte.com), it allows you to create and manage projects, applications, resources, and more.

## Installation

```
git clone https://github.com/taubyte/tau
cd tau
go build -o ~/go/bin/tau
```

### Offline version (Optional)
Fails faster if exploring an unregistered project
```
go build -o ~/go/bin/otau -tags=localAuthClient
```

## Login

`tau login`
    - opens selection with default already selected
    - simply logs in if only default available
    - will open new if no profiles found
`tau login --new` for new
  - `--set-default` for making this new auth the default
`tau login <profile-name>` for using a specific profile


## Environment Variables:
- `TAUBYTE_PROJECT` Selected project
- `TAUBYTE_PROFILE` Selected profile
- `TAUBYTE_APPLICATION` Selected application
- `TAUBYTE_CONFIG (default: ~/tau.yaml)` Config location
- `TAUBYTE_SESSION (default: /tmp/tau-<ppid>)` Session location

## Testing

### All tests
`go test -v ./...`

### Hot reload Spider tests
`$ cd tests`

Edit [air config](tests/.air.toml#L8) `cmd = "go test -v --run <Function|Database|...> [-tags=no_rebuild]`

(Optional) Add `debug: true,` to an individual testMonkey

`$ air`


## Running Individual Prompts

`go run ./prompts/internal`


## Measuring Coverage:

### Calculate coverage for all packages
```shell
go test -v ./... -tags=localAuthClient,projectCreateable,localPatrick,cover,noPrompt -coverprofile cover.out -coverpkg ./...
```

### Display coverage for all packages
```
go tool cover -html=cover.out
go tool cover -func=cover.out
```


# Documentation
For documentation head to [tau.how](https://tau.how/docs/tau)


# Maintainers
 - Sam Stoltenberg @skelouse
 - Tafseer Khan @tafseer-khan
