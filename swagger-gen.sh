#!/bin/bash


out_ext=$1

if [ -n "$out" ]
then
  out_ext="json"
fi

alias swagger="docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"
swagger generate spec -m -o "./swagger.$out_ext"
