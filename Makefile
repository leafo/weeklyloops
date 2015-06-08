
.PHONY: deps default all

default: loop5

all: loop1 loop2 loop3 loop4 loop5

%: %.go loops/*.go loops/ply/*.go
	go build $<

deps: 
	go get github.com/go-gl/gl/v4.1-core/gl
	go get github.com/go-gl/glfw/v3.1/glfw
