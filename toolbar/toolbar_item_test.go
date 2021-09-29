package toolbar

import (
	"github.com/progrium/macdriver/objc"
	"testing"
)

func TestNewNSToolBarItem(t *testing.T) {
	t.Logf("%#v", objc.Get("NSToolbarCloudSharingItemIdentifier"))

	item := NewItem("AppToolbarItem")
	item.SetLabel("ToolBar1")
	item.SetToolTip("This is tooltip")
	item.SetImage("system")

	t.Logf("ToolbarItem: %#v", item)
	t.Logf("ToolbarItem.Identifier: %+v", item.Identifier())
}
