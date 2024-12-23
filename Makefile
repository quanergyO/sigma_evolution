rebuild: clean build


build:
	mkdir build
	cd build && cmake ../src/view/CMakeLists.txt && make
	cd ./build && make

clean:
	rm -rf build