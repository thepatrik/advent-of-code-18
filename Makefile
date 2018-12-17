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

chronalclassification:
	go test -v ./chronalclassification

.PHONY: test cronalcalibration invmgmtsys nomatterhowyousliceit marblemania chronalclassification
