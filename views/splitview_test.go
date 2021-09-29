package views

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"github.com/whimthen/mui"
	"github.com/whimthen/mui/appkit"
	"testing"
)

func TestNewSplitView(t *testing.T) {
	view := NewSplitview(core.Rect(0, 0, 100, 100))
	t.Logf("Splitview: %#v", view)
}

func TestNewSplitviewItem(t *testing.T) {
	rect1 := core.Rect(0, 0, 100, 100)
	view1 := cocoa.NSView_Init(rect1)
	view1.SetBackgroundColor(cocoa.Color(135, 135, 84, 1))
	item1 := NewSplitviewItem()
	item1.Set("viewController:", view1)
	t.Logf("SplitviewItem1: %#v", item1)
}

func TestSplitviewController(t *testing.T) {
	var width, height float64 = 100, 100
	rect1 := core.Rect(0, 0, width, height)
	view1 := cocoa.NSView_Init(rect1)
	view1.SetBackgroundColor(cocoa.Color(135, 135, 84, 1))
	item1 := NewSplitviewItem()
	item1.Set("viewController:", NewViewController(view1))
	t.Logf("SplitviewItem1: %#v", item1)

	rect2 := core.Rect(width, 0, width, height)
	view2 := cocoa.NSView_Init(rect2)
	view2.SetBackgroundColor(cocoa.Color(80, 102, 222, 1))
	item2 := NewSplitviewItem()
	item2.Set("viewController:", NewViewController(view2))

	splitview := NewSplitview(rect1)
	splitview.Set("vertical:", true)
	t.Logf("Splitview: %#v", splitview)
	controller := NewSplitviewController()
	controller.SetSplitview(splitview)
	t.Logf("AddSplitviewItem before, SplitviewController: %#v", controller)
	controller.Send("addSplitViewItem:", item1)
	t.Log("AddSplitviewItem after")
	controller.AddSplitviewItem(item2)
}

func TestSplitview(t *testing.T) {
	//cocoa.TerminateAfterWindowsClose = false
	app := mui.NewAppWithDidLaunch(func(notification objc.Object) {
		winRect := core.Rect(0, 0, 800, 500)
		width := winRect.Size.Width / 2
		height := winRect.Size.Height
		_ = width
		_ = height
		leftView := cocoa.NSView_Init(core.Rect(0, 0, 320, height))
		leftViewController := NewViewController(leftView)
		leftItem := NewSplitviewItemWithContentList(leftViewController)

		rightView := NewVisualEffectView()
		rightView.SetMaterial(NSVisualEffectMaterialSidebar)
		rightItem := NewSplitviewItemWithContentList(NewViewController(rightView))

		splitview := NewSplitviewWithDelegate(TestSplitviewDelegate{})
		splitview.SetVertical(true)
		splitview.SetDividerStyle(NSSplitViewDividerStyleThin)
		controller := NewSplitviewControllerWithDelegate(DefaultSplitviewControllerDelegate{})
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
		//win.SetHasShadow(true)
		win.SetTitlebarAppearsTransparent(true)
		//win.SetTitle("NSTextField")
		//win.Set("minSize:", core.Size(300, 300))
		//win.Set("contentMinSize:", core.Size(300, 300))
		//win.SetIgnoresMouseEvents(false)
		win.SetMovableByWindowBackground(true)
		//win.SetLevel(0)
		win.SetTitleVisibility(cocoa.NSWindowTitleHidden)
		win.MakeKeyAndOrderFront(controller)
		win.SetCollectionBehavior(cocoa.NSWindowCollectionBehaviorDefault)
		win.Center()
	})

	app.Run()
}

type TestSplitviewDelegate struct {
	DefaultSplitviewDelegate
}
func (t TestSplitviewDelegate) SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(obj objc.Object, effectiveRect objc.Object, forDrawnRect objc.Object, dividerIndex int) objc.Object {
	return appkit.ZeroRect
}
func (t TestSplitviewDelegate) SplitViewShouldHideDividerAtIndex(obj objc.Object, dividerIndex int) bool {
	return true
}

func TestSplitview1(t *testing.T) {
	app := mui.NewAppWithDidLaunch(func(notification objc.Object) {
		winRect := core.Rect(0, 0, 800, 500)
		width := winRect.Size.Width / 2
		height := winRect.Size.Height
		_ = width
		_ = height

		// left
		leftView := cocoa.NSView_Init(core.Rect(0, 0, 320, height))
		leftWindow := NewWindowWithController(NewViewController(leftView))
		leftWindow.SetMovableByWindowBackground(false)
		leftWindow.SetIgnoresMouseEvents(false)
		leftWindowController := NewWindowControllerWithWindow(leftWindow)
		leftItem := NewSplitviewItemWithContentList(leftWindowController)

		// right
		rightView := NewVisualEffectView()
		rightView.SetMaterial(NSVisualEffectMaterialSidebar)
		//rightWindow := NewWindowWithController(NewViewController(rightView))
		//rightWindow.SetMovableByWindowBackground(false)
		//rightWindow.SetIgnoresMouseEvents(false)
		//rightWindowController := NewWindowControllerWithWindow(rightWindow)
		rightItem := NewSplitviewItemWithContentList(NewViewController(rightView))

		// split
		splitview := NewSplitview()
		splitview.SetVertical(true)
		controller := NewSplitviewControllerWithSplitview(splitview)
		controller.AddSplitviewItem(leftItem)
		controller.AddSplitviewItem(rightItem)

		win := NewWindowWithController(controller)
		win.SetFrameDisplay(winRect, false)
		win.SetStyleMask(mui.NSWindowStyleMaskTitled |
			mui.NSWindowStyleMaskClosable |
			//mui.NSWindowStyleMaskResizable |
			mui.NSWindowStyleMaskMiniaturizable |
			mui.NSWindowStyleMaskFullSizeContentView)
		//win.SetHasShadow(true)
		//win.SetTitlebarAppearsTransparent(true)
		//win.SetIgnoresMouseEvents(false)
		//win.SetMovableByWindowBackground(false)
		//win.SetLevel(0)
		//win.SetTitleVisibility(cocoa.NSWindowTitleHidden)
		//win.MakeKeyAndOrderFront(controller)
		//win.SetCollectionBehavior(cocoa.NSWindowCollectionBehaviorDefault)
		//win.Center()
	})

	app.Run()
}
