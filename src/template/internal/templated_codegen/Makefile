MAKEFLAGS = -s

TEMPLATES = $(wildcard *.template)
GENERATED = $(TEMPLATES:.template=.go)

all: $(GENERATED)

%.go: %.template
	echo Compiling $<
	go run ./../../bin/compile-template -o $@ $<

clean:
	rm -f $(GENERATED)
