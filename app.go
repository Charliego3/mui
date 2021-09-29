package mui

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"runtime"
)

type Application struct {
	cocoa.NSApplication
	statusbar cocoa.NSMenu
}

type SubMenu struct {
	SubTitle string
	Action   func(object objc.Object)
}

func NewApp() Application {
	return Application{NSApplication: cocoa.NSApp()}
}

func NewAppWithDidLaunch(launch func(notification objc.Object)) Application {
	runtime.LockOSThread()
	app := cocoa.NSApp_WithDidLaunch(launch)
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)
	app.Retain()
	return Application{NSApplication: app}
}

// NewStatusBarApp return an app with status bar
// length: cocoa.NSVariableStatusItemLength
// or cocoa.NSSquareStatusItemLength
func NewStatusBarApp(title string, length float64) Application {
	cocoa.TerminateAfterWindowsClose = false
	runtime.LockOSThread()
	menu := cocoa.NSMenu_New()
	app := cocoa.NSApp_WithDidLaunch(func(notification objc.Object) {
		obj := cocoa.NSStatusBar_System().StatusItemWithLength(length)
		obj.Retain()
		obj.Button().SetTitle(title)
		obj.SetMenu(menu)
	})
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyAccessory)
	app.ActivateIgnoringOtherApps(true)
	return Application{NSApplication: app, statusbar: menu}
}

func (app Application) AddMenuItem(title string, action func(object objc.Object)) {
	obj, sel := core.Callback(action)
	item := cocoa.NSMenuItem_New()
	item.SetTitle(title)
	item.SetAction(sel)
	item.SetTarget(obj)
	app.statusbar.AddItem(item)
}

func (app Application) AddSubMenu(title string, menus ...SubMenu) {
	subMenu := cocoa.NSMenu_New()
	subItem := cocoa.NSMenuItem_New()
	subItem.SetTitle(title)
	subItem.SetSubmenu(subMenu)

	for _, menu := range menus {
		object, selector := core.Callback(menu.Action)
		t1 := cocoa.NSMenuItem_New()
		t1.SetTitle(menu.SubTitle)
		t1.SetTarget(object)
		t1.SetAction(selector)
		subMenu.AddItem(t1)
	}

	app.statusbar.AddItem(subItem)
}

func (app Application) AddMenuItemWithSelector(title string, sel objc.Selector) {
	item := cocoa.NSMenuItem_New()
	item.SetTitle(title)
	item.SetAction(sel)
	app.statusbar.AddItem(item)
}

func (app Application) AddTerminateItem(title ...string) {
	itemTitle := "Quit"
	if len(title) > 0 {
		itemTitle = title[0]
	}
	app.AddMenuItemWithSelector(itemTitle, objc.Sel("terminate:"))
}

func (app Application) AddItemSeparator() {
	app.statusbar.AddItem(cocoa.NSMenuItem_Separator())
}