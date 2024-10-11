// main.go
package main

import (
	"flag"
	"os"

	"github.com/cmbrock1/chip8_emulator_go/pkg/emulator"
	"github.com/sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
)

var logger = logrus.New()

func main() {

	verbosePtr := flag.Bool("v", false, "enable verbose mode for debugging")
	flag.Parse()
	args := flag.Args()

	if *verbosePtr {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.WarnLevel)
	}

	if len(args) == 0 {
		logger.Fatalln("Not Enough Arguments...")
	}

	romFilepath := args[0]
	checkForValidFile(romFilepath)

	sdl.Init(sdl.INIT_EVERYTHING)
	emu := emulator.NewEmulator()
	logger.WithField("emu", emu).Warn("I got a new emulator")
}

func checkForValidFile(romFilepath string) {
	file, err := os.Open(romFilepath)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"file": romFilepath,
		}).Fatalln("Error opening the file")
	}
	defer file.Close()
}
