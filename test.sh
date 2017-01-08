#!/bin/bash

CURRPATH=$(realpath "$0")
BASEDIR=$(dirname "$CURRPATH")
PID="$(ps | grep testServer | grep -v grep | awk '{print $1}')"

kill_test() {
	if [ -n "$PID" ]; then
		kill -9 $PID
	fi
}

main() {
	kill_test
	echo ==============testing===========
	cd "$BASEDIR"
	make test
	echo ===============end==============
}

main $@
