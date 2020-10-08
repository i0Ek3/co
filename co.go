package co

import (
    "os"
    "io"
    "strings"
    //"strconv"
    "unicode"

    log "github.com/sirupsen/logrus"
)

const (
    OB = "obfuscate"
    DE = "deobfuscate"
    NO = ""
)

type Confuse struct {
    input  io.Reader
    status string
    algoed bool
    algoid int
    cobit  int
}

type codeObfuscate struct {
    Confuse
}

type codeDeobfuscate struct {
    Confuse
}

type confuser interface {
    New() *Confuse
    Init()
    checkStatus(status string) bool
    checkID(id int) bool
    processFile(f *os.File) *os.File
    process(code string)

    // for code obfuscate
    Obfuscate(code string, id ...int)
    Coalgo(id int, code []string)
    coalgo1(code string) string
    coalgo2(code string) string
    coalgo3(code string) []string

    // for code deobfuscate
    Deobfuscate(code string, id ...int)
    Dealgo(id int, code []string)
    dealgo1(code string) string
    dealgo2(code string) string
    dealgo3(code string) string
}

func (c *Confuse) New() *Confuse {
    co := &Confuse {
        status: c.status,
        algoed: c.algoed,
        algoid: c.algoid,
        cobit:  c.cobit,
    }
    return co
}

func (c *Confuse) Init() {
    if c.input == nil {
        c.status = NO
        c.algoed = false
        c.cobit  = 0
    }
    if c.checkStatus(c.status) {
        c.algoed = true
        c.cobit  = 8
    }
}

func (c *Confuse) checkStatus(status string) bool {
    if status == OB || status == DE {
        return true
    }
    return false
}

func (c *Confuse) checkID(id int) bool {
    if c.algoid >= 1 && c.algoid <= 3 {
        return true
    }
    return false
}

func (c *codeObfuscate) process(code string) {
    if c.algoed && c.checkStatus(c.status) {
        switch c.status {
        case OB:
            if c.checkID(c.algoid) {
                c.Obfuscate(code, c.algoid)
            } else {
                log.Fatalf("wrong encoding number.")
            }
        }
    } else {
        switch c.status {
        case OB:
            c.Obfuscate(code)
        }
    }
}

func (c *codeObfuscate) Obfuscate(code string, id ...int) {
    if len(id) > 0 {
        c.algoid = id[0]
        if c.checkID(c.algoid) {
            c.Coalgo(c.algoid, code)
        }
    }
}

func (c *codeObfuscate) Coalgo(id int, code string) {
    switch id {
    case 1:
        c.coalgo1(code)
    case 2:
        c.coalgo2(code)
    case 3:
        c.coalgo3(code)
    }
}

// coalgo1 encoding the code string with ordinary offset
// transform, which means map alphabet to next n alphabet
// with all lower case.
func (c *codeObfuscate) coalgo1(code string) string {
    var encode string
    var alphabet map[int]rune

    encode = strings.ToLower(code)

    for i := 0; i < 26; i++ {
        for j := 'a'; j <= 'z'; j++ {
            alphabet[i] = j
        }
    }

    for i, _ := range code {
        num := (int)(code[i]) - 48 - 97 + c.cobit - 26
        if num > 0 {
            encode = string(alphabet[num])
        }
        encode = string(alphabet[num+26])
    }

    return encode
}

// coalgo2 encoding the code string with ordinary offset
// transform, which means map alphabet to next n alphabet
// with all upper case.
func (c *codeObfuscate) coalgo2(code string) string {
    var encode string
    var alphabet map[int]rune

    encode = strings.ToUpper(code)

    for i := 0; i < 26; i++ {
        for j := 'A'; j <= 'Z'; j++ {
            alphabet[i] = j
        }
    }

    for i, _ := range code {
        num := (int)(code[i]) - 48 - 65 + c.cobit - 26
        if num > 0 {
            encode = string(alphabet[num])
        }
        encode = string(alphabet[num+26])
    }

    return encode
}

// coalgo2 encoding the code string with ordinary offset
// transform, but mapped with special characters like _.
func (c *codeObfuscate) coalgo3(code string) []string {
    var alphabetu map[int]rune
    var alphabetl map[int]rune

    upper := 0
    lower := 0
    specChar := []string{"_", "-"}

    for i, _ := range code {
        if unicode.IsUpper((rune)(code[i])) {
            upper++
        } else if unicode.IsLower((rune)(code[i])) {
            lower++
        }
    }

    if upper > lower {
        strings.ToUpper(code)
        for i := 0; i < 26; i++ {
            for j := 'A'; j <= 'Z'; j++ {
                alphabetu[i] = j
            }
        }
    }
    strings.ToLower(code)
    for i := 0; i < 26; i++ {
        for j := 'a'; j <= 'z'; j++ {
            alphabetl[i] = j
        }
    }

    // this flow converts char to _ or - in code. If we only
    // use _ and - without other identifiers then we cannot
    // recognize what exactly alphabet is represent, so we
    // need some different identifiers to distinguish the code,
    // just like _ represents a, then 2_ represents b instead
    // of __.
    var newer []string
    for i := 0; i < 26; i++ {
        if i < c.cobit {
            newer[i] = specChar[0] + string(i + 1)
        }
        newer[i] = specChar[1] + string(i % c.cobit) + specChar[0] + string(i + 1)
    }

    for k, _ := range code {
        n := (int)(code[k]) - 48 - 97
        newer = append(newer, newer[n])
    }
    return newer
}

func (c *codeObfuscate) processFile(f *os.File) *os.File {
    // TODO
    return f
}

func (c *codeDeobfuscate) process(code string) {
    if c.algoed && c.checkStatus(c.status) {
        switch c.status {
        case DE:
            if c.checkID(c.algoid) {
                c.Deobfuscate(code, c.algoid)
            } else {
                log.Fatalf("wrong decoding number.")
            }
        }
    } else {
        switch c.status {
        case DE:
            c.Deobfuscate(code)
        }
    }
}

func (c *codeDeobfuscate) Deobfuscate(code string, id ...int) {
    if len(id) > 0 {
        c.algoid = id[0]
        if c.checkID(c.algoid) {
            c.Dealgo(c.algoid, code)
        }
    }
}

func (c *codeDeobfuscate) Dealgo(id int, code string) {
    switch id {
    case 1:
        c.dealgo1(code)
    case 2:
        c.dealgo2(code)
    case 3:
        c.dealgo3(code)
    }
}

func (c *codeDeobfuscate) dealgo1(code string) string {
    // TODO
    return code
}

func (c *codeDeobfuscate) dealgo2(code string) string {
    // TODO
    return code
}

func (c *codeDeobfuscate) dealgo3(code string) string {
    // TODO
    return code
}

func (c *codeDeobfuscate) processFile(f *os.File) *os.File {
    // TODO
    return f
}
