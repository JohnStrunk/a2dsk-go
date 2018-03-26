GOPATH = $(shell go env GOPATH)
EXECUTABLES = a2dsk-go

.PHONY: all clean test

all: test $(EXECUTABLES)

a2dsk-go:
	go build cmd/a2dsk-go/a2dsk-go.go

clean:
	rm -f $(EXECUTABLES)

test: $(GOPATH)/bin/golint
	go vet ./...
	$(GOPATH)/bin/golint -set_exit_status ./...
	go test -cover ./...

#-- Download golint if not present
$(GOPATH)/bin/golint:
	go get -u golang.org/x/lint/golint
