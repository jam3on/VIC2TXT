package main

import (
	"fmt"
	"os"
	"time"
	"vic2txt/internal/jsontxt"

	"github.com/akamensky/argparse"
	"github.com/briandowns/spinner"
)

// function main of the program
func main() {

	// creation of a new parser
	parser := argparse.NewParser("vic2txt", "Tool for converting the vic database from json format to txt format for forensic tools.")

	// help function
	parser.HelpFunc = func(c *argparse.Command, msg interface{}) string {
		var help string
		help += fmt.Sprintf("./vic2txt.exe -input(-i) C:/Users/User/Desktop/vic_database.json(filepath) -output(-o) C:/Users/User/Desktop/(directory)")
		return help
	}

	// user input elements
	inFilePtr := parser.String("i", "input", &argparse.Options{Required: true, Help: "File containing vic database to parse (.json)"})
	outDirPtr := parser.String("o", "output", &argparse.Options{Required: true, Help: "Destination folder of the created txt file", Default: "C:/"})

	// Parsing arguments
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	// creating an infinite progress bar
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	s.Suffix = " ...Conversion in progress..."
	s.Start()

	// parsing arguments
	jsontxt.RunJsontoTxt(*inFilePtr, *outDirPtr)

	// stopping the progress bar
	s.Stop()
}
