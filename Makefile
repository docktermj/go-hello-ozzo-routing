# Makefile that builds a "go" program.

# PROGRAM_NAME is the name of the GIT repository.
PROGRAM_NAME := $(shell basename `git rev-parse --show-toplevel`)
BUILD_VERSION := $(shell git describe --always --tags --abbrev=0 --dirty)
BUILD_TAG := $(shell git describe --always --tags --abbrev=0)
BUILD_ITERATION := $(shell git log $(BUILD_TAG)..HEAD --oneline | wc -l)


# The first "make" target runs as default.
.PHONY: default
default: build-local

# -----------------------------------------------------------------------------
# Local development
# -----------------------------------------------------------------------------

.PHONY: build-local
build-local:
	go install github.com/docktermj/$(PROGRAM_NAME)


.PHONY: test
test:
	go test github.com/docktermj/$(PROGRAM_NAME)/...


# -----------------------------------------------------------------------------
# Utility targets
# -----------------------------------------------------------------------------

.PHONY: dependencies
dependencies:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure


.PHONY: help
help:
	@echo "Build $(PROGRAM_NAME) version $(BUILD_VERSION)-$(BUILD_ITERATION)".
	@echo "All targets:"
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs
