GOPATH := $(shell go env GOPATH)
MDL := $(shell which mdl 2> /dev/null)
SHELLCHECK := $(shell which shellcheck 2> /dev/null)
YAMLLINT := $(shell which yamllint 2> /dev/null)
EXECUTABLES := a2dsk-go

.PHONY: all clean presubmit test

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

presubmit: test
	@echo -n
ifdef MDL
	@find . -regextype egrep -iregex '.*\.md' -print0 | \
		xargs -0rt -n1 $(MDL)
else
	$(warning "markdownlint (mdl) not found. Not checking markdown.")
endif
ifdef SHELLCHECK
	@find . -regextype egrep -iregex '.*\.(ba)?sh' -print0 | \
		xargs -0rt -n1 $(SHELLCHECK)
else
	$(warning "shellcheck not found. Not checking scripts.")
endif
ifdef YAMLLINT
	@find . -regextype egrep -iregex '.*\.ya?ml' -print0 | \
		xargs -0rt -n1 $(YAMLLINT)
else
	$(warning "yamllint not found. Not checking yaml files.")
endif
	@echo ALL OK.
