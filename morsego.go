package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	inputFile := flag.String("input", "", "input file to read from")
	outputFile := flag.String("output", "", "output file to write to")

	flag.Parse()
	writer := openOutputFile(*outputFile)

	if len(flag.Args()) > 0 { // Input from CLI arguments

		morseFromArray(flag.Args(), writer)

	} else { // Input from file

		reader := openInputFile(*inputFile)

		morseFromFile(reader, writer)

	}

}

func openInputFile(filename string) io.Reader {

	var reader io.Reader

	if filename == "" {
		reader = os.Stdin
	} else {
		var err error
		reader, err = os.Open(filename)

		if err != nil {
			panic(err)
		}
	}

	return reader

}

func openOutputFile(filename string) io.Writer {

	var writer io.Writer

	if filename == "" {
		writer = os.Stdout
	} else {
		var err error
		writer, err = os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)

		if err != nil {
			panic(err)
		}
	}

	return writer
}

func morseFromFile(reader io.Reader, writer io.Writer) {

	buf := make([]byte, 1024)
	n, _ := reader.Read(buf)

	for n > 0 {

		for i := 0; i < n; i++ {
			morseChar := morse(rune(buf[i]))
			fmt.Fprintf(writer, "%s", morseChar)

			if morseChar != "" {
				fmt.Fprintf(writer, " ")
			}

		}

		n, _ = reader.Read(buf)

	}

}

func morseFromArray(texts []string, writer io.Writer) {

	for _, text := range texts {
		words := strings.Fields(text)

		for _, word := range words {
			for _, char := range word {

				morseChar := morse(char)
				fmt.Fprintf(writer, "%s ", morseChar)

			}
		}
	}

}

func morse(letter rune) string {
	switch string(letter) {
	case "A", "a":
		return ".-"
	case "B", "b":
		return "-..."
	case "C", "c":
		return "-.-."
	case "D", "d":
		return "-.."
	case "E", "e":
		return "."
	case "F", "f":
		return "..--."
	case "G", "g":
		return "--."
	case "H", "h":
		return "...."
	case "I", "i":
		return ".."
	case "J", "j":
		return ".---"
	case "K", "k":
		return "-.-"
	case "L", "l":
		return ".-.."
	case "M", "m":
		return "--"
	case "N", "n":
		return "-."
	case "O", "o":
		return "---"
	case "P", "p":
		return ".--."
	case "Q", "q":
		return "--.-"
	case "R", "r":
		return ".-."
	case "S", "s":
		return "..."
	case "T", "t":
		return "-"
	case "U", "u":
		return "..-"
	case "V", "v":
		return "...-"
	case "W", "w":
		return ".--"
	case "X", "x":
		return "-..-"
	case "Y", "y":
		return "-.--"
	case "Z", "z":
		return "--.."
	case "0":
		return "-----"
	case "1":
		return ".----"
	case "2":
		return "..---"
	case "3":
		return "...--"
	case "4":
		return "....-"
	case "5":
		return "....."
	case "6":
		return "-...."
	case "7":
		return "--..."
	case "8":
		return "---.."
	case "9":
		return "----."
	case ".":
		return ".-.-.-"
	case ",":
		return "--..--"
	default:
		return ""
	}
}
