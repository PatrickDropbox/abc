all:
	python gen_op_codes.py
	python gen_access_flags.py
	scalac *.scala

test-rewrite: all
	mkdir -p test_rewrite
	cp *.class test_rewrite
	scala RewriteClassFile test_rewrite/*.class

clean:
	rm *.class
	rm -rf test_rewrite
