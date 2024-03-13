package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

const jobRunnerConfigPath = "jobs.yml"

type Tasks map[string]Task

type Task struct {
	Description *string
	Depends     []string
	Commands    []interface{}
	Error       *struct {
		Tasks
	}
	Success *struct {
		Tasks
	}
}

type Job struct {
	args        []string
	Version     string
	Name        *string
	Description *string
	Tasks
	Success *Task
	Error   *Task
}

type JobRunnerConfig struct {
	Version string
	Root    string
}

type JobRunner struct {
	config *JobRunnerConfig
	job    *Job
}

func NewJobRunner() (*JobRunner, error) {
	config, err := loadRootConfig()
	if err != nil {
		return nil, err
	}
	return &JobRunner{
		config: config,
		job:    nil,
	}, nil
}

func (r *JobRunner) Run(args []string) error {
	err := r.loadJobConfig(args)
	if err != nil {
		return err
	}
	if len(args) == 1 {
		return r.RunTasks(r.job.Tasks)
	}
	if len(args) == 2 {
		return r.RunTask(r.job.args[1])
	}
	return nil
}

func (r *JobRunner) RunTasks(tasks Tasks) error {
	for _, task := range tasks {
		if err := r.RunTask(task); err != nil {
			return err
		}
	}
	return nil
}

func (r *JobRunner) RunTask(key interface{}) error {
	task := r.getTask(key)

	if task.Depends != nil {
		fmt.Printf("[STARTING TASK] %s\n", key)
		fmt.Printf("[EXECUTING TASK DEPENDENCIES] %s\n", task.Depends)

		for _, dep := range task.Depends {
			if err := r.RunTask(dep); err != nil {
				return err
			}
		}
	}

	fmt.Printf("[EXECUTING TASK] %s\n", key)

	for _, command := range task.Commands {
		switch t := command.(type) {
		case string:
			cmd := exec.Command("bash", "-c", t)

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				if task.Error != nil {
					return r.RunTasks(task.Error.Tasks)
				}
				return err
			}
		default:
			//
		}
	}
	if task.Success != nil {
		return r.RunTasks(task.Success.Tasks)
	}
	return nil
}

func (r *JobRunner) getTask(key interface{}) *Task {
	switch t := key.(type) {
	case string:
		task := r.job.Tasks[key.(string)]
		return &task
	case Task:
		return &t
	default:
	}
	return nil
}

func (r *JobRunner) loadJobConfig(args []string) error {
	if len(args) == 0 {
		return errors.New("PLEASE PROVIDE THE JOB FILE NAME")
	}

	filePath := fmt.Sprintf("%s/%s", r.config.Root, args[0])
	fmt.Printf("[LOADING JOB] %s\n", filePath)

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	r.job = &Job{}

	err = yaml.Unmarshal(bytes, r.job)
	if err != nil {
		return err
	}

	r.job.args = args

	return nil
}

func loadRootConfig() (*JobRunnerConfig, error) {
	fmt.Printf("[LOADING ROOT CONFIG] %s\n", jobRunnerConfigPath)

	bytes, err := os.ReadFile(jobRunnerConfigPath)
	if err != nil {
		return nil, err
	}

	config := &JobRunnerConfig{}

	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	args := os.Args[1:]

	jobRunner, err := NewJobRunner()
	if err != nil {
		panic(err)
	}

	if err := jobRunner.Run(args); err != nil {
		panic(err)
	}
}
