package main

import (
	"github.com/kataras/golog"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"github.com/whimthen/mui"
	"github.com/whimthen/mui/appkit"
	"github.com/whimthen/mui/field"
	"github.com/whimthen/mui/views"
	"testing"
)

func TestNewLayoutConstraint2(t *testing.T) {
	app := mui.NewAppWithDidLaunch(func(notification objc.Object) {
		winRect := core.Rect(0, 0, 800, 500)
		width := winRect.Size.Width / 2
		height := winRect.Size.Height
		_ = width
		_ = height

		leftView := cocoa.NSView_Init(core.Rect(0, 0, 320, height))
		leftView.SetWantsLayer(true)

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
		golog.Errorf("%#v before", constraint)
		constraint.Set("active:", true)
		golog.Errorf("%#v after", constraint)
		leftView.Send("addConstraint:", constraint)
		golog.Errorf("addConstraint %#v after", constraint)

		controller := views.NewViewController(leftView)
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

	app.Run()
}
