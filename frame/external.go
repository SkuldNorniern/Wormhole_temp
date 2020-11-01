//Package external frame
//
// MIT License
//
// Copyright (c) 2020 Cordis Terrarum
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
//
package frame

import (
	"fmt"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/SkuldNorniern/Cordium/network"
)

func confirmCallback(response bool) {
	fmt.Println("Responded with", response)
}
func loadDialogGroup(win fyne.Window) *widget.Group {
	return widget.NewGroup("Dialogs",
		widget.NewButton("Info", func() {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Started Hosting",
				Content: "Your Hosting of your files has been begin",
			})
			network.Hostsmb()
		}),
		widget.NewButton("Login", func() {
			username := widget.NewEntry()
			password := widget.NewPasswordEntry()
			content := widget.NewForm(widget.NewFormItem("Username", username),
				widget.NewFormItem("Password", password))

			dialog.ShowCustomConfirm("Login...", "Log In", "Cancel", content, func(b bool) {
				if !b {
					return
				}

				log.Println("Please Authenticate", username.Text, password.Text)
			}, win)
		}))
}

/*func loadWindowGroup() fyne.Widget {
	windowGroup := widget.NewGroup("Windows",
		widget.NewButton("New window", func() {
			w := fyne.CurrentApp().NewWindow("Hello")
			w.SetContent(widget.NewLabel("Hello World!"))
			w.Show()
		}),
		widget.NewButton("Fixed size window", func() {
			w := fyne.CurrentApp().NewWindow("Fixed")
			w.SetContent(fyne.NewContainerWithLayout(layout.NewCenterLayout(), widget.NewLabel("Hello World!")))

			w.Resize(fyne.NewSize(240, 180))
			w.SetFixedSize(true)
			w.Show()
		}),
		widget.NewButton("Centered window", func() {
			w := fyne.CurrentApp().NewWindow("Central")
			w.SetContent(fyne.NewContainerWithLayout(layout.NewCenterLayout(), widget.NewLabel("Hello World!")))

			w.CenterOnScreen()
			w.Show()
		}))

	drv := fyne.CurrentApp().Driver()
	if drv, ok := drv.(desktop.Driver); ok {
		windowGroup.Append(
			widget.NewButton("Splash Window (only use on start)", func() {
				w := drv.CreateSplashWindow()
				w.SetContent(widget.NewLabelWithStyle("Hello World!\n\nMake a splash!",
					fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
				w.Show()

				go func() {
					time.Sleep(time.Second * 3)
					w.Close()
				}()
			}))
	}

	otherGroup := widget.NewGroup("Other",
		widget.NewButton("Notification", func() {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Started Hosting",
				Content: "Your Hosting of your files has been begin",
			})
		}))

	return widget.NewVBox(windowGroup, otherGroup)
}*/
func DialogScreen(win fyne.Window) fyne.CanvasObject {
	return fyne.NewContainerWithLayout(layout.NewAdaptiveGridLayout(2),
		widget.NewScrollContainer(loadDialogGroup(win)))
}
