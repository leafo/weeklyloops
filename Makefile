
.PHONY: deps

loop4: loop4.go loops/*.go
	go build loop4.go

loop3: loop3.go loops/*.go
	go build loop3.go

loop2: loop2.go loops/*.go
	go build loop2.go

loop1: loop1.go loops/*.go
	go build loop1.go

plytest: plytest.go loops/*.go loops/ply/*.go
	go build plytest.go

deps: 
	go get github.com/go-gl/gl/v4.1-core/gl
	go get github.com/go-gl/glfw/v3.1/glfw
