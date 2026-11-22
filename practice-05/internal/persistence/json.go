package persistence

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"example/go-practice-01/practice-05/internal/tasks"
)

func Load(path string) (tasks.State, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return tasks.State{NextID: 1}, nil
	}
	if err != nil {
		return tasks.State{}, fmt.Errorf("чтение %q: %w", path, err)
	}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()

	var state tasks.State

	if err := decoder.Decode(&state); err != nil {
		return tasks.State{}, fmt.Errorf("декодирование %q: %w", path, err)
	}
	var extra any
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		return tasks.State{}, fmt.Errorf("лишние данные после JSON в %s", path)
	}
	store := tasks.NewMemoryStore()
	if err := store.Restore(state); err != nil {
		return tasks.State{}, fmt.Errorf("некорректное состояние в %q: %w", path, err)
	}
	return state, nil
}

// TODO: validate the state, marshal it using encoding/json, save to a
// temporary file in the same directory and replace the target using os.Rename.
// Do not remove the old file before a successful temporary write.
func Save(path string, state tasks.State) error {
	store := tasks.NewMemoryStore()
	if err := store.Restore(state); err != nil {
		return fmt.Errorf("проверка состояния перед сохранением: %w", err)
	}
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return fmt.Errorf("создание JSON: %w", err)
	}
	data = append(data, '\n')
	temp, err := os.CreateTemp(filepath.Dir(path), ".tasks-*.tmp")
	tempPath := temp.Name()
	defer os.Remove(tempPath)
	defer temp.Close()
	if _, err := temp.Write(data); err != nil {
		return fmt.Errorf("запись временного файла: %w", err)
	}
	if err := temp.Sync(); err != nil {
		return fmt.Errorf("синхронизация временного файла: %w", err)
	}
	if err := temp.Close(); err != nil {
		return fmt.Errorf("закрытие временного файла: %w", err)
	}
	if err := os.Rename(tempPath, path); err != nil {
		return fmt.Errorf("замена файла %q: %w", path, err)
	}
	return nil
}
