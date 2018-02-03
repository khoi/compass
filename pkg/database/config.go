package database

import (
	"encoding/csv"
	"strconv"

	"github.com/khoiracle/sextant/pkg/entry"
	"github.com/khoiracle/sextant/pkg/file"
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
	flock := file.NewFlock(f.dbPath)
	if err := flock.Lock(); err != nil {
		return err
	}
	defer flock.Unlock()

	err := flock.File().Truncate(0)
	_, err = flock.File().Seek(0, 0)
	if err != nil {
		return err
	}

	w := csv.NewWriter(flock.File())
	defer w.Flush()

	for _, e := range entries {
		data := []string{e.Path, strconv.Itoa(e.VisitedCount), strconv.Itoa(e.LastVisited)}
		if err := w.Write(data); err != nil {
			return err
		}
	}

	return nil
}

func (f *fileDb) Read() ([]*entry.Entry, error) {
	flock := file.NewFlock(f.dbPath)
	if err := flock.Lock(); err != nil {
		return nil, err
	}
	defer flock.Unlock()

	r := csv.NewReader(flock.File())
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
	flock := file.NewFlock(f.dbPath)
	if err := flock.Lock(); err != nil {
		return err
	}
	defer flock.Unlock()
	flock.File().Truncate(0)
	return flock.File().Sync()
}

func New(filePath string) (DB, error) {
	return &fileDb{
		dbPath: filePath,
	}, nil
}
