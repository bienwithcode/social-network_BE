#!/bin/bash
go run main.go &
while true
do
	echo "Looping forever..."
	sleep 10000
done
