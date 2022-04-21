package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/skratchdot/open-golang/open"
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
	w.Resize(fyne.NewSize(400, 400))

	listTitle := widget.NewLabel("")
	listContainer := container.NewVBox()

	carSelectionView(cars, listContainer, listTitle)

	w.SetContent(container.NewVBox(
		listTitle,
		listContainer,
	))

	w.ShowAndRun()
}

func carSelectionView(cars []string, listContainer *fyne.Container, listTitle *widget.Label) {
	listTitle.SetText("Select a car")
	list := make([]fyne.CanvasObject, len(cars))
	for i, car := range cars {
		myCar := car
		list[i] = widget.NewButton(car, func() {
			trackSelectionView(myCar, getTracks(myCar), listContainer, listTitle)
		})
		listContainer.Add(list[i])
	}
}

func trackSelectionView(car string, tracks []string, listContainer *fyne.Container, listTitle *widget.Label) {
	listTitle.SetText("Select a track")
	listContainer.Objects = nil
	list := make([]fyne.CanvasObject, len(tracks))
	for i, track := range tracks {
		myTrack := track
		list[i] = widget.NewButton(track, func() {
			setupSelectionView(car, myTrack, getSetups(car, myTrack), listContainer, listTitle)
		})
		listContainer.Add(list[i])
	}
}

func setupSelectionView(car, track string, setups []string, listContainer *fyne.Container, listTitle *widget.Label) {
	listTitle.SetText("Select a setup") 
	listContainer.Objects = nil
	list := make([]fyne.CanvasObject, len(setups))
	for i, setup := range setups {
		mySetup := setup
		list[i] = container.NewHBox(
			widget.NewLabel(mySetup),
			widget.NewButton("Open", func() {
				showSetup(car, track, mySetup)
			}),
			widget.NewButton("Copy", func() {
				copySetupToAllTracks(car, track, mySetup)
			}),
		)
		listContainer.Add(list[i])
	}
}

func setupsPath() string {
	home, err := os.UserHomeDir()
	check(err)
	return home + "\\Documents\\Assetto Corsa Competizione\\Setups\\"
}

func getCars() []string {
	files, err := ioutil.ReadDir(setupsPath())
	check(err)

	cars := make([]string, len(files))

	for i, f := range files {
		cars[i] = f.Name()
	}
	return cars
}

func getTracks(car string) []string {
	files, err := ioutil.ReadDir(setupsPath() + car)
	check(err)

	tracks := make([]string, len(files))

	for i, f := range files {
		tracks[i] = f.Name()
	}
	return tracks
}

func getSetups(car, track string) []string {
	files, err := ioutil.ReadDir(setupsPath() + car + "\\" + track)
	check(err)

	setups := make([]string, len(files))

	for i, f := range files {
		setups[i] = f.Name()
	}
	return setups
}

func showSetup(car, track, setup string) {
	// This doesn't work while debugging - it seems to be a bug in the package
	open.Start(setupsPath() + car + "\\" + track + "\\" + setup)
}

func copySetupToAllTracks(car, sourceTrack, setup string) {
	f, err := os.ReadFile("tracklist.txt")
	check(err)
	trackString := string(f)
	for _, destinationTrack := range strings.Split(trackString, "\r\n") {
		if destinationTrack != "" && sourceTrack != destinationTrack {
			if !Contains(getTracks(car), destinationTrack) {
				os.Mkdir(setupsPath()+car+"\\"+destinationTrack, 0755)
			}
			copySetup(car, sourceTrack, destinationTrack, setup)
		}
	}
}

func copySetup(car, sourceTrack, destinationTrack, setup string) {
	source, err := os.Open(setupsPath() + car + "\\" + sourceTrack + "\\" + setup)
	check(err)
	defer source.Close()

	destination, err := os.Create(setupsPath() + car + "\\" + destinationTrack + "\\" + setup)
	check(err)
	defer destination.Close()

	_, err = io.Copy(destination, source)
	check(err)
}

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
