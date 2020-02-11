# Cloning this repository

```
go get github.com/verbruggenjesse/event-handler-mock-template

# Windows
cd %GOPATH%/src/github.com/verbruggenjesse/event-handler-mock-template
go mod download

# Unix / Bash
cd $GOPATH/src/github.com/verbruggenjesse/event-handler-mock-template
go mod download
```

# Implementation

This is just to get started with writing event services.

The mock implementation is not a complete example, and has some weird behaviours that will not be present in the actual implementation that will be provided.

Using multiple subscribers causes some unexpected results for logic processing (events not being processed in order as should be) but this should not be a concern since you will likely seperate features into multiple event handler services.

# Running this example

```
export EVENT_OPERATOR_ADDR=[ip address for the event operator]
export EVENT_CENTRAL_ADDR=[ip address for the event central]
```
Launch this application, using either docker or a `go run ./main.go`

Run the test client for the event store to see the results