.DEFAULT_GOAL:= all

all:
	go test -v ./...

alchemicalreduction:
	go test -v ./alchemicalreduction

chronalclassification:
	go test -v ./chronalclassification

cronalcalibration:
	go test -v ./cronalcalibration

invmgmtsys:
	go test -v ./invmgmtsys

marblemania:
	go test -v ./marblemania

nomatterhowyousliceit:
	go test -v ./nomatterhowyousliceit

memorymaneuver:
	go test -v ./memorymaneuver

starsalign:
	go test -v ./starsalign

reposerecord:
	go test -v ./reposerecord

.PHONY: all cronalcalibration invmgmtsys nomatterhowyousliceit marblemania chronalclassification reposerecord alchemicalreduction memorymaneuver starsalign
