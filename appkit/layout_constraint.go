package appkit

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type LayoutFormatOptions uint
type LayoutAttribute uint
type LayoutRelation uint

const (
	NSLayoutAttributeNotAnAttribute LayoutAttribute = iota
	NSLayoutAttributeLeft
	NSLayoutAttributeRight
	NSLayoutAttributeTop
	NSLayoutAttributeBottom
	NSLayoutAttributeLeading
	NSLayoutAttributeTrailing
	NSLayoutAttributeWidth
	NSLayoutAttributeHeight
	NSLayoutAttributeCenterX
	NSLayoutAttributeCenterY
	NSLayoutAttributeLastBaseline
	NSLayoutAttributeFirstBaseline
	NSLayoutAttributeLeftMargin
	NSLayoutAttributeRightMargin
	NSLayoutAttributeTopMargin
	NSLayoutAttributeBottomMargin
	NSLayoutAttributeLeadingMargin
	NSLayoutAttributeTrailingMargin
	NSLayoutAttributeCenterXWithinMargins
	NSLayoutAttributeCenterYWithinMargins
)

const (
	//NSLayoutRelationLessThanOrEqual    LayoutRelation = iota - 1
	NSLayoutRelationEqual              LayoutRelation = 0
	NSLayoutRelationGreaterThanOrEqual                = 1
)

const (
	NSLayoutFormatSpacingMask                LayoutFormatOptions = 0x1 << 19
	NSLayoutFormatSpacingBaselineToBaseline  LayoutFormatOptions = 1 << 19
	NSLayoutFormatSpacingEdgeToEdge          LayoutFormatOptions = 0 << 19
	NSLayoutFormatDirectionMask              LayoutFormatOptions = 0x3 << 16
	NSLayoutFormatDirectionRightToLeft       LayoutFormatOptions = 2 << 16
	NSLayoutFormatDirectionLeftToRight       LayoutFormatOptions = 1 << 16
	NSLayoutFormatDirectionLeadingToTrailing LayoutFormatOptions = 0 << 16
	NSLayoutFormatAlignmentMask              LayoutFormatOptions = 0xFFFF
	NSLayoutFormatAlignAllFirstBaseline      LayoutFormatOptions = 1 << NSLayoutAttributeFirstBaseline
	NSLayoutFormatAlignAllLastBaseline       LayoutFormatOptions = 1 << NSLayoutAttributeLastBaseline
	NSLayoutFormatAlignAllBaseline           LayoutFormatOptions = NSLayoutFormatAlignAllLastBaseline
	NSLayoutFormatAlignAllCenterY            LayoutFormatOptions = 1 << NSLayoutAttributeCenterY
	NSLayoutFormatAlignAllCenterX            LayoutFormatOptions = 1 << NSLayoutAttributeCenterX
	NSLayoutFormatAlignAllTrailing           LayoutFormatOptions = 1 << NSLayoutAttributeTrailing
	NSLayoutFormatAlignAllLeading            LayoutFormatOptions = 1 << NSLayoutAttributeLeading
	NSLayoutFormatAlignAllBottom             LayoutFormatOptions = 1 << NSLayoutAttributeBottom
	NSLayoutFormatAlignAllTop                LayoutFormatOptions = 1 << NSLayoutAttributeTop
	NSLayoutFormatAlignAllRight              LayoutFormatOptions = 1 << NSLayoutAttributeRight
	NSLayoutFormatAlignAllLeft               LayoutFormatOptions = 1 << NSLayoutAttributeLeft
)

type LayoutConstraint struct {
	objc.Object `objc:"LayoutConstraint : NSLayoutConstraint"`
}

var layoutConstraintClass objc.Object

func lazyRegisterLayoutConstraint() {
	if layoutConstraintClass == nil {
		class := objc.NewClassFromStruct(LayoutConstraint{})
		objc.RegisterClass(class)
		layoutConstraintClass = class
	}
}

func NewLayoutConstraintWithFormat(view1, view2 objc.Object) objc.Object {
	lazyRegisterLayoutConstraint()
	views := core.NSDictionary_Init(core.String("view1"), view1, nil)
	metricsDic := core.NSDictionary_Init(core.String("left"), float32(20), core.String("right"), float32(20), core.String("space"), float32(20), core.String("top"), float32(20), nil)
	return layoutConstraintClass.
		Send("constraintsWithVisualFormat:options:metrics:views:",
			core.String("H:|-left-[view1]-space-[view2(view1)]-right-|"),
			NSLayoutFormatAlignAllTop, metricsDic, views)
}

func NewLayoutConstraintWithAttr(subView objc.Object, subAttribute LayoutAttribute, relation LayoutRelation,
	toItem objc.Object, toAttribute LayoutAttribute, multiplier float32, constant float32) LayoutConstraint {
	lazyRegisterLayoutConstraint()
	return LayoutConstraint{layoutConstraintClass.
		Send("constraintWithItem:attribute:relatedBy:toItem:attribute:multiplier:constant:",
			subView, subAttribute, relation, toItem, float32(toAttribute), multiplier, constant)}
}

func NewLayoutConstraint() LayoutConstraint {
	lazyRegisterLayoutConstraint()
	return LayoutConstraint{Object: layoutConstraintClass.Alloc().Init()}
}

//- (void)viewDidLoad {
//[super viewDidLoad];
//// Do any additional setup after loading the view, typically from a nib.
//self.view.backgroundColor = [UIColor yellowColor];
//
//
//UIView *subView = [[UIView alloc] init];
//subView.backgroundColor = [UIColor redColor];
//// 在设置约束前，先将子视图添加进来
//[self.view addSubview:subView];
//
//// 使用autoLayout约束，禁止将AutoresizingMask转换为约束
//[subView setTranslatesAutoresizingMaskIntoConstraints:NO];
//
//// 设置subView相对于VIEW的上左下右各40像素
//LayoutConstraint *constraint1 = [LayoutConstraint constraintWithItem:subView attribute:NSLayoutAttributeTop relatedBy:NSLayoutRelationEqual toItem:self.view attribute:NSLayoutAttributeTop multiplier:1.0 constant:40];
//LayoutConstraint *constraint2 = [LayoutConstraint constraintWithItem:subView attribute:NSLayoutAttributeLeft relatedBy:NSLayoutRelationEqual toItem:self.view attribute:NSLayoutAttributeLeft multiplier:1.0 constant:40];
//// 由于iOS坐标系的原点在左上角，所以设置右边距使用负值
//LayoutConstraint *constraint3 = [LayoutConstraint constraintWithItem:subView attribute:NSLayoutAttributeBottom relatedBy:NSLayoutRelationEqual toItem:self.view attribute:NSLayoutAttributeBottom multiplier:1.0 constant:-40];
//
//// 由于iOS坐标系的原点在左上角，所以设置下边距使用负值
//LayoutConstraint *constraint4 = [LayoutConstraint constraintWithItem:subView attribute:NSLayoutAttributeRight relatedBy:NSLayoutRelationEqual toItem:self.view attribute:NSLayoutAttributeRight multiplier:1.0 constant:-40];
//
//// 将四条约束加进数组中
//NSArray *array = [NSArray arrayWithObjects:constraint1, constraint2, constraint3, constraint4 ,nil];
//// 把约束条件设置到父视图的Contraints中
//[self.view addConstraints:array];
//}
