# variable
binary=ca

# command
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## build command
	go build ./cmd/ca

.PHONY: test
test: ## run test command
	go test -cover ./...

.PHONY: run
run: build ## run sample
	./$(binary) run test2 -v hoge -sv huga
