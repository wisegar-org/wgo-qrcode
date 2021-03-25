package recognizer

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/liyue201/goqr"
)

func RecognizeFile(inputPath string) ([]*goqr.QRData, error) {

	imgdata, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(bytes.NewReader(imgdata))
	if err != nil {
		return nil, err
	}
	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		return nil, err
	}
	return qrCodes, nil
}

func WriteRecognitionOutput(qrCodes []*goqr.QRData, outputFile string) error {

	f, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	for _, qrCode := range qrCodes {
		_, err = f.Write([]byte(fmt.Sprintf("%s\n", qrCode.Payload)))
		if err != nil {
			f.Close()
			return err
		}
	}
	f.Close()
	return nil
}
