## TODO

## Project

## Time

- [x] Show an elapsed time based on "start" and "stop" command.
- [x] Display the elapsing time every second.
- [x] Replace each tick by the new one.
- [x] Show elapsing time at the top
- [x] On "break" command show the break time and replace the current elapsing
      time by the break time.

## TUI Time

- [ ] Add "Reset" option that set to null start, stop, time and break time

- [ ] On "Break Time" option select, replace the current choice by :
  - [ ] "Start Break Time"
  - [ ] "Stop Break Time"
  - [ ] "Skip Break Time"
    - [ ] Open a quick window and ask user with a "Yes" "No"
          option if he really want to skip the break time.

## Task

- [ ] Create a task.

  - [ ] Title
  - [ ] Description
  - [ ] Tags

- [ ] Write the task to a file in the format of a JSON.
      (Consider using an interface so we can implement other techniques later)
- [ ] Start the application and check if there is a file with tasks.

## TUI Task (later)

# Project goflow

The idea of this project is to emulate the same system of the [Flowmodor](https://flowmodor.com/)
website but in the terminal.

Flowmodor works as follows:

- You can create a task
- You can start a task and it will start counting the time
- You can stop a task and it will stop counting the time.
  - When you stop, a break time will be calculated based on the amount of
    time you worked divided by 5.

The data will be stored using through a REDIS database runned using docker.

## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)

## MakeFile

run all make commands with clean tests

```bash
make all build
```

build the application

```bash
make build
```

run the application

```bash
make run
```

build the docker image of the application

```bash
make docker-build
```

run the docker image of the application and create a volume
to store the data written by the app.

```bash
make docker-run
```

live reload the application

```bash
make watch
```

run the test suite

```bash
make test
```

clean up binary from the last build

```bash
make clean
```
