<div align="center">
<h1>RUNE</h1>
<p>
A mythical and legendary task runner.
</p>
<img src="https://github.com/stuart-haas/rune/assets/12514075/5f7300bb-5f24-4c6e-a379-a9ff3028c7df" />
</div>

## Install

```
go run . internal.yml link
```

## Usage

### Run specific Task

`rune internal.yml verify`

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
