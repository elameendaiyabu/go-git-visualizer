# Go Git Visualizer

Visualize local git commits from the command line - built with go

![git visualizer](./git-visualizer)

## run locally

### prerequisites

- git
- go

### directions

- clone project

```git
git clone git@github.com:elameendaiyabu/go-git-visualizer.git
```

- cd into project

```bash
cd go-git-visualizer
```

- install required go packages

```go
go mod tidy
```

- add folders to include in commits visual

```go
go run . --add ~
```

- go run with your git config email passed as argument

```go
go run . --email your@email.com
```

## Tech Stack

- go
- go-git

## how it works

- get a list of folders to scan
- generate git commits visuals from those folders
