package toolbar

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type SizeMode uint

const (
	NSToolbarSizeModeDefault SizeMode = iota
	NSToolbarSizeModeRegular
	NSToolbarSizeModeSmall
)

type DisplayMode uint

const (
	NSToolbarDisplayModeDefault DisplayMode = iota
	NSToolbarDisplayModeIconAndLabel
	NSToolbarDisplayModeIconOnly
	NSToolbarDisplayModeLabelOnly
)

func MethodSignatureForSelector(obj, sel objc.Object) {}

func DoesNotRecognizeSelector(obj, sel objc.Object) {}

// Toolbar ===============================================================================================================================================
type Toolbar struct {
	objc.Object `objc:"Toolbar : NSToolbar"`
}

var toolbarClass objc.Object

func lazyRegisterToolbar() {
	class := objc.NewClassFromStruct(Toolbar{})
	class.AddMethod("methodSignatureForSelector:", MethodSignatureForSelector)
	class.AddMethod("doesNotRecognizeSelector:", DoesNotRecognizeSelector)
	objc.RegisterClass(class)
	toolbarClass = objc.Get("Toolbar")
}

func NewToolBar(delegate Delegate, identifier ...string) Toolbar {
	if toolbarClass == nil {
		lazyRegisterToolbar()
	}
	var obj objc.Object
	alloc := toolbarClass.Alloc()
	if len(identifier) > 0 {
		obj = alloc.Send("initWithIdentifier:", core.String(identifier[0]))
	} else {
		obj = alloc.Init()
	}
	toolbar := Toolbar{Object: obj}
	toolbar.Set("delegate:", delegate)
	return toolbar
}

func (t *Toolbar) InsertItem(item string, index int) {
	t.Send("insertItemWithItemIdentifier:atIndex:", core.String(item), index)
}

func (t *Toolbar) SetAllowsUserCustomization(allow bool) {
	t.Set("allowsUserCustomization:", allow)
}

func (t *Toolbar) SetAutosavesConfiguration(autoSave bool) {
	t.Set("autosavesConfiguration:", autoSave)
}

func (t *Toolbar) SetDisplayMode(mode DisplayMode) {
	t.Set("displayMode:", mode)
}

func (t *Toolbar) SetSizeMode(mode SizeMode) {
	t.Set("sizeMode:", mode)
}

type Delegate interface {
	// AllowedItemIdentifiers all the item identifier
	AllowedItemIdentifiers(obj objc.Object) objc.Object

	// DefaultItemIdentifiers real to show item identifier
	DefaultItemIdentifiers(obj objc.Object) objc.Object

	// SelectableItemIdentifiers selectable item identifiers for a toolbar
	SelectableItemIdentifiers(obj objc.Object) core.NSString

	// GenerateItem 根据item标识 返回每个具体的NSToolbarItem对象实例
	// obj is Toolbar instance & flag is willBeInsertedIntoToolbar
	// https://developer.apple.com/documentation/appkit/nstoolbardelegate/1516985-toolbar
	GenerateItem(obj objc.Object, itemIdentifier objc.Object, flag bool) objc.Object
}
