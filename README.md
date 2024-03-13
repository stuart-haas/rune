# Rune

> A mythical and legendary task runner.

![dune_787564](https://github.com/stuart-haas/rune/assets/12514075/96ce5749-2047-407c-8b20-03ab4eb08b1b)

## Install

```
go run . internal.yml link
```

## Usage

### Run default task

`rune internal.yml`

### Run specific Task

`rune internal.yml verify`

## TODO

* Add Cobra/Viper for the CLI
* More useful keys for tasks
* Add flags
  ```
  --ignore-errors # skips errors
  --task # specify tasks or tasks to run in sequence
  --sequential # execute tasks in sequence (default)
  --parallel # execute tasks in parallel
  --ignore-tasks # specify task or tasks to ignore
  --config # specifiy location of runner config
  --work-dir # specifiy working directory
  --install # build and make symbol link in /usr/local/bin
  --docker # support executing from Docker image
  ```
