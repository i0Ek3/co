# co

`co` is a Go tool to obfuscate the code, so we call it co(code obfuscation). But for now, co just support simple code obfuscation with my own encoding algorithm.

- basic code obfuscation
- customize encoding algorithm
- file obfuscation/deobfuscation(not implement yet)

## Install

`go get github.com/i0Ek3/co`

## Usage

```Go
package main

import (
    co "github.com/i0Ek3/co"
)

func main() {
    ...
    co.Obfuscate(id)
    //co.Obfuscate(os.File)
    co.Deobfuscate(id)
    //co.Deobfuscate(os.File)
}
```

## License

MIT.
