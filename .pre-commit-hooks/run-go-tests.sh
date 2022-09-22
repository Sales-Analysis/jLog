#!/bin/bash

go test $(go list ./... | grep -v /vendor/)


if [ $? -eq 1 ]; then
  echo "tests failed"
  exit 1
fi
