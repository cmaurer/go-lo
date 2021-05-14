
.EXPORT_ALL_VARIABLES:

SRC_DIRS := slice

all: generate copy-generated test

dirs:
	@echo $(SRC_DIRS)

generate:
	@go run ./codegen

copy-generated:
	@for dir in $(SRC_DIRS); do \
		cp -r ./codegen/$$dir/generated/*.go $$dir; \
	done

test:
	@for dir in $(SRC_DIRS); do \
		echo $$dir && cd $$dir && go test . -coverpkg=./... -covermode=atomic -coverprofile=./coverage.txt -v -count 1; \
	done

build:



.PHONY: generate copy-generated test build