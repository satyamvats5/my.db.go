##################################
#                                #
#          Makefile              #
#                                #
##################################

# Linker Flags
LDFLAGS=

# Source Directory
SRC=src

# Third-party Sources
# LIB=vendor

# Output Directory
BIN=bin

# Executable File and Path
EXE=DB

# Source Package(s)
SRCDIRS=$(dir $(wildcard $(SRC)/*/))
PACKAGES=$(sort $(patsubst $(SRC)/%/, %\\n, $(SRCDIRS)))

# Compiler
CC=go

# Compiler Commands
GOGET=install
GOBUILD=build
GOCLEAN=clean

# Go Variables
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/$(SRC)
GOBIN := $(GOBASE)/$(BIN)
GOFILES := $(wildcard $(SRC)/**/*.go)

# Build Executable from All Source Files
all: get
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(CC) $(GOBUILD) $(LDFLAGS) -o $(GOBIN)/$(EXE) $(GOFILES)

.PHONY: all

# Run executable
run: all
	@-echo "Starting Application"
	@$(BIN)/$(EXE)

.PHONY: run

# Fetch Dependency
get:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(CC) $(GOGET) $(GOFILES)

.PHONY: get

# Help Option
help: Makefile
	@-echo "Base Path:" $(GOBASE)
	@-echo "Source Path:" $(GOPATH)
	@-echo "Binary Path:" $(GOBIN)
	@-echo "Local Packages:\n" $(PACKAGES)
	@-echo "Run 'make all' to build Application"

.PHONY: help

# House-keeping
clean:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(CC) $(GOCLEAN)

.PHONY: clean
