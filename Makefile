all: build

DIR?=d:\downloads
NOWIND=-ldflags -H=windowsgui

run: sort.go
	go run sort.go $(DIR)

build:
	go build $(NOWIND) -o ../installer/sortall.exe sort.go $(DIR)