package main

import (
	_ "embed"
	"github.com/kataras/golog"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"github.com/whimthen/mui"
	"github.com/whimthen/mui/appkit"
	"github.com/whimthen/mui/field"
	"github.com/whimthen/mui/views"
)

//go:embed Splitview.app/Contents/Resources/9911696_14.jpg
var image []byte

func main() {
	app := mui.NewAppWithDidLaunch(func(notification objc.Object) {
		winRect := core.Rect(0, 0, 800, 500)
		width := winRect.Size.Width / 2
		height := winRect.Size.Height
		_ = width
		_ = height

		leftView := cocoa.NSView_Init(core.Rect(0, 0, 320, height))
		leftView.SetWantsLayer(true)

		imageView := views.NewImageViewWithBytes(image)
		imageView.SetImageAlignment(views.NSImageAlignCenter)
		imageView.SetImageFrameStyle(views.NSImageFrameGroove)
		imageView.SetImageScaling(views.NSImageScaleProportionallyDown)
		imageView.SetAnimates(true)
		golog.Errorf("Image: %#v", imageView)
		leftView.Send("addSubview:", &imageView)

		tfRect := core.Rect(10, 200, 300, 21)
		textField := field.NewNSTextField(tfRect)
		textField.Set("placeholderString:", core.String("PlaceholderString"))
		textField.Set("drawsBackground:", true)
		textField.Set("translatesAutoresizingMaskIntoConstraints:", false)
		leftView.Send("addSubview:", &textField)

		constraint := appkit.NewLayoutConstraintWithAttr(textField,
			appkit.NSLayoutAttributeRight,
			appkit.NSLayoutRelationEqual,
			leftView,
			appkit.NSLayoutAttributeRightMargin,
			1.0, 20.0,
		)
		golog.Errorf("%#v", constraint)
		constraint.Set("active:", true)
		//leftView.Send("addConstraint:", constraint)

		leftView.AddSubviewPositionedRelativeTo(imageView, 1, textField)
		golog.Errorf("leftView: %#v", leftView)
		leftViewController := views.NewViewController(leftView)
		leftItem := views.NewSplitviewItemWithContentList(leftViewController)

		rightView := views.NewVisualEffectView()
		rightView.SetMaterial(views.NSVisualEffectMaterialSidebar)
		rightItem := views.NewSplitviewItemWithContentList(views.NewViewController(rightView))

		splitview := views.NewSplitviewWithDelegate(views.DefaultSplitviewDelegate{})
		splitview.SetVertical(true)
		splitview.SetDividerStyle(views.NSSplitViewDividerStyleThin)
		controller := views.NewSplitviewControllerWithDelegate(views.DefaultSplitviewControllerDelegate{})
		controller.SetSplitview(splitview)
		controller.AddSplitviewItem(leftItem)
		controller.AddSplitviewItem(rightItem)

		win := cocoa.NSWindow_WithContentViewController(controller)
		win.SetFrameDisplay(winRect, true)
		win.SetStyleMask(mui.NSWindowStyleMaskTitled |
			mui.NSWindowStyleMaskClosable |
			mui.NSWindowStyleMaskResizable |
			mui.NSWindowStyleMaskMiniaturizable |
			mui.NSWindowStyleMaskFullSizeContentView)
		win.SetTitlebarAppearsTransparent(true)
		win.SetMovableByWindowBackground(true)
		win.SetTitle("Welcome")
		win.SetTitleVisibility(cocoa.NSWindowTitleHidden)
		win.MakeKeyAndOrderFront(controller)
		win.SetCollectionBehavior(cocoa.NSWindowCollectionBehaviorDefault)
		win.Center()
	})

	menu := app.MainMenu()
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
	_ = menu
	menu.AddItem(item)
	app.Set("windowsMenu:", mainMenu)

	app.Run()
}