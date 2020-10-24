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
- use `[]struct{}{}` 

 
## Content

- hello
- integer
- iteration
- array
- slice
- smi: struct/method/interface


## Credit

Thanks to [@quii](https://github.com/quii) for his contribution.

## License

MIT.
