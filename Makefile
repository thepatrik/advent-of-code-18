.DEFAULT_GOAL:= all

all: alchemicalreduction chronalclassification chronalcoordinates cronalcalibration invmgmtsys marblemania memorymaneuver nomatterhowyousliceit reposerecord starsalign sumofparts

alchemicalreduction:
	go test -v ./alchemicalreduction

chronalclassification:
	go test -v ./chronalclassification

chronalcoordinates:
	go test -v ./chronalcoordinates

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

.PHONY: all alchemicalreduction chronalclassification chronalcoordinates cronalcalibration invmgmtsys marblemania memorymaneuver nomatterhowyousliceit reposerecord starsalign sumofparts
