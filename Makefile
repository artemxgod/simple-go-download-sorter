all: run

DIR?=~/Downloads
EXEDIR?= d:\downloads\installer\sortall.exe
NOWIND=-ldflags -H=windowsgui

run: sort.go
	go run sort.go $(DIR)

build:
	go build $(NOWIND) -o $(EXEDIR) sort.go

.PHONY: clean

clean:
	rm $(EXEDIR)