package field

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type SearchField struct {
	cocoa.NSView
}

var nsSearchField = objc.Get("NSSearchField")

func NewSearchField(frame core.NSRect) SearchField {
	return SearchField{NSView: cocoa.NSView{Object: nsSearchField.Alloc().Send("initWithFrame:", frame)}}
}
