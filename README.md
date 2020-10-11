# co

`co` is a Go tool to obfuscate/deobfuscate the code, there are three algorithms you can choose to obfuscate/deobfuscate the code string. But for now, co only support simple code obfuscation/deobfuscation, we'll add some complicated encoding/decoding algorithms later.

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
    co.processOB(code)
    //co.Obfuscate(file)
    co.processDE(code)
    //co.Deobfuscate(file)
}
```

## License

MIT.
