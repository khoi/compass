package database

import (
	"encoding/csv"
	"os/user"
	"path/filepath"
	"strconv"

	"github.com/khoiln/sextant/pkg/entry"
)

const (
	defaultFileName = ".sextant"
)

type DB interface {
	Read() ([]*entry.Entry, error)
	Write([]*entry.Entry) error
	Truncate() error
}

type fileDb struct {
	dbPath string
}

func (f *fileDb) Write(entries []*entry.Entry) error {
	file, err := createOrOpenLockedFile(f.dbPath)
	if err != nil {
		return err
	}
	defer closeLockedFile(file)

	w := csv.NewWriter(file)
	defer w.Flush()

	for _, entry := range entries {
		data := []string{entry.Path, strconv.Itoa(entry.VisitedCount), strconv.Itoa(entry.LastVisited)}
		if err := w.Write(data); err != nil {
			return nil
		}
	}

	return nil
}

func (f *fileDb) Read() ([]*entry.Entry, error) {
	file, err := createOrOpenLockedFile(f.dbPath)
	if err != nil {
		return nil, err
	}
	defer closeLockedFile(file)

	r := csv.NewReader(file)
	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	entries := make([]*entry.Entry, 0, len(records))
	for _, r := range records {
		visitedCount, err := strconv.Atoi(r[1])
		lastVisited, err := strconv.Atoi(r[2])

		if err != nil {
			continue
		}

		entries = append(entries, &entry.Entry{
			Path:         r[0],
			VisitedCount: visitedCount,
			LastVisited:  lastVisited,
		})
	}

	return entries, nil
}

func (f *fileDb) Truncate() error {
	file, err := createOrOpenLockedFile(f.dbPath)
	if err != nil {
		return err
	}
	defer closeLockedFile(file)

	file.Truncate(0)
	return file.Sync()
}

func New(filePath string) (DB, error) {
	return &fileDb{
		dbPath: filePath,
	}, nil
}

func NewDefault() (DB, error) {
	usr, err := user.Current()

	if err != nil {
		return nil, err
	}

	fullFilePath := filepath.Join(usr.HomeDir, defaultFileName)

	return New(fullFilePath)
}
