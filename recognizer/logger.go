package recognizer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var filename = ""

func SetFilename(inputFile string) {
	inputFileDir := filepath.Dir(inputFile)
	timenow := time.Now()
	timeset := fmt.Sprintf("%d%s%d", timenow.Day(), timenow.Month(), timenow.Year())
	filename = inputFileDir + "/" + timeset + ".log"
}

func Log(info ...interface{}) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	log.SetOutput(f)
	log.Println(info...)
	f.Close()
}

func LogFatal(info ...interface{}) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	log.SetOutput(f)
	log.Fatalln(info...)
	f.Close()
}
