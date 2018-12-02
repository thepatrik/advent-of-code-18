.DEFAULT_GOAL:= test

test:
	go test -v ./...

cronalcalibration:
	go test -v ./cronalcalibration

invmgmtsys:
	go test -v ./invmgmtsys

.PHONY: test cronalcalibration invmgmtsys
