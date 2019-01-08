.DEFAULT_GOAL:= all

all: alchemicalreduction chronalcalibration chronalclassification chronalcoordinates invmgmtsys marblemania memorymaneuver nomatterhowyousliceit reposerecord starsalign subterraneansustainability sumofparts

alchemicalreduction:
	go test -v ./alchemicalreduction

chronalcalibration:
	go test -v ./chronalcalibration

chronalclassification:
	go test -v ./chronalclassification

chronalcoordinates:
	go test -v ./chronalcoordinates

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

subterraneansustainability:
	go test -v -timeout 10h ./subterraneansustainability

sumofparts:
	go test -v ./sumofparts

test:
	go test ./...

.PHONY: all alchemicalreduction chronalcalibration chronalclassification chronalcoordinates invmgmtsys marblemania memorymaneuver nomatterhowyousliceit reposerecord starsalign subterraneansustainability sumofparts test
