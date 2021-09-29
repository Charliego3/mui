package appkit

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

func PerformKeyEquivalent(obj, event objc.Object) bool {
	modifierFlags := event.Get("modifierFlags").Int()
	character := event.Get("charactersIgnoringModifiers")
	window := obj.Get("window")
	responder := window.Get("firstResponder")

	// 点击 esc 取消焦点
	if modifierFlags&0xffff0000 == 0 && character.String() == "\x1b" {
		window.Send("makeFirstResponder:", nil)
		return true
	}

	// 判断是否 Command 组合键
	if modifierFlags&0xffff0000 != 1<<20 {
		return obj.SendSuper("performKeyEquivalent:", event).Bool()
	}

	var rtn objc.Object
	switch character.String() {
	case "a":
		rtn = obj.SendSuper("sendAction:to:", objc.Sel("selectAll:"), responder)
	case "c":
		rtn = obj.SendSuper("sendAction:to:", objc.Sel("copy:"), responder)
	case "v":
		rtn = obj.SendSuper("sendAction:to:", objc.Sel("paste:"), responder)
	case "x":
		rtn = obj.SendSuper("sendAction:to:", objc.Sel("cut:"), responder)
	case "\u007f": // Command + delete
		obj.Set("stringValue:", core.String(""))
		return true
	}
	if rtn == nil {
		return true
	}
	return rtn.Bool()
}
