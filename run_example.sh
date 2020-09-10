#!/bin/sh

# Uncomment this to see the commands as they're being run
# set -x

echo "Running main.go program"
go build && ./mockexample && rm mockexample

echo "Running tests for log package"
(cd log && go test)
echo ""

echo "Running tests for service package"
(cd service && go test)
echo ""