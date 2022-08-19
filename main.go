package main

import (
	"image/color"
	imageservice "rock-paper-scissors/ImageService"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	appWindow := myApp.NewWindow("Random Image")

	// Menu
	fileMenu := fyne.NewMenu(
		"File",
		fyne.NewMenuItem(
			"Quit",
			func() { myApp.Quit() },
		),
	)

	helpMenu := fyne.NewMenu(
		"Help",
		fyne.NewMenuItem(
			"About",
			func() {
				dialog.ShowCustom(
					"About",
					"Close",
					container.NewVBox(
						widget.NewLabel("Welcome to Unsplash random images"),
						widget.NewLabel("Version: 1.0"),
						widget.NewLabel("Author: Rebekka Orth"),
					),
					appWindow,
				)
			},
		),
	)

	mainMenu := fyne.NewMainMenu(fileMenu, helpMenu)

	appWindow.SetMainMenu(mainMenu)

	text := canvas.NewText("Display random image", color.Black)
	text.Alignment = fyne.TextAlignCenter

	// Define image
	var resource, _ = fyne.LoadResourceFromURLString(imageservice.GetRandomUnsplashImage())
	gopherImg := canvas.NewImageFromResource(resource)
	gopherImg.SetMinSize(fyne.Size{Width: 500, Height: 500})

	// Button
	button := widget.NewButton("New image", func() {
		resource, _ := fyne.LoadResourceFromURLString(imageservice.GetRandomUnsplashImage())
		gopherImg.Resource = resource

		gopherImg.Refresh()
	})

	button.Importance = widget.HighImportance

	box := container.NewVBox(
		text,
		gopherImg,
		button,
	)

	appWindow.SetContent(box)

	appWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	appWindow.ShowAndRun()
}
