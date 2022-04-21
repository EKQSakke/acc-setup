package main

import (
	"io/ioutil"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cars := getCars()

	app := tview.NewApplication()
	list := tview.NewList()
	for _, car := range cars {
		list.AddItem(car, "", 'a', func() {
			list.Clear()
			tracks := getTracks(car)
			for _, track := range tracks {
				list.AddItem(track, "", 'a', func() {
					list.Clear()
					setups := getSetups(car, track)
					for _, setup := range setups {
						list.AddItem(setup, "", 'a', func() {
							showSetup(car, track, setup, app)
						})
					}
					list.AddItem("Quit", "", 'q', func() {
						app.Stop()
					})

				})
			}
			list.AddItem("Quit", "", 'q', func() {
				app.Stop()
			})

		})
	}
	list.AddItem("Quit", "", 'q', func() {
		app.Stop()
	})

	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}

func getCars() []string {
	home, err := os.UserHomeDir()
	check(err)
	files, err := ioutil.ReadDir(home + "\\Documents\\Assetto Corsa Competizione\\Setups")
	check(err)

	cars := make([]string, len(files))

	for i, f := range files {
		cars[i] = f.Name()
	}
	return cars
}

func getTracks(car string) []string {
	home, err := os.UserHomeDir()
	check(err)
	files, err := ioutil.ReadDir(home + "\\Documents\\Assetto Corsa Competizione\\Setups\\" + car)
	check(err)

	tracks := make([]string, len(files))

	for i, f := range files {
		tracks[i] = f.Name()
	}
	return tracks
}

func getSetups(car string, track string) []string {
	home, err := os.UserHomeDir()
	check(err)
	files, err := ioutil.ReadDir(home + "\\Documents\\Assetto Corsa Competizione\\Setups\\" + car + "\\" + track)
	check(err)

	setups := make([]string, len(files))

	for i, f := range files {
		setups[i] = f.Name()
	}
	return setups
}

func showSetup(car string, track string, setup string, app *tview.Application) {
	home, err := os.UserHomeDir()
	check(err)
	setupFile, err := os.Open(home + "\\Documents\\Assetto Corsa Competizione\\Setups\\" + car + "\\" + track + "\\" + setup)
	check(err)
	defer setupFile.Close()

	setupBytes, err := ioutil.ReadAll(setupFile)
	check(err)

	textView := tview.NewTextView()
	textView.SetText(string(setupBytes))

	textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			textView.SetText("")
			list := tview.NewList()

			setups := getSetups(car, track)
			for _, setup := range setups {
				list.AddItem(setup, "", 'a', func() {
					showSetup(car, track, setup, app)
				})
			}
			list.AddItem("Quit", "", 'q', func() {
				app.Stop()
			})

			if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
				panic(err)
			}
		}
		return event
	})

	if err := app.SetRoot(textView, true).SetFocus(textView).Run(); err != nil {
		panic(err)
	}
}
