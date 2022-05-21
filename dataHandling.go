package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

func setupsPath() string {
	home, err := os.UserHomeDir()
	check(err)
	return home + "\\Documents\\Assetto Corsa Competizione\\Setups\\"
}

// Get a list of cars with folders
func getCars() []string {
	files, err := ioutil.ReadDir(setupsPath())
	check(err)

	cars := make([]string, len(files))

	for i, f := range files {
		cars[i] = f.Name()
	}
	return cars
}

// Get a list of track folders for a car
func getTracks(car string) []string {
	files, err := ioutil.ReadDir(setupsPath() + car)
	check(err)

	tracks := make([]string, len(files))

	for i, f := range files {
		tracks[i] = f.Name()
	}
	return tracks
}

// Get a list of setups for a car and track
func getSetups(car, track string) []string {
	files, err := ioutil.ReadDir(setupsPath() + car + "\\" + track)
	check(err)

	setups := make([]string, len(files))

	for i, f := range files {
		setups[i] = f.Name()
	}
	return setups
}

// Displays a setup json in a text editor
func showSetup(car, track, setup string) {
	// This doesn't work while debugging - it seems to be a bug in the package
	open.Start(setupsPath() + car + "\\" + track + "\\" + setup)
}

// Copies a setup to all tracks
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

// Copies a setup from one track to another
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
