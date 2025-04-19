package store

import (
	"encoding/json"
	"errors"
	"os"
)

type Store struct {
	FilePath string
	Data     map[string]uint32
}

func NewStore(filePath string) *Store {
	return &Store{
		FilePath: filePath,
		Data:     make(map[string]uint32),
	}
}

func (s *Store) SaveRecord() error {
	data, err := json.Marshal(s.Data)
	if err != nil {
		return err
	}
	return os.WriteFile(s.FilePath, data, 0644)
}

func (s *Store) LoadRecord() error {
	data, err := os.ReadFile(s.FilePath)
	if errors.Is(err, os.ErrNotExist) {
		return s.SaveRecord()
	} else if err != nil {
		return err
	}
	return json.Unmarshal(data, &s.Data)
}
