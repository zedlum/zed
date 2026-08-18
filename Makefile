BINARY := zed
VENV := .zedpy

.PHONY: build install test venv freeze

build:
	go build -o bin/$(BINARY) ./cmd

install: build
	cp bin/$(BINARY) $(GOPATH)/bin/$(BINARY)

test:
	go test ./...

venv:
	@test -d $(VENV) || python3 -m venv $(VENV)
	@$(VENV)/bin/pip install --quiet --upgrade commitizen pre-commit

freeze: venv
	@$(VENV)/bin/pip freeze > requirements.txt
