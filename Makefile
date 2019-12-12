
.DEFAULT_GOAL := all
all: wasm web

wasm:
    # To build the WASM, GOOS=js and GOARCH=wasm must be defined.
    # then the output must be .wasm (-o main.wasm)
	@ GOOS=js GOARCH=wasm go build -o main.wasm cmd/main.go

.PHONY: web
web:
	@ cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js web/js/wasm_exec.js

clean:
	@ if [ -f main.wasm ] ; then  rm main.wasm; fi;
	@ if [ -f web/js/main.wasm ] ; then  rm web/js/main.wasm; fi;
	@ if [ -f web/js/wasm_exec.js ] ; then  rm web/js/wasm_exec.js; fi;

run: all
	@ docker run --rm \
            -v $(shell pwd)/configs/Caddyfile:/etc/Caddyfile:ro \
            -v $(shell pwd)/main.wasm:/srv/js/main.wasm:ro \
            -v $(shell pwd)/web:/srv:ro \
            -p 80:80 \
            abiosoft/caddy
