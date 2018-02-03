package file

import (
	"os"
	"syscall"
)

type Flock struct {
	path string
	file *os.File
}

func NewFlock(path string) *Flock {
	return &Flock{path: path}
}

func (f *Flock) File() *os.File {
	return f.file
}

func (f *Flock) Lock() error {
	if f.file == nil {
		if err := f.createOrOpenFile(); err != nil {
			return err
		}
	}

	if err := syscall.Flock(int(f.file.Fd()), syscall.LOCK_EX); err != nil {
		return err
	}

	return nil
}

func (f *Flock) Unlock() error {
	if err := syscall.Flock(int(f.file.Fd()), syscall.LOCK_UN); err != nil {
		return err
	}
	return f.file.Close()
}

func (f *Flock) createOrOpenFile() error {
	file, err := os.OpenFile(f.path, os.O_CREATE|os.O_RDWR, os.FileMode(0644))
	if err != nil {
		return err
	}
	f.file = file
	return nil
}
