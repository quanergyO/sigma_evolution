rebuild: clean build


build:
	mkdir build
	cd build && cmake ../view/CMakeLists.txt && make
	cd ./build && make

clean:
	rm -rf build