#!/usr/bin/env bash

# run tests locally with gotest if it is installed (for colorized output)
if hash gotest 2>/dev/null; then
    go_test_cmd="gotest"
else
    go_test_cmd="go test"
fi

test_dirs=("shared" "services")

for i in "${test_dirs[@]}"
do
   eval $go_test_cmd -v ./...
done
