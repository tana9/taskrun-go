package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type TaskRunner struct {
	Dir string
}

type Taskfile struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (t Task) String() string {
	return fmt.Sprintf("%s\t%s", t.Name, t.Desc)
}

func (r *TaskRunner) ListAll() ([]Task, error) {
	out, err := exec.Command("task", "--dir", r.Dir, "--list-all", "--json").Output()
	if err != nil {
		return nil, err
	}
	tasks, err := parseTasks(out)
	if err != nil {
		return nil, err
	}
	return tasks, err
}

func (r *TaskRunner) Run(task Task) error {
	cmd := exec.Command("task", "--dir", r.Dir, task.Name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func parseTasks(data []byte) ([]Task, error) {
	var taskfile Taskfile
	if err := json.Unmarshal(data, &taskfile); err != nil {
		return nil, err
	}
	return taskfile.Tasks, nil
}
