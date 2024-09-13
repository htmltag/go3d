package main

import (
	"syscall/js"

	"github.com/htmltag/go3d/core"
	"github.com/htmltag/go3d/renderers"
)

func main() {
    c := make(chan struct{}, 0)

    js.Global().Set("Go3D", js.ValueOf(map[string]interface{}{
        "NewObject3D":     js.FuncOf(newObject3D),
        "NewWebGLRenderer": js.FuncOf(newWebGLRenderer),
    }))

    <-c
}

func newObject3D(this js.Value, args []js.Value) interface{} {
    return js.ValueOf(core.NewObject3D())
}

func newWebGLRenderer(this js.Value, args []js.Value) interface{} {
    if len(args) < 1 {
        return nil
    }
    canvasID := args[0].String()
    return js.ValueOf(renderers.NewWebGLRenderer(canvasID))
}
