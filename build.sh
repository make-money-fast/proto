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

function run_build() {
  dirs=$(find . -type d -not -path '*/.*')
  for i in $dirs; do
    if [[ $i == "." ]]; then
      continue
    fi
    build_dir $i
  done
}

function on_success() {
  curl https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa \
          -H 'Content-Type: application/json' \
          -d '{"msgtype": "text","text": {"content": "构建完成"}}'
}

function fail() {
   curl https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=693axxx6-7aoc-4bc4-97a0-0ec2sifa5aaa \
            -H 'Content-Type: application/json' \
            -d '{"msgtype": "text","text": {"content": "构建失败"}}'
}

case $1 in
"build")
    run_build
  ;;
"success")
  on_success
  ;;
"fail")
  on_fail
  ;;
esac