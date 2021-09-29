package views

import "github.com/progrium/macdriver/cocoa"

// VisualEffectMaterial ==============================================================================
type VisualEffectMaterial uint

const (
	NSVisualEffectMaterialTitleBar  VisualEffectMaterial = iota + 3 // 3
	NSVisualEffectMaterialSelection                                 // 4
	NSVisualEffectMaterialMenu                                      // 5
	NSVisualEffectMaterialPopover                                   // 6
	NSVisualEffectMaterialSidebar                                   // 7
	_
	_
	NSVisualEffectMaterialHeaderView       // 10
	NSVisualEffectMaterialSheet            // 11
	NSVisualEffectMaterialWindowBackground // 12
	NSVisualEffectMaterialHUDWindow        // 13
	_
	NSVisualEffectMaterialFullScreenUI // 15
	_
	NSVisualEffectMaterialToolTip           // 17
	NSVisualEffectMaterialContentBackground // 18
	_
	_
	NSVisualEffectMaterialUnderWindowBackground // 21
	NSVisualEffectMaterialUnderPageBackground   // 22
)

// VisualEffectBlendingMode ==============================================================================
type VisualEffectBlendingMode uint

const (
	NSVisualEffectBlendingModeBehindWindow VisualEffectBlendingMode = iota // 0
	NSVisualEffectBlendingModeWithinWindow                                 // 1
)

// BackgroundStyle ==============================================================================
type BackgroundStyle uint

const (
	NSBackgroundStyleNormal BackgroundStyle = iota
	NSBackgroundStyleEmphasized
	NSBackgroundStyleRaised
	NSBackgroundStyleLowered
)

// VisualEffectState ==============================================================================
type VisualEffectState uint

const (
	NSVisualEffectStateFollowsWindowActiveState VisualEffectState = iota
	NSVisualEffectStateActive
	NSVisualEffectStateInactive
)

type VisualEffectView struct {
	cocoa.NSVisualEffectView
}

func NewVisualEffectView() VisualEffectView {
	return VisualEffectView{cocoa.NSVisualEffectView_New()}
}

func (view VisualEffectView) SetMaterial(material VisualEffectMaterial) {
	view.Set("material:", material)
}

func (view VisualEffectView) SetBlendingMode(blendingMode VisualEffectBlendingMode) {
	view.Set("blendingMode:", blendingMode)
}

// SetEmphasized 指示是否强调材质的外观
func (view VisualEffectView) SetEmphasized(emphasized bool) {
	view.Set("emphasized:", emphasized)
}

// SetInteriorBackgroundStyle 视图的内部背景样式
func (view VisualEffectView) SetInteriorBackgroundStyle(style BackgroundStyle) {
	view.Set("interiorBackgroundStyle:", style)
}

// SetMaskImage 其 alpha 通道掩盖了视觉效果视图的材质
func (view VisualEffectView) SetMaskImage(style cocoa.NSImage) {
	view.Set("maskImage:", style)
}

// SetState 指示视图是否应用了视觉效果
func (view VisualEffectView) SetState(state VisualEffectState) {
	view.Set("state:", state)
}
