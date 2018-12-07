# Database Information
## Getting Started
A simple tool for checking database information such as database version, database active client, and some database health metrics. This repository is web version from previous project, [dbcheck](https://github.com/onkiit/dbcheck).

### Grab to your local
Get this tool by execute this command
```
go get github.com/onkiit/db-monitor
```
With go, you can simply get all required packages by its tool with a single command.
First go to root project:
```
cd $(go env GOPATH)/src/github.com/onkiit/db-monitor

```
then hit: 
```
go get -v -u
```

### Configuration
Configuration file located in config/config.yaml. You can change with your own setting such as database uri, client or server host.

### Usage (LINUX)
There are 2 steps for using this tool.
#### Activate server 
```
cd $(go env GOPATH)/src/github.com/onkiit/db-monitor/cmd/monitor
go build
./monitor
```
#### Activate client 
```
cd $(go env GOPATH)/src/github.com/onkiit/db-monitor/client
npm install
npm run serve
```
