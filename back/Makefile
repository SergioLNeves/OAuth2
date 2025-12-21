.PHONY: run install-air

install-air:
	@which air > /dev/null || go install github.com/air-verse/air@latest

run: install-air
	~/go/bin/air -c .air.toml
