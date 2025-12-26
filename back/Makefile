.PHONY: run install-air install-lint lint

install-air:
	@which air > /dev/null || go install github.com/air-verse/air@latest

install-lint:
	@which golangci-lint > /dev/null || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: install-lint
	~/go/bin/golangci-lint run --config .golangci.yml

run: install-air
	~/go/bin/air -c .air.toml
