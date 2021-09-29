package views

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
	"github.com/whimthen/mui/appkit"
)

type DividerStyle uint

const (
	NSSplitViewDividerStyleThick DividerStyle = iota + 1
	NSSplitViewDividerStyleThin
	NSSplitViewDividerStylePaneSplitter
)

// SplitviewDelegate =========================================================================================================
type SplitviewDelegate interface {
	SplitViewResizeSubviewsWithOldSize(obj objc.Object, oldSize objc.Object)
	SplitViewWillResizeSubviews(notification objc.Object)
	SplitViewDidResizeSubviews(notification objc.Object)
	SplitViewCanCollapseSubview(obj objc.Object, subview objc.Object) bool
	SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex(obj objc.Object, subview objc.Object, dividerIndex int) bool
	SplitViewShouldAdjustSizeOfSubview(obj objc.Object, view objc.Object) bool
	SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(obj objc.Object, effectiveRect objc.Object, forDrawnRect objc.Object, dividerIndex int) objc.Object
	SplitViewShouldHideDividerAtIndex(obj objc.Object, dividerIndex int) bool
	SplitViewAdditionalEffectiveRectOfDividerAtIndex(obj objc.Object, dividerIndex int) objc.Object
	SplitViewConstrainMaxCoordinateOfSubviewAt(obj objc.Object, proposedMaximumPosition float32, dividerIndex int) float32
	SplitViewConstrainMinCoordinateOfSubviewAt(obj objc.Object, proposedMinimumPosition float32, dividerIndex int) float32
	SplitViewConstrainSplitPositionOfSubviewAt(obj objc.Object, proposedMinimumPosition float32, dividerIndex int) float32
}

// DefaultSplitviewDelegate =========================================================================================================
type DefaultSplitviewDelegate struct{}

func (t DefaultSplitviewDelegate) SplitViewResizeSubviewsWithOldSize(obj objc.Object, oldSize objc.Object) {}
func (t DefaultSplitviewDelegate) SplitViewWillResizeSubviews(notification objc.Object) {}
func (t DefaultSplitviewDelegate) SplitViewDidResizeSubviews(notification objc.Object)  {}
func (t DefaultSplitviewDelegate) SplitViewCanCollapseSubview(obj objc.Object, subview objc.Object) bool {
	return true
}
func (t DefaultSplitviewDelegate) SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex(obj objc.Object, subview objc.Object, dividerIndex int) bool {
	return true
}
func (t DefaultSplitviewDelegate) SplitViewShouldAdjustSizeOfSubview(obj objc.Object, view objc.Object) bool {
	return true
}
func (t DefaultSplitviewDelegate) SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(obj objc.Object, effectiveRect objc.Object, forDrawnRect objc.Object, dividerIndex int) objc.Object {
	return effectiveRect
}
func (t DefaultSplitviewDelegate) SplitViewShouldHideDividerAtIndex(obj objc.Object, dividerIndex int) bool {
	return true
}
func (t DefaultSplitviewDelegate) SplitViewAdditionalEffectiveRectOfDividerAtIndex(obj objc.Object, dividerIndex int) objc.Object {
	return nil
}
func (t DefaultSplitviewDelegate) SplitViewConstrainMaxCoordinateOfSubviewAt(obj objc.Object, proposedMaximumPosition float32, dividerIndex int) float32 {
	return proposedMaximumPosition
}
func (t DefaultSplitviewDelegate) SplitViewConstrainMinCoordinateOfSubviewAt(obj objc.Object, proposedMinimumPosition float32, dividerIndex int) float32 {
	return proposedMinimumPosition
}
func (t DefaultSplitviewDelegate) SplitViewConstrainSplitPositionOfSubviewAt(obj objc.Object, proposedMinimumPosition float32, dividerIndex int) float32 {
	return proposedMinimumPosition
}

// Splitview ===============================================================================================================
type Splitview struct {
	cocoa.NSView `objc:"Splitview : NSSplitView"`
}

var splitViewClass objc.Object

func lazyRegisterSplitview() {
	class := objc.NewClassFromStruct(Splitview{})
	objc.RegisterClass(class)
	splitViewClass = class
}

func NewSplitview(frame ...core.NSRect) Splitview {
	if splitViewClass == nil {
		lazyRegisterSplitview()
	}
	var obj objc.Object
	if len(frame) > 0 {
		obj = splitViewClass.Alloc().Send("initWithFrame:", frame[0])
	} else {
		obj = splitViewClass.Alloc().Init()
	}
	return Splitview{cocoa.NSView{Object: obj}}
}

func NewSplitviewWithDelegate(delegate SplitviewDelegate) Splitview {
	if splitViewClass == nil {
		lazyRegisterSplitview()
	}
	class := objc.Get("Splitview")
	class.AddMethod("splitView:resizeSubviewsWithOldSize:", delegate.SplitViewResizeSubviewsWithOldSize)
	class.AddMethod("splitViewWillResizeSubviews:", delegate.SplitViewWillResizeSubviews)
	class.AddMethod("splitViewDidResizeSubviews:", delegate.SplitViewDidResizeSubviews)
	class.AddMethod("splitView:canCollapseSubview:", delegate.SplitViewCanCollapseSubview)
	class.AddMethod("splitView:shouldCollapseSubview:forDoubleClickOnDividerAtIndex:", delegate.SplitViewShouldCollapseSubviewForDoubleClickOnDividerAtIndex)
	class.AddMethod("splitView:shouldAdjustSizeOfSubview:", delegate.SplitViewShouldAdjustSizeOfSubview)
	class.AddMethod("splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:", delegate.SplitViewEffectiveRectForDrawnRectOfDividerAtIndex)
	class.AddMethod("splitView:shouldHideDividerAtIndex:", delegate.SplitViewShouldHideDividerAtIndex)
	class.AddMethod("splitView:additionalEffectiveRectOfDividerAtIndex:", delegate.SplitViewAdditionalEffectiveRectOfDividerAtIndex)
	class.AddMethod("splitView:constrainMaxCoordinate:ofSubviewAt:", delegate.SplitViewConstrainMaxCoordinateOfSubviewAt)
	class.AddMethod("splitView:constrainMinCoordinate:ofSubviewAt:", delegate.SplitViewConstrainMinCoordinateOfSubviewAt)
	class.AddMethod("splitView:constrainSplitPosition:ofSubviewAt:", delegate.SplitViewConstrainSplitPositionOfSubviewAt)
	return Splitview{cocoa.NSView{Object: class.Alloc().Init()}}
}

func (s Splitview) SetVertical(vertical bool) {
	s.Set("vertical:", vertical)
}

func (s Splitview) SetDividerStyle(style DividerStyle) {
	s.Set("dividerStyle:", style)
}

func (s Splitview) SetDelegate(delegate SplitviewDelegate) {
	s.Set("delegate:", delegate)
}

func (s Splitview) ArrangedSubviews() []cocoa.NSView {
	_ = s.Get("arrangedSubviews")
	return nil
}

// SplitviewController =======================================================================================================
type SplitviewController struct {
	objc.Object `objc:"SplitviewController : NSSplitViewController"`
}

var splitViewControllerClass objc.Object

func lazyRegisterSplitviewController() {
	class := objc.NewClassFromStruct(SplitviewController{})
	objc.RegisterClass(class)
	splitViewControllerClass = class
}

func NewSplitviewController() SplitviewController {
	if splitViewControllerClass == nil {
		lazyRegisterSplitviewController()
	}
	return SplitviewController{splitViewControllerClass.Alloc().Init()}
}

func NewSplitviewControllerWithSplitview(splitview Splitview) SplitviewController {
	controller := NewSplitviewController()
	controller.SetSplitview(splitview)
	return controller
}

func NewSplitviewControllerWithDelegate(delegate SplitviewControllerDelegate) SplitviewController {
	class := objc.NewClass("GoSplitViewController", "NSSplitViewController")
	//class.AddMethod("splitView:effectiveRect:forDrawnRect:ofDividerAtIndex:", delegate.SplitViewEffectiveRectForDrawnRectOfDividerAtIndex)
	//class.AddMethod("splitView:shouldHideDividerAtIndex:", delegate.SplitViewShouldHideDividerAtIndex)
	objc.RegisterClass(class)
	return SplitviewController{class.Alloc().Init()}
}

func (s SplitviewController) SetSplitview(view Splitview) {
	s.Set("splitView:", view)
}

func (s SplitviewController) AddSplitviewItem(item SplitviewItem) {
	s.Send("addSplitViewItem:", item)
}

type SplitviewControllerDelegate interface {
	SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(obj objc.Object, effectiveRect objc.Object, forDrawnRect objc.Object, dividerIndex int) objc.Object
	SplitViewShouldHideDividerAtIndex(obj objc.Object, dividerIndex int) bool
}

type DefaultSplitviewControllerDelegate struct{}

func (d DefaultSplitviewControllerDelegate) SplitViewEffectiveRectForDrawnRectOfDividerAtIndex(obj objc.Object, effectiveRect objc.Object, forDrawnRect objc.Object, dividerIndex int) objc.Object {
	return appkit.ZeroRect
}

func (d DefaultSplitviewControllerDelegate) SplitViewShouldHideDividerAtIndex(obj objc.Object, dividerIndex int) bool {
	return true
}
