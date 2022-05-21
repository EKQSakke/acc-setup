package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cars := getCars()
	a := app.New()
	w := a.NewWindow("ACC Setup Manager")
	w.Resize(fyne.NewSize(500, 20))

	listContainer := container.NewVBox()
	topContainer := container.NewHBox()

	carSelectionView(cars, listContainer, topContainer)

	w.SetContent(container.NewVBox(
		topContainer,
		listContainer,
	))

	w.ShowAndRun()
}

func carSelectionView(cars []string, listContainer *fyne.Container, topContainer *fyne.Container) {
	title := widget.NewLabel("Select a car")
	topContainer.Objects = nil
	topContainer.Add(title)

	listContainer.Objects = nil
	list := make([]fyne.CanvasObject, len(cars))
	for i, car := range cars {
		myCar := car
		list[i] = widget.NewButton(car, func() {
			trackSelectionView(myCar, getTracks(myCar), listContainer, topContainer)
		})
		listContainer.Add(list[i])
	}
}

func trackSelectionView(car string, tracks []string, listContainer *fyne.Container, topContainer *fyne.Container) {
	title := widget.NewLabel("Select a track")
	backButton := widget.NewButton("<", func() {
		carSelectionView(getCars(), listContainer, topContainer)
	})
	topContainer.Objects = nil
	topContainer.Add(backButton)
	topContainer.Add(title)

	listContainer.Objects = nil
	list := make([]fyne.CanvasObject, len(tracks))
	for i, track := range tracks {
		myTrack := track
		list[i] = widget.NewButton(track, func() {
			setupSelectionView(car, myTrack, getSetups(car, myTrack), listContainer, topContainer)
		})
		listContainer.Add(list[i])
	}
}

func setupSelectionView(car, track string, setups []string, listContainer *fyne.Container, topContainer *fyne.Container) {
	title := widget.NewLabel("Select a setup")
	backButton := widget.NewButton("<", func() {
		trackSelectionView(car, getTracks(car), listContainer, topContainer)
	})
	topContainer.Objects = nil
	topContainer.Add(backButton)
	topContainer.Add(title)

	listContainer.Objects = nil
	list := make([]fyne.CanvasObject, len(setups))
	for i, setup := range setups {
		mySetup := setup
		list[i] = container.NewHBox(
			widget.NewLabel(mySetup),
			widget.NewButton("Open", func() {
				showSetup(car, track, mySetup)
			}),
			widget.NewButton("Copy to all tracks", func() {
				copySetupToAllTracks(car, track, mySetup)
			}),
		)
		listContainer.Add(list[i])
	}
}
