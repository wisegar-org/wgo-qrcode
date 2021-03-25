package main

import (
	"os"

	"github.com/wisegar-org/wgo-qrcode/recognizer"
)

func main() {

	inputFile := os.Args[1]
	recognizer.SetFilename(inputFile)

	recognizer.Log("Executing QR Code recognizer")

	recognizer.Log("Checking input img: ", inputFile)
	inputImageFile, err := recognizer.CheckInputFileParam(inputFile)
	if err != nil {
		recognizer.Log("Impossible to validate the input image file")
		recognizer.Log(err)
		os.Exit(1)
	}

	outputTextFile, err := recognizer.CheckOutputFileParam(inputFile)
	if err != nil {
		recognizer.Log("Impossible to validate the output file")
		recognizer.Log(err)
		os.Exit(1)
	}

	recognizer.Log("Recognizing %s image file", inputImageFile)
	reconizedPayload, err := recognizer.RecognizeFile(inputImageFile)
	if err != nil {
		recognizer.Log("Impossible to validate the input image file")
		recognizer.Log(err)
		os.Exit(1)
	}

	recognizer.Log("Recording image text results on output file  %s", outputTextFile)
	errWriting := recognizer.WriteRecognitionOutput(reconizedPayload, outputTextFile)
	if errWriting != nil {
		recognizer.Log("Impossible to write results the output text file")
		recognizer.Log(errWriting)
		os.Exit(1)
	}
	recognizer.Log("QR Code reconizer sucessfully executed")
}
