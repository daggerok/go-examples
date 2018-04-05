#!/bin/bash
export GOPATH="$GOPATH:$(pwd)"
go install golang-methods
./bin/golang-methods
