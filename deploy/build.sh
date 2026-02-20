#!/bin/bash
set -e

GOOS=linux GOARCH=arm64 go build -o dist/lambda/bootstrap ./cmd/lambda
