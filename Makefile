all: build

run: sort.go
	go run sort.go 

build:
	go build -ldflags -H=windowsgui -o ../installer/sortall.exe sort.go