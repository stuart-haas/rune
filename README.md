# GO Job Runner

## Install

```
go build -o runner
chmod +x ./runner
```

## Usage

### Run Task in Sequence

`./runner nginx.yml`

### Run specific Task

`./runner nginx.yml verify`

## TODO

* Add Cobra/Viper for the CLI
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