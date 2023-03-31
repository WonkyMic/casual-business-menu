package views

import (
	// "wonkymic/casual-business-menu/views"

	"github.com/rivo/tview"
)

type Views struct {
	weather Weather
}

func NewViews() Views {
	return Views{NewWeather()}
}

func (v Views) Main() *tview.Pages {
	pages := tview.NewPages()
	pages.AddPage("home", HomeMain(), true, true)
	pages.AddPage("market", tview.NewTextView().SetText("market"), true, false)
	pages.AddPage("crypto", tview.NewTextView().SetText("crypto"), true, false)
	pages.AddPage("stocks", tview.NewTextView().SetText("stocks"), true, false)
	pages.AddPage("music", tview.NewTextView().SetText("music"), true, false)
	pages.AddPage("weather", v.weather.Main(), true, false)
	return pages
}

func (v Views) Menu() *tview.Pages {
	pages := tview.NewPages()
	pages.AddPage("home", HomeMenu(), true, true)
	pages.AddPage("market", tview.NewTextView().SetText("market"), true, false)
	pages.AddPage("crypto", tview.NewTextView().SetText("crypto"), true, false)
	pages.AddPage("stocks", tview.NewTextView().SetText("stocks"), true, false)
	pages.AddPage("music", tview.NewTextView().SetText("music"), true, false)
	pages.AddPage("weather", v.weather.Menu(), true, false)
	return pages
}

func (v Views) Sidebar() *tview.Pages {
	pages := tview.NewPages()
	pages.AddPage("home", HomeSidebar(), true, true)
	pages.AddPage("market", tview.NewTextView().SetText("market"), true, false)
	pages.AddPage("crypto", tview.NewTextView().SetText("crypto"), true, false)
	pages.AddPage("stocks", tview.NewTextView().SetText("stocks"), true, false)
	pages.AddPage("music", tview.NewTextView().SetText("music"), true, false)
	pages.AddPage("weather", v.weather.Sidebar(), true, false)
	return pages
}