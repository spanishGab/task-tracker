package tasks

import (
	"encoding/json"
	"fmt"
	"slices"
	"tasktracker/src/ports"

	"time"
)

type TaskRepository struct {
	DataSource   string
	dbConnection ports.IFileHandler
}

func NewTaskRepository(dataSource string, dbConnection ports.IFileHandler) *TaskRepository {
	return &TaskRepository{
		DataSource:   dataSource,
		dbConnection: dbConnection,
	}
}

func (tr *TaskRepository) CreateOne(task Task) (*Task, error) {
	tasks, err := tr.getAllTasks()
	if err != nil {
		return nil, err
	}
	task.ID = tr.getNextId(tasks)
	now := time.Now().UTC()
	task.CreatedAt = now
	task.UpdatedAt = now

	tasks = append(tasks, task)

	data, err := tr.tasksToBytes(tasks)
	if err != nil {
		return nil, err
	}
	if _, err := tr.dbConnection.Write(data); err != nil {
		return nil, err
	}
	return &task, nil
}

func (tr *TaskRepository) DeleteOne(id uint64) error {
	tasks, err := tr.getAllTasks()
	if err != nil {
		return err
	}
	var taskPosition int
	for i, task := range tasks {
		if id == task.ID {
			taskPosition = i
			break
		}
	}
	tasks = slices.Delete(tasks, taskPosition, taskPosition+1)

	data, err := tr.tasksToBytes(tasks)
	if err != nil {
		return err
	}
	if _, err := tr.dbConnection.Write(data); err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) UpdateOne(task Task) (*Task, error) {
	tasks, err := tr.getAllTasks()
	if err != nil {
		return nil, err
	}
	for i := range tasks {
		if tasks[i].ID == task.ID {
			tasks[i].Description = task.Description
			tasks[i].Status = task.Status
			tasks[i].UpdatedAt = time.Now().UTC()
			break
		}
	}

	data, err := tr.tasksToBytes(tasks)
	if err != nil {
		return nil, err
	}
	if _, err := tr.dbConnection.Write(data); err != nil {
		return nil, err
	}
	return &task, nil
}

func (tr *TaskRepository) getAllTasks() ([]Task, error) {
	unmarshalledTasks, err := tr.dbConnection.Read()
	if err != nil {
		return nil, fmt.Errorf("error while trying to read all tasks: %s", err.Error())
	}
	if len(unmarshalledTasks) <= 0 {
		return []Task{}, nil
	}

	var tasks []Task
	err = json.Unmarshal(unmarshalledTasks, &tasks)
	if err != nil {
		return nil, fmt.Errorf("error while trying to unmarshal all tasks: %s", err.Error())
	}
	return tasks, nil
}

func (tr *TaskRepository) tasksToBytes(tasks []Task) ([]byte, error) {
	data, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("error while trying to marshal all tasks: %s", err.Error())
	}
	return data, nil
}

func (tr *TaskRepository) getNextId(tasks []Task) uint64 {
	var biggestID uint64 = 1
	for _, task := range tasks {
		if task.ID > biggestID {
			biggestID = task.ID
		}
	}
	biggestID++
	return biggestID
}
