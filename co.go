package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	log "github.com/sirupsen/logrus"
)

// code confuse operations
const (
	OB = "obfuscate"
	DE = "deobfuscate"
	NO = ""
)

// Confuse defines code confuse fileds
type Confuse struct {
	status string
	algoed bool
	algoid int
	cobit  int
}

// confuser defines code confuse interface
type confuser interface {
	New() *Confuse
	checkStatus(status string) bool
	checkID(id int) bool
	caseTransform(code string)
	isCodeEmpty(code string) bool

	// for code obfuscate
	processOB(code string)
	Obfuscate(code string, id ...int)
	Coalgo(id int, code []string)
	coalgo1(code string) string
	coalgo2(code string) string
	coalgo3(code string) []string
	processFileOB(f *os.File) *os.File

	// for code deobfuscate
	processDE(code string)
	Deobfuscate(code string, id ...int)
	Dealgo(id int, code []string)
	dealgo1(code string) string
	dealgo2(code string) string
	dealgo3(code string) string
	processFileDE(f *os.File) *os.File
}

// New creates new code confuse instance
func (c *Confuse) New() *Confuse {
	co := &Confuse{
		status: c.status,
		algoed: c.algoed,
		algoid: c.algoid,
		cobit:  c.cobit,
	}
	return co
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

func (c *Confuse) caseTransform(code string) {
	var alphabetu map[int]rune
	var alphabetl map[int]rune

	upper := 0
	lower := 0

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
}

func (c *Confuse) isCodeEmpty(code string) bool {
	switch {
	case code == "":
		return false
	default:
		return true
	}
}

func (c *Confuse) processOB(code string) {
	if c.isCodeEmpty(code) {
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
}

// Obfuscate obfuscates the code
func (c *Confuse) Obfuscate(code string, id ...int) {
	if len(id) > 0 {
		c.algoid = id[0]
		if c.checkID(c.algoid) {
			c.Coalgo(c.algoid, code)
		}
	}
}

// Coalgo defines code obfuscation algorithms
func (c *Confuse) Coalgo(id int, code string) {
	switch id {
	case 1:
		c.coalgo1(code)
	case 2:
		c.coalgo2(code)
	case 3:
		c.coalgo3(code)
	}
}

// TODO: refactor coalgo1 and coalgo2, maybe incorporate them into one.
// coalgo1 encoding the code string with ordinary offset
// transform, which means map alphabet to next n alphabet
// with all lower case.
func (c *Confuse) coalgo1(code string) string {
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
func (c *Confuse) coalgo2(code string) string {
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

// coalgo3 encoding the code string with ordinary offset
// transform, but mapped with special characters like _.
func (c *Confuse) coalgo3(code string) []string {
	specChar := []string{"_", "-"}
	c.caseTransform(code)

	// this flow converts char to _ or - in code. If we only
	// use _ and - without other identifiers then we cannot
	// recognize what exactly alphabet is represent, so we
	// need some different identifiers to distinguish the code,
	// just like _ represents a, then 2_ represents b instead
	// of __.
	var newer []string
	for i := 0; i < 26; i++ {
		if i < c.cobit {
			newer[i] = specChar[0] + string(i+1)
		}
		newer[i] = specChar[1] + string(i%c.cobit) + specChar[0] + string(i+1)
	}

	var ret []string
	for k, _ := range code {
		n := (int)(code[k]) - 48 - 97
		ret = append(ret, newer[n])
	}
	return ret
}

func (c *Confuse) processFileOB(f *os.File) *os.File {
	// TODO
	return f
}

func (c *Confuse) processDE(code string) {
	if c.isCodeEmpty(code) {
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
}

// Deobfuscate deobfuscates the code
func (c *Confuse) Deobfuscate(code string, id ...int) {
	if len(id) > 0 {
		c.algoid = id[0]
		if c.checkID(c.algoid) {
			c.Dealgo(c.algoid, code)
		}
	}
}

// Dealgo defines code deobfuscation algorithms
func (c *Confuse) Dealgo(id int, code string) {
	switch id {
	case 1:
		c.dealgo1(code)
	case 2:
		c.dealgo2(code)
	case 3:
		c.dealgo3(code)
	}
}

func (c *Confuse) dealgo1(code string) string {
	var decode string
	var alphabet map[int]rune

	for i := 0; i < 26; i++ {
		for j := 'i'; j <= 'r'; j++ {
			alphabet[i] = j
		}
		for k := 'a'; k <= 'h'; k++ {
			alphabet[i] = k
		}
	}

	for i, _ := range code {
		num := (int)(code[i]) + 48 + 97 - c.cobit + 26
		if num > 26 {
			decode = string(alphabet[num-26])
		}
		decode = string(alphabet[num])
	}

	return decode
}

func (c *Confuse) dealgo2(code string) string {
	var decode string
	var alphabet map[int]rune

	for i := 0; i < 26; i++ {
		for j := 'I'; j <= 'R'; j++ {
			alphabet[i] = j
		}
		for k := 'A'; k <= 'H'; k++ {
			alphabet[i] = k
		}
	}

	for i, _ := range code {
		num := (int)(code[i]) + 48 + 65 - c.cobit + 26
		if num > 26 {
			decode = string(alphabet[num-26])
		}
		decode = string(alphabet[num])
	}

	return decode
}

func (c *Confuse) dealgo3(code string) string {
	// TODO
	return code
}

func (c *Confuse) processFileDE(f *os.File) *os.File {
	// TODO
	return f
}

func main() {
	var input string
	var code string
Again:
	fmt.Printf("Please input OB for obfuscation or DE for deobfuscation: ")
	fmt.Scan(&input)

	if input != "OB" && input != "DE" {
		log.Warnf("None of OB or DE, please input again!")
		goto Again
	}
	co := &Confuse{input, true, 3, 8}

	fmt.Printf("Please input the code string: ")
	fmt.Scan(&code)

	if input == "OB" {
		co.processOB(code)
	}
	co.processDE(code)
}
