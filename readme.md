## go-archive

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ferhatelmas/go-archive)
[![Build Status](https://travis-ci.org/ferhatelmas/go-archive.png?branch=master)](https://travis-ci.org/ferhatelmas/go-archive)

> Check if a filepath is an archive file.

### Install

```
go get github.com/ferhatelmas/go-archive
```

### Usage

```go
import "github.com/ferhatelmas/go-archive"

archive.Is("src/unicorn.zip")
//=> true

archive.Is("src/unicorn.txt")
//=> false
```
### Related
[goarchext](https://github.com/shamsher31/goarchext)

### License

MIT © [ferhat elmas](http://ferhatelmas.com)
