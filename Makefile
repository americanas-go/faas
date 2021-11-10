#!make

GOLANGCI_STATUS :=$(shell command -v "golangci-lint")

define installGolangci
    $(if $(findstring linux,$(OS)),sudo curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin v1.38.0,brew install golangci-lint && brew upgrade golangci-lint)
endef

verify-golangci:
ifeq (,$(findstring golangci,$(GOLANGCI_STATUS)))
	$(call installGolangci)
endif

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; \
	{printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}' | \
	sort

tests: ## Run all tests
	@CGO_ENABLED=0 GOFLAGS="-count=1" go test  ./...

changelog: ## Autogenerate CHANGELOG.md
	@docker run -t -v "$(shell pwd)":/app/ orhunp/git-cliff:latest --config cliff.toml --output CHANGELOG.md

lint: verify-golangci ## Run linter and display errors
	@golangci-lint run ./...

lint-fix: verify-golangci ## Run linter and fix code when possible
	@golangci-lint run ./... --fix

.PHONY: help tests changelog lint lint-fix
