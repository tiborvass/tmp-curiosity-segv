
LDFLAGS=-lm


all: main plugin.so

main:
	gcc -o main main.c $(LDFLAGS)

plugin.so:
	go build --buildmode=c-shared -o plugin.so plugin.go

clean:
	rm -f plugin.so main

.PHONY: all clean
