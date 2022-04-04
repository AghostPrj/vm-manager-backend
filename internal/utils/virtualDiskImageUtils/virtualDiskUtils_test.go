package virtualDiskImageUtils

import (
	"github.com/ggg17226/aghost-go-base/pkg/utils/fileUtils"
	"github.com/ggg17226/aghost-go-base/pkg/utils/randomUtils"
	"os"
	"testing"
)

func TestCreateNewImage(t *testing.T) {
	randomFileName := "/tmp/" + randomUtils.RandStringWithLength(16) + ".img"
	defer deleteFileIfExists(randomFileName)
	err := CreateNewImage(randomFileName, 16777216)
	if err != nil {
		t.Fatalf("create image file error:  %s", err.Error())
	}
}

func deleteFileIfExists(path string) {
	if fileUtils.FileExists(path) {
		os.Remove(path)
	}
}

func TestCopyImage(t *testing.T) {
	randomSourceFileName := "/tmp/" + randomUtils.RandStringWithLength(16) + ".img"
	randomTargetFileName := "/tmp/" + randomUtils.RandStringWithLength(16) + ".img"

	defer deleteFileIfExists(randomSourceFileName)
	defer deleteFileIfExists(randomTargetFileName)

	err := CreateNewImage(randomSourceFileName, 16777216)
	if err != nil {
		t.Fatalf("create image file error:  %s", err.Error())
	}

	err = CopyImage(randomSourceFileName, randomTargetFileName)
	if err != nil {
		t.Fatalf("copy image file error:  %s", err.Error())
	}
}
