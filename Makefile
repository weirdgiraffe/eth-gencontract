.PHONY: help
help: ## Print list of all commands
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-50s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## Run all of the tests
	go generate ./...
	go test ./...
