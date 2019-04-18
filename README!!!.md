# Trains

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

### Structer of project

This prject have main two main part of code

  - Root : location of main.go file controll main process of this application
  - Core : location of go file to calculate logic

### Flow of application
  - Receive input and parse into two dimension map
  - From 10 problem we can separate into 4 types of problems 
    - (1-5) find distance from specific route
    - (6,7) find number of route by limit number of stop
    - (8-9) find shorted path
    - (10) find number of route by limit of distance
    
  - Send map to logic function for calculate answer
  - Receive answer from each fucntion and print to console

### Calculate logic
  - GROUP-1 find distance from specific route :
    go through the route follow by the problem, get distance from map, sum and response anser

  - GROUP-2 find number of route by limit number of stop : recursive through possible route until find the end-point or until reached limit of stop, sum correct route and return

  - GROUP-3 find shorted path : recursive through possible route until find the end-point or until not found any new route, remember lowest distace, sum lowest distace and return

  - GROUP-4 find number of route by limit of distance : recursive through possible route until find the end-point or until reached limit of distance, sum correct route and return

### Installation!!!

  - Install the dependencies and devDependencies, build application.

    ```sh
    $ go mod tidy
    $ go build
    ```

### Run

It root of working directory after build application
Run application with graph path as input.

Example (run with binary file): 

```sh
$ ./ThoughWorkAssignMent.exe "AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"
```
or
Example (run with main.go file): 

```sh
$ go run main.go "AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"
```

### Run Test
To run all test use this command in root directory
```
$ go test ./... -v
```




