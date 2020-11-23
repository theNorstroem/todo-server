#!/usr/bin/env bash

CLIENTPROJECT=../todo-client
HERE=`pwd`

## build a fresh client
cd $CLIENTPROJECT
npm install
npm run build:modern

## update package
cd $HERE
rm -rf public && mv $CLIENTPROJECT/dist public
pkger