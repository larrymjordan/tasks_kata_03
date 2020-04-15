package models

import (
	"encoding/json"

	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
)

type TaskStorage interface {
	Add(*Task)
	List() []Task
	Clear()
}

// Tasks is not required by pop and may be deleted
type MemoryTaskStorage []Task

// Add adds a new task to the task storage.
func (t *MemoryTaskStorage) Add(task *Task) {
	task.ID = uuid.Must(uuid.NewV4())
	*t = append(*t, *task)
}

// List get all tasks stored in the tasks storage.
func (t MemoryTaskStorage) List() []Task {
	return t
}

// Clear removes all tasks stored in tasks storage.
func (t *MemoryTaskStorage) Clear() {
	*t = MemoryTaskStorage{}
}

// String is not required by pop and may be deleted
func (t MemoryTaskStorage) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

type DatabaseTaskStorage struct {
	Tx *pop.Connection
}

func (d DatabaseTaskStorage) List() []Task {
	tasks := []Task{}
	d.Tx.All(&tasks)
	return tasks
}

func (d DatabaseTaskStorage) Add(task *Task) {
	d.Tx.Create(task)
}

func (d DatabaseTaskStorage) Clear() {
	for _, t := range d.List() {
		d.Tx.Destroy(&t)
	}
}
