package views

import (
	"github.com/progrium/macdriver/objc"
)

type SplitviewItem struct {
	objc.Object `objc:"SplitviewItem : NSSplitViewItem"`
}

var splitviewItemClass objc.Object

func lazyRegisterSplitviewItem() {
	class := objc.NewClassFromStruct(SplitviewItem{})
	objc.RegisterClass(class)
	splitviewItemClass = class
}

func getSplitviewItemClass() objc.Object {
	if splitviewItemClass == nil {
		lazyRegisterSplitviewItem()
	}
	return splitviewItemClass
}

func NewSplitviewItem() SplitviewItem {
	return SplitviewItem{Object: getSplitviewItemClass().Alloc().Init()}
}

func (item SplitviewItem) SetViewController(controller ViewController) {
	item.Set("viewController:", controller)
}

func (item SplitviewItem) SetSpringLoaded(spring bool) {
	item.Set("isSpringLoaded:", spring)
}

func (item SplitviewItem) SetCanCollapse(collapse bool) {
	item.Set("isSpringLoaded:", collapse)
}

func NewSplitviewItemWithSidebar(controller objc.Object) SplitviewItem {
	return SplitviewItem{Object: getSplitviewItemClass().Send("sidebarWithViewController:", controller)}
}

func NewSplitviewItemWithContentList(controller objc.Object) SplitviewItem {
	return SplitviewItem{Object: getSplitviewItemClass().Send("contentListWithViewController:", controller)}
}

func NewSplitviewItemWithView(controller objc.Object) SplitviewItem {
	return SplitviewItem{Object: getSplitviewItemClass().Send("splitViewItemWithViewController:", controller)}
}