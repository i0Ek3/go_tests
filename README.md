# Learning Go with TDD

In this repo, you will see some examples of Go basic syntax, and each directory contains the corresponding topic test file and source code, related topics come from repo [here](https://github.com/quii/learn-go-with-tests). And I expanded it properly and rewrited some code with fully tested. Whatever you learn it for work or something else, it's really worth it, let's Go.


## Content

- hello
- integer
- iteration
- array
- slice
- smi: struct/method/interface
- pe: pointer/error
- maps: map
- di: dependency injection
- mock: mocking
- concurrency
- select
- reflect // a little hard
- sync
- context // a little hard
- rnr: A test based on property // a little hard
- math


## YSK

> YSK: You Should Know.

- platform: go1.15.2 darwin/amd64
- run common test: `go test -v`
- run benchmark test: `go test -bench=.`
- use `go test -cover` to get the test cover ratio
- test file named as xxx_test.go
- test function named as TestXxx and it only recive one arugment
- don't expect to write perfect code at once, please iterate slowly
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
- use `net/http/httptest` to create a mock http server
- use `select` to implement process synchronization
- use `time.After()` to prevent your system from being permanently blocked
- don't use reflection unless you really need it
- nobody likes anonymous nested stuct in the complicated code, so please be nice
- pointer type value cannot use the NumField method in reflect, you need to call Elem() to extract the underlying value before executing this method
- slice have no method NumField in reflect, we should use method Len and Index 
- use `sync.WaitGroup` to synchronize concurrent processes
- use `sync.Mutex` to solve data race issue, and the zero value for a Mutex is an unlocked mutex
- a Mutex must not be copied after first use
- `Use a sync.Mutex or a channel?` see [here](https://github.com/golang/go/wiki/MutexOrChannel)
    - use channels when passing ownership of data 
    - use mutexes for managing state
    - in my ex-company, they use both Mutex and channel, and they use channel to customize a new Mutex
- use `go vet` to check your code always
- don't use type embedding, you'll ignore the impact it brings, and that's hard to track bug down
- use `context` to manage long-running processes
- don't use `context.Value` or you got fired, here's [why](https://faiface.github.io/post/context-should-go-away-go2/)
- use `strings.Builder` to build a string with less memory copy
- use `testing/quick` to test the code with random numbers quickly


## TODO

- error checking
- refactor 


## Credit

Thanks to [@quii](https://github.com/quii) for his contribution. And also I found a good Go tutorial to make you learn Go better, please visit [here](https://golangbot.com/learn-golang-series/) to check it out.


## License

MIT.
