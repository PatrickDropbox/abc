MAKEFLAGS = -s

all:
	go run github.com/pattyshack/gt/tools/lr/bin/generate-lr-parser -language rust -o parser.rs ../piecemeal-parsing/grammar.lr

clean:
	rm -f parser.rs
