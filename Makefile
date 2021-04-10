
LDFLAGS=-L. -lplugin


all: main

main: libplugin.so
	gcc -g -ggdb -o main main.c $(LDFLAGS)

libplugin.so:
	go build --buildmode=c-shared -o libplugin.so plugin.go

clean:
	rm -f *plugin.so main

.PHONY: all clean
