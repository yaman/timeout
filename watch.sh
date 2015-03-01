#!/bin/zsh

red='\033[0;32m'
NC='\033[0m' # No Color
echo "${red}running tests${NC}"
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
