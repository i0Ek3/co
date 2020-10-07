package co

import (
    "os"
    "io"

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
    Coalgo(id int, code string)
    coalgo1(code string) string
    coalgo2(code string) string
    coalgo3(code string) string

    // for code deobfuscate
    Deobfuscate(code string, id ...int)
    Dealgo(id int, code string)
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

func (c *codeObfuscate) coalgo1(code string) string {
    // TODO
    return code
}

func (c *codeObfuscate) coalgo2(code string) string {
    // TODO
    return code
}

func (c *codeObfuscate) coalgo3(code string) string {
    // TODO
    return code
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
                log.Fatalf("wrong encoding number.")
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
