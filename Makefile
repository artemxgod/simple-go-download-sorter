all: build


build:
	go build -ldflags -H=windowsgui -o ../installer/sortall.exe sort.go