# co

`co` is a Go tool to obfuscate and deobfuscate the code string or code file, there are three algorithms you can choose to obfuscate and deobfuscate them. But for now, co only support simple code obfuscation and deobfuscation, we'll add some complicated obfuscation and deobfuscation algorithms later. Also there are so many flaws, so, please be nice.


## Feature

- basic code string obfuscation/deobfuscation
- file obfuscation/deobfuscation
- multiple alternative obfuscation/deobfuscation algorithms
- support enable and disable debug mode
- no third-party libraries 
- customize the error message display


## Install

`go get github.com/i0Ek3/co`


## Usage

`go run co.go` or `go build ; ./co`


## TODO

- need to refactor
- write good comments
- perfect the tests


## License

MIT.
