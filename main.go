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
	Commands    []interface{}
	Error       *struct {
		Tasks
	}
	Success *struct {
		Tasks
	}
}

type Job struct {
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
}

func NewJobRunner() (*JobRunner, error) {
	config, err := loadRootConfig()
	if err != nil {
		return nil, err
	}
	return &JobRunner{
		config,
	}, nil
}

func (r *JobRunner) Run(args []string) error {
	if len(args) == 0 {
		return errors.New("PLEASE PROVIDE THE JOB FILE NAME")
	}
	job, err := r.loadJobConfig(args[0])
	if err != nil {
		return err
	}
	if len(args) == 1 {
		return r.RunTasks(job.Tasks)
	}
	if len(args) == 2 {
		task := job.Tasks[args[1]]
		return r.RunTask(args[1], task)
	}
	return nil
}

func (r *JobRunner) RunTasks(tasks map[string]Task) error {
	for key, task := range tasks {
		if err := r.RunTask(key, task); err != nil {
			return err
		}
	}
	return nil
}

func (r *JobRunner) RunTask(key string, task Task) error {
	fmt.Printf("[EXECUTING TASKS] %s\n", key)

	if task.Description != nil {
		fmt.Printf("  %s\n", *task.Description)
	}

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

func (r *JobRunner) loadJobConfig(jobName string) (*Job, error) {
	filePath := fmt.Sprintf("%s/%s", r.config.Root, jobName)
	fmt.Printf("[LOADING JOB] %s\n", filePath)

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	job := &Job{}

	err = yaml.Unmarshal(bytes, job)
	if err != nil {
		return nil, err
	}

	if job.Name != nil {
		fmt.Println(job.Name)
	}

	return job, nil
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
