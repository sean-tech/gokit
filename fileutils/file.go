package fileutils

import (
	"bufio"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}


func ReadFile(path string) ([]byte, error) {
	inputFile, inputError := os.Open(path)
	if inputError != nil {
		return nil, inputError
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	buf := make([]byte, 1024)
	for {

		n, err := inputReader.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if (n == 0) { break}
	}
	return buf, nil
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	srcFile, err := os.Open(srcName)
	if err != nil {
		return -1, err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dstName)
	if err != nil {
		return -1, err
	}
	defer dstFile.Close()

	if written, err = io.Copy(dstFile, srcFile); err != nil {
		return
	}
	if err := dstFile.Sync(); err != nil {
		return -1, err
	}
	return
}

func ClearFile(srcName string) error {
	var file *os.File
	var err error
	if file, err = os.OpenFile(srcName, os.O_RDWR, 0644); err != nil {
		return err
	}
	defer file.Close()
	if err = file.Truncate(0); err != nil {
		return err
	}
	return nil
}

func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}

func GetExt(fileName string) string {
	return path.Ext(fileName)
}

func CheckExist(src string) bool {
	_, err := os.Stat(src)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

func MKDirIfNotExist(src string) error {
	if exist := CheckExist(src); exist == false {
		if err := MKDir(src); err != nil {
			return err
		}
	}
	return nil
}

func MKDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

