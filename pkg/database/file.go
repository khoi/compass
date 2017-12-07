package database

import (
	"os"
	"syscall"
)

func createOrOpenLockedFile(filePath string) (file *os.File, err error) {
	if _, serr := os.Stat(filePath); os.IsNotExist(serr) {
		file, err = os.Create(filePath)
	} else {
		file, err = os.OpenFile(filePath, os.O_RDWR, 0644)
	}

	if err != nil {
		return
	}

	if err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX); err != nil {
		return
	}

	return
}

func closeLockedFile(file *os.File) error {
	if err := syscall.Flock(int(file.Fd()), syscall.LOCK_UN); err != nil {
		return err
	}
	return file.Close()
}
