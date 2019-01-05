.DEFAULT_GOAL:= all

all: cronalcalibration invmgmtsys nomatterhowyousliceit marblemania chronalclassification reposerecord alchemicalreduction memorymaneuver starsalign sumofparts

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

reposerecord:
	go test -v ./reposerecord

starsalign:
	go test -v ./starsalign

sumofparts:
	go test -v ./sumofparts

.PHONY: all cronalcalibration invmgmtsys nomatterhowyousliceit marblemania chronalclassification reposerecord alchemicalreduction memorymaneuver starsalign sumofparts
