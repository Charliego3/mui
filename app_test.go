package mui

import (
	"github.com/kataras/golog"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"github.com/whimthen/mui/field"
	"github.com/whimthen/mui/toolbar"
	"testing"
)

func TestApp(t *testing.T) {
	app := NewApp()
	t.Log(app.Get("mainWindow"))
}

func TestLaunching(t *testing.T) {
	t.Logf("%#v", objc.Get("NSDictionaryOfVariableBindings"))

	app := cocoa.NSApp_WithDidLaunch(func(notification objc.Object) {
		win := cocoa.NSWindow_Init(
			core.Rect(0, 0, 600, 665),
			cocoa.NSClosableWindowMask|
				cocoa.NSResizableWindowMask|
				cocoa.NSMiniaturizableWindowMask|
				//cocoa.NSTexturedBackgroundWindowMask|
				cocoa.NSFullSizeContentViewWindowMask|
				cocoa.NSTitledWindowMask,
			cocoa.NSBackingStoreRetained,
			false,
		)
		win.SetHasShadow(true)
		//win.SetTitlebarAppearsTransparent(true)

		tfRect := core.Rect(10, 200, 200, 21)
		textField := field.NewNSTextField(tfRect)
		textField.Set("placeholderString:", core.String("PlaceholderString"))
		textField.Set("drawsBackground:", true)
		t.Logf("TextField: %#v", textField)

		sfRect := core.Rect(10, 100, 200, 21)
		searchField := field.NewSearchField(sfRect)

		view := cocoa.NSView_Init(win.Frame())
		view.SetWantsLayer(true)
		view.Send("addSubview:", &textField)
		view.Send("addSubview:", &searchField)

		delegate := NewTestToolbarDelegate()
		bar := toolbar.NewToolBar(&delegate)
		t.Logf("Toolbar: %v", bar)
		bar.SetAutosavesConfiguration(true)
		bar.SetDisplayMode(toolbar.NSToolbarDisplayModeIconAndLabel)
		//bar.SetAllowsUserCustomization(true)
		bar.SetSizeMode(toolbar.NSToolbarSizeModeRegular)

		item := toolbar.NewItem("AppToolbarItem")
		item.SetLabel("ToolBar1")
		item.SetToolTip("This is tooltip")
		item.SetImage("system")
		item.SetMinSize(core.Size(25, 25))
		item.SetMaxSize(core.Size(100, 100))
		item.SetTarget(bar)
		item.SetTag(1)

		bar.InsertItem(item.Identifier(), 0)
		win.Set("showsToolbarButton:", true)
		win.Set("toolbar:", bar)

		win.SetTitle("NSTextField")
		win.SetContentView(view)
		win.Send("setMinSize:", core.Size(300, 300))
		win.SetIgnoresMouseEvents(false)
		win.SetMovableByWindowBackground(false)
		win.SetLevel(0)
		win.SetTitleVisibility(cocoa.NSWindowTitleHidden)
		win.MakeKeyAndOrderFront(view)
		win.SetCollectionBehavior(cocoa.NSWindowCollectionBehaviorCanJoinAllSpaces)
		win.Center()
	})

	mainMenu := cocoa.NSMenu_New()
	mainMenu.SetTitle("MainMenu")
	rootMenu := cocoa.NSMenu_New()
	rootMenu.SetTitle("root_menu")

	obj, sel := core.Callback(func(object objc.Object) {})
	item := cocoa.NSMenuItem_New()
	item.SetTitle("menu1")
	item.SetAction(sel)
	item.SetTarget(obj)
	//mainMenu.Send("setSubmenu:forItem:", rootMenu, item)
	mainMenu.AddItem(item)

	t.Logf("NSApp.mainMenu: %v", mainMenu)
	app.Set("windowsMenu:", mainMenu)

	app.Retain()
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)
	app.Run()
}

// TestToolbarDelegate ============================================================================================================================
type TestToolbarDelegate struct {
	objc.Object `objc:"TestToolbarDelegate : NSObject"`
}

var testDelegateIdentifier = core.String("AppToolbarItem")

func (t *TestToolbarDelegate) AllowedItemIdentifiers(obj objc.Object) objc.Object {
	golog.Errorf("Toolbar.AllowedItemIdentifiers, Obj: %#v", obj)
	return core.NSArray_WithObjects(testDelegateIdentifier, core.String("colors"))
}

func (t *TestToolbarDelegate) DefaultItemIdentifiers(obj objc.Object) objc.Object {
	golog.Errorf("Toolbar.DefaultItemIdentifiers, Obj: %#v", obj)
	return core.NSArray_WithObjects(testDelegateIdentifier, core.String("colors"))
}

func (t *TestToolbarDelegate) GenerateItem(obj objc.Object, itemIdentifier objc.Object, willBeInsertedIntoToolbar bool) objc.Object {
	golog.Errorf("Toolbar.GenerateItem, Obj: %#v, item: %#v, willBeInserted: %v", obj, itemIdentifier.String(), willBeInsertedIntoToolbar)
	item := toolbar.NewItem(itemIdentifier.String())
	item.SetLabel("ToolBar1")
	item.SetPaletteLabel("ToolBar1")
	item.SetToolTip("This is tooltip")
	item.SetImage("left")
	//item.Set("title:", core.String("Title"))
	btn1 := cocoa.NSView{Object: objc.Get("NSButton").Alloc().Send("initWithFrame:", core.Rect(0, 0, 50, 22))}
	btn1.Set("title:", core.String("Change Title"))
	btn1.Send("setButtonType:", 1)
	btn1.Send("setButtonType:", 0)
	item.Set("view:", btn1)
	item.SetMinSize(core.Size(25, 25))
	item.SetMaxSize(core.Size(100, 100))
	item.SetTarget(obj)
	item.SetTag(1)
	return item
}

func (t *TestToolbarDelegate) SelectableItemIdentifiers(obj objc.Object) core.NSString {
	return testDelegateIdentifier
}

var testToolbarDelegate objc.Object

func NewTestToolbarDelegate() TestToolbarDelegate {
	if testToolbarDelegate == nil {
		lazyRegisterToolbarDelegate()
	}
	return TestToolbarDelegate{Object: testToolbarDelegate}
}

func lazyRegisterToolbarDelegate() {
	delegate := TestToolbarDelegate{}
	class := objc.NewClassFromStruct(delegate)
	class.AddMethod("methodSignatureForSelector:", toolbar.MethodSignatureForSelector)
	class.AddMethod("doesNotRecognizeSelector:", toolbar.DoesNotRecognizeSelector)
	class.AddMethod("toolbarAllowedItemIdentifiers:", (*TestToolbarDelegate).AllowedItemIdentifiers)
	class.AddMethod("toolbarDefaultItemIdentifiers:", (*TestToolbarDelegate).DefaultItemIdentifiers)
	class.AddMethod("toolbarSelectableItemIdentifiers:", (*TestToolbarDelegate).SelectableItemIdentifiers)
	class.AddMethod("toolbar:itemForItemIdentifier:willBeInsertedIntoToolbar:", (*TestToolbarDelegate).GenerateItem)
	objc.RegisterClass(class)
	testToolbarDelegate = class.Alloc().Init()
}