#!/bin/bash

set -e
set -x

godep go build -a github.com/jimmidyson/jolokia-proxy

docker build -t jimmidyson/jolokia-proxy:latest .
