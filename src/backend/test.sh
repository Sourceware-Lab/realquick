#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

go test ./...
go test -tags=integration ./...
./fuzz.sh
