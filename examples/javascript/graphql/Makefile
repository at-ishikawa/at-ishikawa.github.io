start:
	node index.js

query/hello:
	@curl -XPOST -H 'Content-Type: application/json' -H 'Accept: application/json' --data-binary '{"query":"{ hello }"}' http://localhost:4000

query/error:
	@curl -v -XPOST -H 'Content-Type: application/json' -H 'Accept: application/json' --data-binary '{"query":"{ hello, error }"}' http://localhost:4000
