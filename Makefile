VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT  := $(shell git log -1 --format='%H')

###############################################################################
###                                Build flags                              ###
###############################################################################

LD_FLAGS = -X github.com/bro-n-bro/osjuno/version.Version=$(VERSION) \
	-X github.com/bro-n-bro/osjuno/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(LD_FLAGS)'

###############################################################################
###                                  Build                                  ###
###############################################################################

build: go.sum
ifeq ($(OS),Windows_NT)
	@echo "building osjuno binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/osjuno.exe ./cmd
else
	@echo "building osjuno binary..."
	@go build -mod=readonly $(BUILD_FLAGS) -o build/osjuno ./cmd
endif
.PHONY: build