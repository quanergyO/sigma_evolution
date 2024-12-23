run:
	go run cmd/main.go

build:
	mkdir build
	go build cmd/main.go -o build/sigmaEvolution.out

clean:
	rm -rf build