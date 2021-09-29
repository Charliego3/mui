package appkit

import (
	"fmt"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

var ZeroRect = MakeRect(0, 0, 0, 0)

type NSRect struct {
	objc.Object
	Origin core.NSPoint
	Size   core.NSSize
}

func (r NSRect) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v)", r.Origin.X, r.Origin.Y, r.Size.Width, r.Size.Height)
}

func MakeRect(x, y, w, h float64) NSRect {
	return NSRect{
		Origin: core.NSPoint{
			X: x, Y: y,
		},
		Size: core.NSSize{
			Width: w, Height: h,
		},
	}
}
