compile:
	./node_modules/.bin/google-closure-compiler --js hello.js --js_output_file hello-compiled.js
	# python ./node_modules/closure-library/closure/bin/build/closurebuilder.py \
	# 	--root=./node_modules/closure-library/ \
	# 	--root=./myproject/ \
	# 	--namespace="myproject.start" \
	# 	--output_mode=compiled \
	# 	--compiler_jar=./node_modules/google-closure-compiler-java/compiler.jar
	./node_modules/.bin/google-closure-compiler \
		--js myproject/start.js \
		--js myproject/klass.js \
		--js_output_file myproject/start-compiled.js
