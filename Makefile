BIN=simplehttp

all: bin/$(BIN)

.PHONY: gb
gb:
	go get -v -u github.com/constabulary/gb/... && which gb

bin/$(BIN): gb
	gb build 

clean:
	rm -r -f bin
	rm -r -f pkg
