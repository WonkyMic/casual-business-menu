package views

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Weather struct {
	street string
	city string
	state string
	zip string
}

func (w Weather) setStreet(s string) {
	w.street = s
}

func (w Weather) setCity(c string) {
	w.city = c
}

func (w Weather) setState(s string) {
	w.state = s
}

func (w Weather) setZip(z string) {
	w.zip = z
}

func NewWeather() Weather {
	return Weather{"100 Military Plaza #4", "San Antonio", "Texas", "78254"}
}

func results() *tview.Pages {
	pages := tview.NewPages()
		pages.AddPage("blank", tview.NewTextView().SetText("blank"), true, true)
		pages.AddPage("result", tview.NewTextView().SetText("result").SetBackgroundColor(tcell.ColorBlue), true, false)
	return pages
}

func (w Weather) Main() tview.Primitive {
	results := results()
	form := tview.NewForm().
		AddInputField("Street", w.street, 30, nil, w.setStreet).
		AddInputField("City", w.city, 20, nil, w.setCity).
		AddInputField("State", w.state, 20, nil, w.setState).
		AddInputField("Zip", w.zip, 10, nil, w.setZip).
		AddButton("Submit", func() {
			results.SwitchToPage("result")
		}).
		AddButton("Clear", func() {
			results.SwitchToPage("blank")
		})

	grid := tview.NewGrid().
			SetRows(2).
			AddItem(form, 0, 0, 3, 1, 0, 0, false).
			AddItem(results, 3, 0, 2, 1, 0, 0, false)
	grid.SetBorder(true).SetTitle("Weather App").SetTitleAlign(tview.AlignLeft)
	return grid
}

func (w Weather) Menu() tview.Primitive {
	return tview.NewTextView().SetText("WeatherMenu")
}

func (w Weather) Sidebar() tview.Primitive {
	return tview.NewTextView().SetText("WeatherSidebar")
}