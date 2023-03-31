package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	form := tview.NewForm().
		AddInputField("Street", "", 30, nil, nil).
		AddInputField("City", "", 20, nil, nil).
		AddInputField("State", "", 20, nil, nil).
		AddInputField("Zip", "", 10, nil, nil).
		AddButton("Submit", func() {
			
		}).
		AddButton("Clear", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Weather App").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}