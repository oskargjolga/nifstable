# nifstable
[![Go Test](https://github.com/oskargjolga/nifstable/workflows/Go/badge.svg)](https://github.com/oskargjolga/nifstable/actions)


## Getting Started

### Dependencies
* Recent version of golang or docker installed

### Clone
```bash
git clone https://github.com/oskargjolga/nifstable.git
```
### Executing Program if Golang is Installed
* `go run main.go` will compute a table based on full time results
* `go run main.go -halftime` will compute a table based on half time results

### Executing Program in a Docker Container
#### Using the Makefile
* `make table` for fulltime results
* `make table-halftime` for halftime results

#### Without Make
* `docker build --tag nifstable`
* `docker run nifstable` for fulltime results
* `docker run nifstable -halftime` for halftime results


