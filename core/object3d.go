package core

import (
	"github.com/htmltag/go3d/math"
)

type Object3D struct {
	Position *math.Vector3
	Rotation *math.Vector3
	Scale    *math.Vector3
	Children []*Object3D
}

func NewObject3D() *Object3D {
	return &Object3D{
		Position: math.NewVector3(0, 0, 0),
		Rotation: math.NewVector3(0, 0, 0),
		Scale:    math.NewVector3(1, 1, 1),
		Children: make([]*Object3D, 0),
	}
}

func (o *Object3D) Add(child *Object3D) {
	o.Children = append(o.Children, child)
}

func (o *Object3D) Remove(child *Object3D) {
	for i, c := range o.Children {
		if c == child {
			o.Children = append(o.Children[:i], o.Children[i+1:]...)
			break
		}
	}
}

func (o *Object3D) Update() {
	// Update transformation matrix based on position, rotation, and scale

	// Update children
	for _, child := range o.Children {
		child.Update()
	}
}
