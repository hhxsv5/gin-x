package helper

import "os"

func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

func FileIsDir(file string) bool {
	f, err := os.Stat(file)
	if err != nil {
		return false
	}
	return f.IsDir()
}
