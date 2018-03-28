#!/bin/bash
export GOPATH="$GOPATH:$PWD"
go install hello-world
./bin/hello-world
