
.PHONY: deps

loop2: loop2.go loops/*.go
	go build loop2.go

loop1: loop1.go loops/*.go
	go build loop1.go

deps: 
	go get github.com/go-gl/gl/v4.1-core/gl
	go get github.com/go-gl/glfw/v3.1/glfw
