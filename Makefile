##################################
#                                #
#          Makefile              #
#                                #
##################################

# MACRO Definitions
define RECURSIVE_DIR
./...
endef

define TESTING_COMMENT
@-echo "Testing Application"
@-echo
endef

define START_COMMENT
@-echo
@-echo
@-echo "Starting Application"
@-echo
endef

# Linker Flags
LDFLAGS=-race

# Source Directory
SRC=src

# Output Directory
BIN=bin

# Executable File Name
EXE=cli

# Source Package(s)
PACKAGES=$(sort $(patsubst $(SRC)/%/, %\\n, $(dir $(wildcard $(SRC)/*/))))

# Compiler
CC=go

# Compiler Commands
GOGET=install
GOBUILD=build
GORUN=run
GOTEST=test
GOCLEAN=clean

# Go Variables
GOBASE=$(shell pwd)
GOPATH=$(GOBASE)/$(SRC)
GOBIN=$(GOBASE)/$(BIN)
GOSOURCE=$(sort $(patsubst $(SRC)/%/, %, $(dir $(wildcard $(SRC)/**/*.go))))

# Default Makefile Target
all: build
.PHONY: all

# Build Executable from All Source Files
build: get
	@GOPATH=$(GOBASE) GOBIN=$(GOBIN) $(CC) $(GOBUILD) $(LDFLAGS) $(GOSOURCE)

.PHONY: build

# Unit Test Modules
test: get
	$(TESTING_COMMENT)
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(CC) $(GOTEST) $(RECURSIVE_DIR)

.PHONY: test

# Run executable
run: build test
	$(START_COMMENT)
	@GOPATH=$(GOBASE) GOBIN=$(GOBIN) $(CC) $(GORUN) $(LDFLAGS) $(GOSOURCE)

.PHONY: run

# Fetch Dependency
get:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(CC) $(GOGET) $(RECURSIVE_DIR)

.PHONY: get

# Help Option
help: Makefile
	@-echo "Base Path:" $(GOBASE)
	@-echo "Source Path:" $(GOPATH)
	@-echo "Binary Path:" $(GOBIN)
	@-echo "Local Packages:\n" $(PACKAGES)
	@-echo "Run 'make all' to build Application"
	@-echo "Run 'make test' to test Application"

.PHONY: help

# House-keeping
clean:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(CC) $(GOCLEAN)

.PHONY: clean
