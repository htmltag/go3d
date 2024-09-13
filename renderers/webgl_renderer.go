package renderers

import (
	"syscall/js"

	"github.com/htmltag/go3d/core"
)

type WebGLRenderer struct {
	canvas  js.Value
	context js.Value
}

func NewWebGLRenderer(canvasID string) *WebGLRenderer {
	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", canvasID)
	context := canvas.Call("getContext", "webgl")

	return &WebGLRenderer{
		canvas:  canvas,
		context: context,
	}
}

func (r *WebGLRenderer) Render(scene *core.Object3D, camera *core.Object3D) {
	// Clear the canvas
	r.context.Call("clearColor", 0, 0, 0, 1)
	r.context.Call("clear", 16384) // COLOR_BUFFER_BIT

	// Render the scene
	r.renderObject(scene)
}

func (r *WebGLRenderer) renderObject(object *core.Object3D) {
	// TODO: Render the object

	// Render children
	for _, child := range object.Children {
		r.renderObject(child)
	}
}
