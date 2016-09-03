#!/bin/bash
env GOOS=windows GOARCH=amd64 go build -x -o ./bin/win64/wincase.exe .
env GOOS=linux   GOARCH=amd64 go build -x -o ./bin/linux64/wincase .
env GOOS=darwin  GOARCH=amd64 go build -x -o ./bin/darwin64/wincase .
