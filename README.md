# nifstable
[![Go Test](https://github.com/oskargjolga/nifstable/workflows/Go/badge.svg)](https://github.com/oskargjolga/nifstable/actions)
[![Coverage Status](https://coveralls.io/repos/github/oskargjolga/nifstable/badge.svg?branch=main)](https://coveralls.io/github/oskargjolga/nifstable?branch=main)

## Getting Started

### Dependencies
* Recent version of Golang or Docker installed

### Clone
```bash
git clone https://github.com/oskargjolga/nifstable.git
```
### Usage (if Golang is Installed)
Compute a table based on full time results:
```bash
go run main.go
```
Compute a table based on half time results:
```bash
go run main.go - halftime
```
### Usage (Docker Container)
#### With Make
```bash
make table
```
or:
```bash
make table-halftime
```

#### Without Make
Build image:
```bash
docker build --tag nifstable .
```
Run image as a container:
```bash
docker run nifstable
```
or:
```bash
docker run nifstable -halftime
```
