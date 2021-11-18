MAKEFLAGS = -s

TEMPLATES = $(wildcard *.template)
GENERATED = $(TEMPLATES:.template=.go)

all: clean $(GENERATED)
	go run .

%.go: %.template
	echo Compiling $<
	go run ./../bin/compile-template -use-manual-code-gen=false -o $@ $<

clean:
	rm -f $(GENERATED)
