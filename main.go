package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/common-nighthawk/go-figure"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"

	"github.com/gugadev/virusgotal/requests"
	"github.com/gugadev/virusgotal/utils"
)

func main() {
	var FileUtils utils.Files
	var Console utils.Console

	// console outs
	green := color.New(color.FgGreen).Add(color.Bold)
	white := color.New(color.FgWhite).Add(color.Bold)
	whiteUnderline := color.New(color.FgWhite).Add(color.Bold, color.Underline)
	ascii := figure.NewFigure("VirusGotal", "graffiti", false)

	// spin
	spin := spinner.New(spinner.CharSets[39], 200*time.Millisecond)

	// arguments
	key := flag.String("key", "Your VT API Key", "Ex.: dg2nfng304ngdjng234fng4tfnfjn34")
	path := flag.String("file", "Path of the file", "/home/janedoe/suspectfile.ext")
	flag.Parse()

	if key == nil {
		log.Fatal("You need to provide your VirusTotal API key")
		os.Exit(-1)
	}
	if path == nil {
		log.Fatal("You need to provide the file path")
		os.Exit(-1)
	}

	// open the file to send to VT
	file := FileUtils.Open(*path)
	fileStat, _ := file.Stat()

	Console.Clear()

	ascii.Print()

	green.Print("\nScanning file ")
	white.Printf("%s ", fileStat.Name())
	spin.Start()

	// upload and scan the file
	scan := requests.Upload(*key, file)

	spin.Stop()
	whiteUnderline.Print("\n\nScan results:\n\n")

	// get the report
	report := requests.GetReport(*key, scan.Resource)

	// show the report into a table
	report.Show()
}
