PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

# TODO: Update the ldflags with the app, client & server names
#TODO  at the moment calling version.name = cosmos as its the default when no version to enable debugging with keyring-test
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=cosmos \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=dald \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=dalcli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

all: install

install: go.sum
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/dald
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/dalcli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

# Uncomment when you have some tests
# test:
# 	@go test -mod=readonly $(PACKAGES)

# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify
