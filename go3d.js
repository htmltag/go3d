const go = new Go();
WebAssembly.instantiateStreaming(fetch("go3d.wasm"), go.importObject).then(
  (result) => {
    go.run(result.instance);
  },
);
