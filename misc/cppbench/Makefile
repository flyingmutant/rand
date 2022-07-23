all: bench-c bench-g

clean:
	rm bench-c bench-g *.o

bench-c: bench-c.o nanobench-c.o
	clang++                  -flto -O3                                             -o bench-c bench-c.o nanobench-c.o

bench-g: bench-g.o nanobench-g.o
	g++                      -flto -O3 -fno-unroll-loops -fno-move-loop-invariants -o bench-g bench-g.o nanobench-g.o

bench-c.o: vendor/nanobench.h bench.cpp
	clang++ -std=c++17 -Wall -flto -O3                                             -o bench-c.o -c bench.cpp

bench-g.o: vendor/nanobench.h bench.cpp
	g++     -std=c++17 -Wall -flto -O3 -fno-unroll-loops -fno-move-loop-invariants -o bench-g.o -c bench.cpp

nanobench-c.o: vendor/nanobench.h nanobench.cpp
	clang++ -std=c++17 -Wall -flto -O3                                             -o nanobench-c.o -c nanobench.cpp

nanobench-g.o: vendor/nanobench.h nanobench.cpp
	g++     -std=c++17 -Wall -flto -O3 -fno-unroll-loops -fno-move-loop-invariants -o nanobench-g.o -c nanobench.cpp
