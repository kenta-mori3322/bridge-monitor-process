VERSION := 0.0.1
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=Human \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=humans \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'
DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
PROJECT_NAME = $(shell git remote get-url origin | xargs basename -s .git)

all: install

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/processord

build: go.sum clean
		go build -mod=mod $(BUILD_FLAGS) -o build/processord ./cmd/processord

build-linux:
	GOOS=linux GOARCH=amd64 $(MAKE) build

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify

# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify

# devnet

# Create log files
log-files:
	sudo mkdir -p /var/log/processord && sudo touch /var/log/processord/processord.log && sudo touch /var/log/processord/processord_error.log
