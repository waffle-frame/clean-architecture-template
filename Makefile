# TODO: ---
PROJECTNAME=$(shell basename "$(PWD)")
PID=/tmp/.$(PROJECTNAME)-api-server.pid

# Documentation generation
.PHONY: docs
docs:
	# Generation
	@swagger generate spec -o $(PWD)/pkg/docs/swagger.yaml
	@swagger generate spec -m -o $(PWD)/pkg/docs/swagger.yaml

# Build
.PHONY: build
build:
	@cd ./cmd && go build -o $(PWD)/build/app

# Run project as dev
.PHONY: run
run:
	@cd ./cmd && go run main.go
