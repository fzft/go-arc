VERSION = (shell git decribe --tags --long --dirty --always)
#BRANCH = (shell git rev-parse --abbrev-ref HEAD)

build:
	go build -o dist/go-arc ./cmd/main.go



run: build
	./dist/go-arc



