// Instantiate Go runtime and register the WASM module
const go = new Go();
console.log("Loading main.wasm");
WebAssembly.instantiateStreaming(fetch("js/main.wasm"), go.importObject)
    .then((result) => {
        console.log("main.wasm loaded");
        go.run(result.instance);
    });

// event handler that will call a function in WASM
document.querySelector("#button-1").onclick = event => goFunction("from JS -> WASM")
