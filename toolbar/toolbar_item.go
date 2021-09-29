package toolbar

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type Item struct {
	objc.Object `objc:"ToolbarItem : NSToolbarItem"`
}

var toolbarItem objc.Object

func init() {
	c := objc.NewClassFromStruct(Item{})
	c.AddMethod("methodSignatureForSelector:", MethodSignatureForSelector)
	c.AddMethod("doesNotRecognizeSelector:", DoesNotRecognizeSelector)
	objc.RegisterClass(c)
	toolbarItem = objc.Get("ToolbarItem")
}

func NewItem(identifier string) Item {
	return Item{Object: toolbarItem.Alloc().Send("initWithItemIdentifier:", core.String(identifier))}
}

func (item *Item) SetLabel(label string) {
	item.Set("label:", core.String(label))
}

func (item *Item) SetPaletteLabel(label string) {
	item.Set("paletteLabel:", core.String(label))
}

func (item *Item) SetToolTip(toolTip string) {
	item.Set("toolTip:", core.String(toolTip))
}

func (item *Item) SetImage(image string) {
	item.Set("image:", cocoa.NSImage_ImageNamed(image))
}

func (item *Item) SetMinSize(size core.NSSize) {
	item.Set("minSize:", size)
}

func (item *Item) SetMaxSize(size core.NSSize) {
	item.Set("maxSize:", size)
}

func (item *Item) SetTarget(obj objc.Object) {
	item.Set("target:", obj)
}

func (item *Item) SetTag(tag int) {
	item.Set("target:", tag)
}

func (item *Item) Identifier() string {
	return item.Get("itemIdentifier").String()
}