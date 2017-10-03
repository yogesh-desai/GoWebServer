# GoWebServer [![Build Status](https://travis-ci.org/yogesh-desai/GoWebServer.svg?branch=master)](https://travis-ci.org/yogesh-desai/GoWebServer)

Simple Web Server to get the files and and report if file is not present.

The tree structure if as follows,
```
.
├── api
│ └── v1
│     └── getfile
│         ├── file1.txt
│         ├── file2.txt
│         └── file3.txt
├── main.go
├── README.md
└── test.txt

```
Basically we have handler to check the dir location of api/v1/getfile.

## Dependencies
Currently, there are no dependencies. We have used all the in-built libraries and packages.

## Usage

```
$ go run main.go

* Open any one URL in browser window.
      1. http://localhost:8080/api/v1/getfile/<fileName>
      2. http://localhost:8080/test

```


## ToDo's

Many things can be done. This is just the basic example.
Please feel free to generate pull requests or issues. :)