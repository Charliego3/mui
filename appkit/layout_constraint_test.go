package appkit

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"testing"
)

func TestNewLayoutConstraint(t *testing.T) {
	constraint := NewLayoutConstraint()
	t.Logf("Constraint: %#v", constraint)
}

func TestNewLayoutConstraintWithAttr(t *testing.T) {
	rect := core.Rect(0, 0, 100, 100)
	rootView := cocoa.NSView_Init(rect)
	subView := cocoa.NSView_Init(rect)

	constraint := NewLayoutConstraintWithAttr(subView,
		NSLayoutAttributeLeft,
		NSLayoutRelationEqual,
		rootView,
		NSLayoutAttributeLeft,
		1.0, 10.0,
	)

	t.Logf("%#v", constraint)
}

func TestNewLayoutConstraintWithFormat(t *testing.T) {
	rect := core.Rect(0, 0, 100, 100)
	rootView := cocoa.NSView_Init(rect)
	subView := cocoa.NSView_Init(rect)

	constraint := NewLayoutConstraintWithFormat(rootView, subView)

	t.Logf("ConstraintWithFormat: %#v", constraint)
}