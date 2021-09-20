#!/bin/bash

for dir in $SRC_DIRS; do \
    echo "copy generated $dir" && \
    cp -r ./codegen/$dir/generated/*.go $dir; \
done
