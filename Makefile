all: local
build_local:
	@go build -o bin/DB
run_local:
	@./bin/DB
local: build_local run_local
