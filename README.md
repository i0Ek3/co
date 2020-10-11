# co

`co` is a Go tool to obfuscate/deobfuscate the code, there are three algorithms you can choose to obfuscate/deobfuscate. But for now, co just support simple code obfuscation/deobfuscation with my own encoding/decoding algorithm, we'll add some complicated encoding/decoding algorithms later.

- basic code obfuscation/deobfuscation
- customize encoding/decoding algorithm
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
    co.process(code)
    co.Obfuscate(algoid)
    //co.Obfuscate(file)
    co.Deobfuscate(algoid)
    //co.Deobfuscate(file)
}
```

## License

MIT.
