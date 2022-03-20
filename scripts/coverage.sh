#!/bin/bash
go test -cover -coverprofile=./scripts/coverage.out ./...

TEST_ERROR=$?

function cleanUp {
  rm -f ./scripts/coverage.out
  rm -f ./scripts/coverage_cleaned.out
}

if [ $TEST_ERROR -ne 0 ]; then
    exit $TEST_ERROR
fi

# ignore mocks and main.go
grep -v 'mock_' ./scripts/coverage.out | grep -v 'main.go' | grep -v "/models/" > ./scripts/coverage_cleaned.out

go tool cover -func=./scripts/coverage_cleaned.out
go tool cover -html=./scripts/coverage_cleaned.out -o ./scripts/coverage.html

trap cleanUp EXIT
