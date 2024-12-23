run:
	go run src/cmd/main.go

build:
	mkdir build
	go build src/cmd/main.go -o build/sigmaEvolution.out

clean:
	rm -rf build