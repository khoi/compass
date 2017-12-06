package database

import (
	"encoding/csv"
	"os/user"
	"path/filepath"
	"strconv"

	"github.com/khoiln/sextant/search"
)

const (
	defaultFileName = ".sextant"
)

type DB interface {
	Read() ([]*search.Entry, error)
	Write([]*search.Entry) error
	Truncate() error
}

type fileDb struct {
	dbPath string
}

func (f *fileDb) Write(entries []*search.Entry) error {
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

func (f *fileDb) Read() ([]*search.Entry, error) {
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

	entries := make([]*search.Entry, 0, len(records))
	for _, r := range records {
		visitedCount, err := strconv.Atoi(r[1])
		lastVisited, err := strconv.Atoi(r[2])

		if err != nil {
			continue
		}

		entries = append(entries, &search.Entry{
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

func New(fileName string) (DB, error) {
	usr, err := user.Current()

	if err != nil {
		return nil, err
	}

	fullFilePath := filepath.Join(usr.HomeDir, fileName)

	return &fileDb{
		dbPath: fullFilePath,
	}, nil
}

func NewDefault() (DB, error) {
	return New(defaultFileName)
}
