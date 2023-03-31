package src

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Navigation struct {
	main *tview.Pages
	menu *tview.Pages
	sidebar *tview.Pages
}

func NewNavigation(main *tview.Pages, menu *tview.Pages, sidebar *tview.Pages) Navigation {
	return Navigation{main, menu, sidebar}
}

func marketNode() *tview.TreeNode {
	node := tview.NewTreeNode("market").
			SetColor(tcell.ColorYellow).
			SetSelectable(true).
			SetExpanded(false).
			AddChild(navigationNode("crypto")).
			AddChild(navigationNode("stocks"))
	return node
}

func navigationNode(name string) *tview.TreeNode {
	node := tview.NewTreeNode(name).
			SetSelectable(true)
	return node
}

func (nav Navigation) Draw() *tview.Grid {
	root := navigationNode("home").
			SetColor(tcell.ColorYellow).
			AddChild(marketNode()).
			AddChild(navigationNode("music")).
			AddChild(navigationNode("weather"))
	tree := tview.NewTreeView().
			SetRoot(root).
			SetCurrentNode(root).
			SetSelectedFunc(func(node *tview.TreeNode) {
				key := node.GetText()
				if "home" != key {
					node.SetExpanded(!node.IsExpanded())
				}
				nav.main.SwitchToPage(key)
				nav.menu.SwitchToPage(key)
				nav.sidebar.SwitchToPage(key)
			})
	title := tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText("Navigation")
	grid := tview.NewGrid().
			SetRows(2).
			SetBorders(false).
			AddItem(title, 0, 0, 1, 1, 0, 0, false).
			AddItem(tree, 1, 0, 1, 1, 0, 0, false)
	return grid
}