.DEFAULT_GOAL:= test

test:
	go test -v ./...

cronalcalibration:
	go test -v ./cronalcalibration

invmgmtsys:
	go test -v ./invmgmtsys

nomatterhowyousliceit:
	go test -v ./nomatterhowyousliceit

marblemania:
	go test -v ./marblemania

.PHONY: test cronalcalibration invmgmtsys nomatterhowyousliceit marblemania
