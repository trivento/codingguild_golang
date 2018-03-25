# Go

The following GO modules are present here:

* Basic: examples containing some basic GO code constructs.
* Gossip: a network gossiper for creating a network of communicating nodes. 

```
export GOPATH=$HOME/go
go get github.com/trivento/codingguild_golang/basics
go get github.com/trivento/codingguild_golang/gossip
```

## Creating a new package
``` 
mkdir $GOPATH/src/github.com/trivento/codingguild_golang/yourpackage
touch  $GOPATH/src/github.com/trivento/codingguild_golang/yourpackage/main.go
```
and add the following code to `main.go`
```
package main

import (
	"fmt"
)

func main() {

	fmt.Println("Ready to go")

}
```
Create an run your package
```
    go install github.com/trivento/codingguild_golang/yourpackage && $GOPATH/bin/yourpackage
```
Note that in Go the resuling exectuable will contain all dependent code.

## Tooling
* Visual Studio Code [https://code.visualstudio.com/docs/languages/go]
* Goland/IntelliJ

## Book: The Go Programming Language 

* Website [http://www.gopl.io]. 
* Source code [https://github.com/adonovan/gopl.io/]
* You can download the first chapter [http://www.gopl.io/ch1.pdf]

You can use all the examples like this:

```
export GOPATH=$HOME/gobook
go get gopl.io/ch1/helloworld
$GOPATH/bin/helloworld
``` 

