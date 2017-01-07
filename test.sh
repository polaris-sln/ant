#!/bin/bash

CURRPATH=$(realpath "$0")
BASEDIR=$(dirname "$CURRPATH")
main() {
	echo ==============testing===========
	cd "$BASEDIR"
	make test
	echo ===============end==============
}

main $@
