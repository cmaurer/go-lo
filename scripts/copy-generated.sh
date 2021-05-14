#!/bin/bash

SRC_DIRS=slice

for dir in $SRC_DIRS; do \
    cp -r ./codegen/$dir/generated/*.go $dir; \
done
