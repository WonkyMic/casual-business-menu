package main

import (
	"wonkymic/casual-business-menu/src"
	"wonkymic/casual-business-menu/views"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Anchor struct {
	main *tview.Pages
	menu *tview.Pages
	sidebar *tview.Pages
	navigation src.Navigation
}

func NewAnchor() Anchor {
	appView := views.NewViews()
	main := appView.Main()
	menu := appView.Menu()
	sidebar := appView.Sidebar()
	return Anchor{
		main,
		menu,
		sidebar,
		src.NewNavigation(main, menu, sidebar),
	}
}

func (a Anchor) Grid() *tview.Grid {
	grid := tview.NewGrid().
		SetRows(5, 0, 5).
		SetColumns(30, 0, 30).
		SetBorders(true)
	grid.AddItem(a.navigation.Draw(), 0, 0, 2, 1, 0, 10, false).
		AddItem(a.menu, 2, 0, 3, 1, 0, 100, false).
		AddItem(a.main, 0, 1, 5, 1, 0, 100, false).
		AddItem(a.sidebar, 0, 2, 5, 1, 0, 100, false)
	return grid
}

func main() {
	app := tview.NewApplication()
	anchor := NewAnchor()
	frame := tview.NewFrame(anchor.Grid()).
		SetBorders(0, 0, 0, 0, 0, 0).
		AddText("Anchor", true, tview.AlignCenter, tcell.ColorYellow)
	if err := app.SetRoot(frame, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}