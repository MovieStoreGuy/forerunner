# ForeRunner
[![GoDoc](https://godoc.org/github.com/MovieStoreGuy/forerunner?status.svg)](https://godoc.org/github.com/MovieStoreGuy/forerunner)
[![Maintainability](https://api.codeclimate.com/v1/badges/3b6eff078d4c45b158d0/maintainability)](https://codeclimate.com/github/MovieStoreGuy/forerunner/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/MovieStoreGuy/forerunner)](https://goreportcard.com/report/github.com/MovieStoreGuy/forerunner)
[![Build Status](https://travis-ci.org/MovieStoreGuy/forerunner.svg?branch=master)](https://travis-ci.org/MovieStoreGuy/forerunner)  
Forerunner is a Golang application to allow for automated CI/CD manor of testing of Docker images.

## Requirements
iCurrently, in order to build Forerunner you will need:
- Golang v1.8+

To run ForeRunner, you will need:
- Docker install
    - Docker daemon running
- Golang v1.8+ 
    - If you are running this via the source code

## Usage
In order to use the forerunner application, you will need to do the following:
```sh
forerunner --path path/to/config.yaml image [images...]
```
With forerunner, it is possuible to test mutliple images consecutively but does require
that they each use the same forerunner config.

The yaml file looks like this:
```yaml
---
# Non Optional Arguements
Commands:
    - cmd1
    - cmd2
    - cmd3
# Optional configs
Network: <bridge|host|none|custom>
Environemnt:
    - <var>=<value>
    - ...
```
