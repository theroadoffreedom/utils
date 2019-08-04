package utils

import (
	"os"
	"testing"
)

func TestIsFileExist(t *testing.T) {

	tmpFile := "./tmp_file"
	exist, err := IsFileExist(tmpFile)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if exist {
		t.Error("file should not exist")
		return
	}

	// create file to check
	emptyFile, err := os.Create(tmpFile)
	if err != nil {
		t.Error(err)
		return
	}
	emptyFile.Close()
	exist, err = IsFileExist(tmpFile)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if !exist {
		t.Error("file should be exist")
		return
	}

	// delete file
	err = os.Remove(tmpFile)
	if err != nil {
		t.Error("clean up error")
	}
	t.Log("test IsFileExist success")
}
