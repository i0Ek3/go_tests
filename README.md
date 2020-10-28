# Learning Go with TDD

In this repo, each directory contains the corresponding topic test file and source code, related topics come from repo: `https://github.com/quii/learn-go-with-tests`. And I expanded it properly and rewrited some code with fully tested. Whatever you learn it for work or something else, it's really worth it, let's Go.

## YSK

> YSK: You Should Know.

- platform: go1.15.2 darwin/amd64
- run common test: `go test -v`
- run benchmark test: `go test -bench=.`
- use `go test -cover` to gets the test cover ratio
- test file named as xxx_test.go
- test function named as TestXxx and it only recive one arugment
- write the simple test and test it first, then write the main code and refactor them, REPL
- example function cannot excute if you forget the comment `// Output: result`
- write the good comment for all you functions please, especially exported ones
- use table driven tests to make test file easier to expand and maintain
- use `[]struct{}{}` to make your code hierarchical and maintain better
- when a function or method was invocated, parameters will be copied. So, always use `*Struct` be a reciver 
- type alias can be useful, also you can declare methods on them: `type Transaction uintptr`
- import `errors`, cause of error checking always useful, and read [this one](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
- use `_, ok` to check if operation is ok, ok is a bool value
- never initialize an empty map variable like `var m map[int]string`, cause of nil pointer exception. You should do this: `m = make(map[int]string)` or `m = map[int]string{}`
- if you want to put the data somewhere, please use `io.Writer`, it is a good general interface, better than `bytes.Buffer`
- DI makes separation of concerns: decoupling where the data arrives and how it is generated, and reuse the code in the different situation
- try to get your code to be tested as soon as possible
- when we want to start a goroutine, we often use anonymous functions, just like this: `go func() {}()`
- to enable race detector, run the test with the race flag: `go test -race` 
- use channel to pass the data in goroutine, `ch <- data` is send data and `data := <-ch` is receive data


## Content

- hello
- integer
- iteration
- array
- slice
- smi: struct/method/interface
- pe: pointer/error
- maps: map
- **di: dependency injection**
- **mock: mocking**
- **concurrency: goroutine, channel**
 

## Credit

Thanks to [@quii](https://github.com/quii) for his contribution. And also I found a good Go tutorial to make you learn Go better, please visit [here](https://golangbot.com/learn-golang-series/) to check it out.

## License

MIT.
