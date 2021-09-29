package views

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type ImageFrameStyle uint
type ImageAlignment uint
type ImageScaling uint

const (
	NSImageFrameNone ImageFrameStyle = iota
	NSImageFramePhoto
	NSImageFrameGrayBezel
	NSImageFrameGroove
	NSImageFrameButton
)

const (
	NSImageAlignCenter ImageAlignment = iota
	NSImageAlignTop
	NSImageAlignTopLeft
	NSImageAlignTopRight
	NSImageAlignLeft
	NSImageAlignBottom
	NSImageAlignBottomLeft
	NSImageAlignBottomRight
	NSImageAlignRight
)

const (
	NSImageScaleProportionallyDown ImageScaling = iota
	NSImageScaleAxesIndependently
	NSImageScaleNone
	NSImageScaleProportionallyUpOrDown
)

type ImageView struct {
	objc.Object
}

func NewImageView(image cocoa.NSImage) ImageView {
	return ImageView{objc.Get("NSImageView").Send("imageViewWithImage:", image)}
}

func NewImageViewWithBytes(bytes []byte) ImageView {
	data := core.NSData_WithBytes(bytes, uint64(len(bytes)))
	image := cocoa.NSImage_InitWithData(data)
	return NewImageView(image)
}

func (i ImageView) SetImageFrameStyle(style ImageFrameStyle) {
	i.Set("imageFrameStyle:", style)
}

func (i ImageView) SetImageAlignment(alignment ImageAlignment) {
	i.Set("imageAlignment:", alignment)
}

func (i ImageView) SetImageScaling(scaling ImageScaling) {
	i.Set("imageScaling:", scaling)
}

func (i ImageView) SetAnimates(animates bool) {
	i.Set("animates:", animates)
}
