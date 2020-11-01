package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/SkuldNorniern/Cordium/frame"
	"github.com/SkuldNorniern/Cordium/network"
	"rsc.io/quote"
)

func main() {
	a := app.NewWithID("Wormhole")
	w := a.NewWindow("Wormhole")
	w.SetMaster()
	fmt.Println(quote.Hello())
	frame.DialogScreen(w)
	network.Hostsmb()
	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()
	w.ShowAndRun()
}
