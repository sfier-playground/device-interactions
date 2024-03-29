GO_BINARY_NAME=device-interaction # <- change this value to your binary name
VERSION=$(shell git describe --tags || git rev-parse --short HEAD || echo "unknown version")
LDFLAGS+= -X "github.com/sifer169966/device-interactions/pkg/flags.Version=$(VERSION)"
LDFLAGS+= -X "github.com/sifer169966/device-interactions/pkg/flags.GoVersion=$(shell go version | sed -r 's/go version go(.*)\ .*/\1/')"



# run this command once, when you work on this project for the first time
init:
	@echo "== 👩‍🌾 ci init =="
	@if [[ -z "$(shell node -v)" ]]; then \
		brew install node; \
	fi;
	@if [[ -z "$(shell pre-commit --version)" ]]; then \
		brew install pre-commit; \
	fi;
	brew install golangci-lint
	brew upgrade golangci-lint

	@echo "== pre-commit setup =="
	pre-commit install

	@echo "== install hook =="
	$(MAKE) precommit.rehooks

# installs the pre-commit hooks defined in the .pre-commit-config.yaml, specifically installs the commit-msg hook. \
The commit-msg hook is a special type of pre-commit hook that runs after you write your commit message \
but before the commit is finalized.
precommit.rehooks:
	pre-commit autoupdate
	pre-commit install --install-hooks
	pre-commit install --hook-type commit-msg

# Always turn on go module when use `go` command.
GO := GO111MODULE=on go

# Build GO application
# -v
# print packages name.
# -a
# force re-building of packages that are already up-to-date.
# -o
# -ldsflags
# update flag variable that link into the application.
.PHONY: build
build:
	$(GO) build -ldflags '$(LDFLAGS)' -a -v -o $(GO_BINARY_NAME) main.go

start:
	go run main.go serve-rest

docker.up:
	docker-compose up

.PHONY: test
test:
	go test ./... -coverprofile coverage.out

# Clean up when build the application on local directory.
.PHONY: clean
clean:
	@rm -rf $(GO_BINARY_NAME) ./vendor

mock:
	mockgen -source=./internal/core/port/repository.go -package=mocks  -destination=./mocks/repository/repository.go
	mockgen -source=./internal/core/port/service.go -package=mocks  -destination=./mocks/service/service.go
