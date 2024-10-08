all: generate tidy fmt test

help: ## Display this help section
ifeq ($(OSTYPE), linux-gnu)
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z\$$/]+.*:.*?##\s/ {printf "\033[36m%-38s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
else
	@gawk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z\$$/]+.*:.*?##\s/ {printf "\033[36m%-38s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
endif

init: ## install required deps for building
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install go.uber.org/mock/mockgen@latest

generate: ## Generate mocks for testing
	go generate ./...

tidy: ## run go mod tidy to update go.mod
	go mod tidy

lint: tidy fmt generate ## Run Golangci-lint
	golangci-lint run --allow-parallel-runners --timeout '3m'

fmt: ## run go fmt against code
	go fmt ./...

test: tidy fmt generate ## run go test against code
	go test -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm -f coverage.out

cov: tidy fmt generate ## run go test against code and generate coverage report
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm -f coverage.out
