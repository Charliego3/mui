package views

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type Window struct {
	cocoa.NSWindow
}

func NewWindow(rect core.NSRect, windowStyle core.NSUInteger, bufferingType cocoa.NSBackingStoreType, deferCreation bool) Window {
	return Window{cocoa.NSWindow_Init(rect, windowStyle, bufferingType, deferCreation)}
}

func NewWindowWithController(controller objc.Object) Window {
	return Window{cocoa.NSWindow_WithContentViewController(controller)}
}


// WindowController ==============================================================================
type WindowController struct {
	objc.Object `objc:"WindowController : NSWindowController"`
}

var windowControllerClass objc.Object

func lazyRegisterWindowController() {
	class := objc.NewClassFromStruct(WindowController{})
	objc.RegisterClass(class)
	windowControllerClass = class
}

func NewWindowControllerWithWindow(window Window) WindowController {
	if windowControllerClass == nil {
		lazyRegisterWindowController()
	}
	return WindowController{windowControllerClass.Alloc().Send("initWithWindow:", window)}
}
