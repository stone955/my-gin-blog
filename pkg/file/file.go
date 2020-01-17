package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

func GetSize(f multipart.File) (int, error) {
	bytes, err := ioutil.ReadAll(f)
	return len(bytes), err
}

func GetExt(fileName string) string {
	return path.Ext(fileName)
}

func CheckNotExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return os.IsNotExist(err)
}

func CheckPermission(fileName string) bool {
	_, err := os.Stat(fileName)
	return os.IsPermission(err)
}

func MkDirIfNotExist(fileName string) error {
	if notExist := CheckNotExist(fileName); notExist {
		if err := Mkdir(fileName); err != nil {
			return err
		}
	}
	return nil
}

func Mkdir(fileName string) error {
	if err := os.MkdirAll(fileName, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
