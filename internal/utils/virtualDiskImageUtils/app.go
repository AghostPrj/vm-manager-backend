package virtualDiskImageUtils

import (
	"errors"
	"github.com/AghostPrj/vm-manager-backend/internal/utils/commandUtils"
	"github.com/ggg17226/aghost-go-base/pkg/utils/fileUtils"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

func CreateNewImage(path string, size uint64) error {
	if !filepath.IsAbs(path) {
		return errors.New("path must be abs")
	}

	if fileUtils.FileExists(path) {
		return errors.New("file already exist")
	}
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 755)
	if err != nil {
		return err
	}
	cmd := "/usr/bin/qemu-img create -f " + VirtualDiskFormat + " '" + path + "' " + strconv.FormatUint(size, 10)
	exitCode, err := commandUtils.ExecCommand(cmd)
	if err != nil {
		return err
	}
	if exitCode != 0 {
		if fileUtils.FileExists(path) {
			os.Remove(path)
		}
		return errors.New("create image file error")
	}

	return nil
}

func checkFileOpened(path string) (bool, error) {
	if !fileUtils.FileExists(path) {
		return false, nil
	}
	cmd := "/usr/bin/lsof '" + path + "'"
	exitCode, err := commandUtils.ExecCommand(cmd)
	if err != nil {
		return false, err
	}
	return exitCode == 0, nil
}

func CheckImage(path string) error {
	if !filepath.IsAbs(path) {
		return errors.New("path must be abs")
	}

	if !fileUtils.FileExists(path) {
		return errors.New("file not exist")
	}
	cmd := "/usr/bin/qemu-img check '" + path + "'"
	exitCode, err := commandUtils.ExecCommand(cmd)
	if err != nil {
		return err
	}
	if exitCode != 0 {
		return errors.New("check image file error")
	}
	return nil
}

func CopyImage(source string, target string) error {
	if (!filepath.IsAbs(source)) || (!filepath.IsAbs(target)) {
		return errors.New("path must be abs")
	}

	if !fileUtils.FileExists(source) {
		return errors.New("source file not exist")
	}

	if fileUtils.FileExists(target) {
		return errors.New("target file already exist")
	}

	opened, err := checkFileOpened(source)
	if err != nil {
		return err
	}

	if opened {
		return errors.New("source file has opened")
	}

	err = CheckImage(source)
	if err != nil {
		return err
	}

	dir := filepath.Dir(target)
	err = os.MkdirAll(dir, 755)
	if err != nil {
		return err
	}

	targetFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	_, err = io.Copy(targetFile, sourceFile)
	return err

}
