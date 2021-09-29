package toolbar

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"testing"
)

func TestNewNSToolBar(t *testing.T) {
	class := objc.NewClass("NSToolBar", "NSObject")
	class.AddMethod("init:", func(obj objc.Object) {
		obj.SendSuper("init:")
	})
	class.AddMethod("methodSignatureForSelector:", MethodSignatureForSelector)
	class.AddMethod("doesNotRecognizeSelector:", DoesNotRecognizeSelector)
	objc.RegisterClass(class)

	//class = objc.Get("NSToolBar")
	t.Logf("%#v", class)
	alloc := class.Alloc()
	t.Logf("%#v", alloc)
	object := alloc.Send("initWithIdentifier:", core.String("AppToolbar"))
	//object := alloc.Init()
	t.Logf("%#v", object)
}

func TestNewToolBar(t *testing.T) {
	toolBar := NewToolBar(nil)
	t.Logf("%#v", toolBar)
}

func TestToolBar(t *testing.T) {
	class := objc.Get("NSToolbar")
	t.Logf("%#v", class)
	alloc := class.Alloc()
	t.Logf("%#v", alloc)
	object := alloc.Send("initWithIdentifier:", core.String("AppToolbar"))
	t.Logf("%#v", object)
}