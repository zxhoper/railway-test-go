build:
	go build -o myapp

run:
	./myapp


SHELL := /bin/bash
# CURL_PREAMBLE := curl -v --compressed
CURL_PREAMBLE := curl 

X := $(shell ./random.sh)
Y := $(shell ./random.sh)



test-a-index:
	@echo "Get '/':"
	${CURL_PREAMBLE} localhost:1323/
	@echo ""

test-b-get-user:
	@echo "GET a user using :  '/users/:ID':"
	${CURL_PREAMBLE} localhost:1323/users/Joe${X}
	@echo ""


test-c-get-query-string:
	@echo "GET a user using :  '/show?team=AAAAAAA&member=ABCD':"
	${CURL_PREAMBLE} localhost:1323/show?"team=x-${X}&member=y-${Y}"
	@echo ""
