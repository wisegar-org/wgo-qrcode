package recognizer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func getAllowedImgs() [2]string {
	return [2]string{".jpg", ".png"}
}

func isImgAllowed(ext string) bool {
	allowedImages := getAllowedImgs()
	isAllowed := false
	for _, img := range allowedImages {
		if strings.EqualFold(strings.ToLower(ext), strings.ToLower(img)) {
			isAllowed = true
			break
		}
	}
	return isAllowed
}

func CheckInputFileParam(inputFile string) (string, error) {
	_, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return "", errors.New("input file not valid")
	}

	inputFileExt := filepath.Ext(inputFile)
	inputFileName := strings.Replace(inputFile, inputFileExt, "", 1)

	if isImgAllowed(inputFileExt) {
		return inputFileName + strings.ToLower(inputFileExt), nil
	}
	return "", errors.New("input file extension not supported")
}

func CheckOutputFileParam(inputFile string) (string, error) {

	inputFileDir := filepath.Dir(inputFile)
	inputFileExt := filepath.Ext(inputFile)
	inputFileName := strings.Replace(inputFile, inputFileDir, "", 1)

	outputFile := inputFileDir + strings.Replace(inputFileName, inputFileExt, "", 1) + ".txt"

	f, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		errorValue := fmt.Sprintf("error opening output file: %v", err)
		return "", errors.New(errorValue)
	}
	f.Close()

	e, err := os.Stat(outputFile)
	if err != nil {
		errorValue := fmt.Sprintf("output file not valid: %v", err)
		return "", errors.New(errorValue)
	}
	if e.IsDir() {
		return "", errors.New("output file can not be a folder")
	}
	return outputFile, nil
}
