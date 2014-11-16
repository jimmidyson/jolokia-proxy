#!/bin/bash

set -e
set -x

CGO_ENABLED=0 godep go build -a github.com/jimmidyson/jolokia-proxy

docker build -t jimmidyson/jolokia-proxy:latest .
