MAKEFLAGS = -s

all:
	go run ./../../bin/generate-lr-parser -language rust -o parser.rs ../piecemeal-parsing/grammar.lr

clean:
	rm -f parser.rs
