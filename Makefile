GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOLINT=golint "-set_exit_status"

test:
	$(GOTEST) ./...
vet:
	$(GOVET) ./...
lint:
	$(GOLINT) ./...
fmt:
	go fmt ./...
.PHONY: install test fmt release