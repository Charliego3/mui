package field

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"github.com/whimthen/mui/appkit"
)

type NSTextField struct {
	cocoa.NSView `objc:"TextField : NSTextField"`
}

var textFieldObj objc.Object

func init() {
	class := objc.NewClassFromStruct(NSTextField{})
	class.AddMethod("performKeyEquivalent:", appkit.PerformKeyEquivalent)
	objc.RegisterClass(class)
	textFieldObj = objc.Get("TextField")
}

func NewNSTextField(frame core.NSRect) NSTextField {
	return NSTextField{NSView: cocoa.NSView{Object: textFieldObj.Alloc().Send("initWithFrame:", frame)}}
}

func (t NSTextField) SetBackgroundColor(color cocoa.NSColor) {
	t.Set("backgroundColor:", &color)
}

func (t NSTextField) SetBordered(isBordered bool) {
	t.Set("bordered:", isBordered)
}

func (t NSTextField) SetStringValue(val string) {
	t.Set("stringValue:", core.String(val))
}
