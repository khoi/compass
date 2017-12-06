PKGS := $(shell go list ./... | grep -v /vendor)

.PHONY: test
test:
	go test -v $(PKGS)
