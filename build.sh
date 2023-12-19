#!/bin/bash

set -ex

function build_dir() {
  currentdir=$(pwd)
  cd $1
  protoc -I. --go_out=. --go-grpc_out=. ./*.proto
  mv $1/*.go .
  rm -rf $1
  cd $currentdir
}

dirs=$(find . -type d -not -path '*/.*')
for i in $dirs; do
  if [[ $i == "." ]]; then
    continue
  fi
  build_dir $i
done