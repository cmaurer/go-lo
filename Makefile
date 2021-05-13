
.EXPORT_ALL_VARIABLES:

# SRC_DIRS := $(shell find . -depth 2 -type f -name '*.go' | sed -E 's|/[^/]+$$||' |uniq)

SRC_DIRS := slice

dirs:
	@echo $(SRC_DIRS)

generate:
	@go run ./codegen

copy-generated:
	@for dir in $(SRC_DIRS); do \
		cp -r ./codegen/$$dir/generated/*.go $$dir; \
	done

# cp -r ./codegen/$$dir/generated/*.go $$dir; \

test:
	@for dir in $(SRC_DIRS); do \
		echo $$dir && cd $$dir && go test . -v -count 1; \
	done


.PHONY: generate test