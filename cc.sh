#!/bin/bash
mkdir -p bin/linux64
mkdir -p bin/darwin64
env GOOS=darwin  GOARCH=amd64 go build -x -o ./bin/darwin64/wincase .
env GOOS=linux   GOARCH=amd64 go build -x -o ./bin/linux64/wincase .
