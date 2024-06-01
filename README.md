# TODO

- [ ] Faire une image docker de l'app
- [ ] Setup un github workflow pour deploy l'image sur le registry github

- [ ] Show an elapsed time when "start" is enter.
- [ ] Create a task.
  - [ ] Title
  - [ ] Description
  - [ ] Tags

# Project goflow

The idea of this project is to emulate the same system of the [Fomodor](https://flowmodor.com/)
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

Create DB container

```bash
make docker-run
```

Shutdown DB container

```bash
make docker-down
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

