package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/SkuldNorniern/Cordium/frame"
)

func firstScreen(a fyne.App) fyne.CanvasObject {

	return widget.NewVBox(
		layout.NewSpacer(),
	)
}

func main() {
	a := app.NewWithID("Wormhole")

	w := a.NewWindow("Wormhole")
	w.SetMaster()

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Settings", theme.HomeIcon(), firstScreen(a)),
		widget.NewTabItemWithIcon("Device", theme.DocumentCreateIcon(), firstScreen(a)),
		widget.NewTabItemWithIcon("Host", theme.DocumentCreateIcon(), frame.DialogScreen(w)))

	//tabs.SelectTabIndex(a.Preferences().Int(preferenceCurrentTab))
	//network.Hostsmb()
	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()
	tabs.SetTabLocation(widget.TabLocationLeading)
	w.SetContent(tabs)
	w.ShowAndRun()
	//a.Preferences().SetInt(preferenceCurrentTab, tabs.CurrentTabIndex())
}
