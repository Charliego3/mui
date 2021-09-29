package views

import (
	"github.com/progrium/macdriver/objc"
)

type ViewController struct {
	objc.Object `objc:"ViewController : NSViewController"`
}

var viewControllerClass objc.Object

func lazyRegisterViewController() {
	class := objc.NewClassFromStruct(ViewController{})
	objc.RegisterClass(class)
	viewControllerClass = class
}

func NewViewController(view objc.Object) ViewController {
	if viewControllerClass == nil {
		lazyRegisterViewController()
	}
	controller := ViewController{viewControllerClass.Alloc().Init()}
	controller.Set("view:", view)
	return controller
}
