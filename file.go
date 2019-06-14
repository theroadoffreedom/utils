package utils

import (
	"io"
	"net/http"
	"os"
)

const (
	TMP_DIR = "/tmp"
)

func Download(url string, filePath string) error {

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func Delete(filePath string) error {
	err := os.Remove(filePath)
	return err
}

func RandomTmpFilePath() string {
	return TMP_DIR + "/" + NewUUID()
}
