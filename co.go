package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"

	log "github.com/sirupsen/logrus"
)

// code confuse operations
const (
	OB = "obfuscate"
	DE = "deobfuscate"
	NO = ""

	AN = 26 // alphabets number
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
	caseTransform(code string, mode ...string)
	isCodeEmpty(code string) (ok bool)

	// for code obfuscate
	processOB(code string)
	Obfuscate(code string, id ...int)
	coalgo(id int, code string) (newdata string)
	coalgo1(code string) string
	coalgo2(code string) string
	coalgo3(code string) string
	processFileOB(filename string, algoid int) string

	// for code deobfuscate
	processDE(code string)
	Deobfuscate(code string, id ...int)
	dealgo(id int, code string) (newdata string)
	dealgo1(code string) string
	dealgo2(code string) string
	dealgo3(code string) string
	processFileDE(filename string, algoid int) string
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

func (c *Confuse) caseTransform(code string, mode ...string) {
	alphabetu := make(map[int]rune, AN)
	alphabetl := make(map[int]rune, AN)

	upper := 0
	lower := 0
	if len(mode) == 0 {
		for i, _ := range code {
			if unicode.IsUpper((rune)(code[i])) {
				upper++
			} else if unicode.IsLower((rune)(code[i])) {
				lower++
			}
		}

		if upper > lower {
			strings.ToUpper(code)
			for i := 0; i < AN; i++ {
				for j := 'A'; j <= 'Z'; j++ {
					alphabetu[i] = j
				}
			}
		} else {
			strings.ToLower(code)
			for i := 0; i < AN; i++ {
				for j := 'a'; j <= 'z'; j++ {
					alphabetl[i] = j
				}
			}
		}
	} else {
		if mode[0] == "lower" {
			strings.ToLower(code)
			for i := 0; i < AN; i++ {
				for j := 'a'; j <= 'z'; j++ {
					alphabetl[i] = j
				}
			}
		} else if mode[0] == "upper" {
			strings.ToUpper(code)
			for i := 0; i < AN; i++ {
				for j := 'A'; j <= 'Z'; j++ {
					alphabetu[i] = j
				}
			}
		}
	}
}

// FIXME
func (c *Confuse) isCodeEmpty(code string) bool {
	if code == "" || code == "\n" {
		return false
	}
	return true
}

func (c *Confuse) processOB(code string) (ok bool) {
	if c.isCodeEmpty(code) {
		ok = true
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
	return ok == true
}

// Obfuscate obfuscates the code
func (c *Confuse) Obfuscate(code string, id ...int) {
	if len(id) > 0 {
		c.algoid = id[0]
		if c.checkID(c.algoid) {
			c.coalgo(c.algoid, code)
		}
	}
}

// coalgo defines code obfuscation algorithms
func (c *Confuse) coalgo(id int, code string) (newdata string) {
	switch id {
	case 1:
		newdata = c.coalgo1(code)
	case 2:
		newdata = c.coalgo2(code)
	case 3:
		newdata = c.coalgo3(code)
	}
	return
}

// coalgo1 encoding the code string with ordinary offset
// transform, which means map alphabet to next n alphabet
// with all lower case.
func (c *Confuse) coalgo1(code string) string {
	var encode string
	var alphabet map[int]rune

	encode = strings.ToLower(code)

	for i := 0; i < AN; i++ {
		for j := 'a'; j <= 'z'; j++ {
			alphabet[i] = j
		}
	}

	for i, _ := range code {
		num := (int)(code[i]) - 48 - 97 + c.cobit - AN
		if num > 0 {
			encode = string(alphabet[num])
		}
		encode = string(alphabet[num+AN])
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

	for i := 0; i < AN; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			alphabet[i] = j
		}
	}

	for i, _ := range code {
		num := (int)(code[i]) - 48 - 65 + c.cobit - AN
		if num > 0 {
			encode = string(alphabet[num])
		}
		encode = string(alphabet[num+AN])
	}

	return encode
}

// coalgo3 encoding the code string with ordinary offset
// transform, but mapped with special characters like _.
func (c *Confuse) coalgo3(code string) string {
	specChar := []string{"_", "-"}
	mode := "lower"
	c.caseTransform(code, mode)

	// this flow converts char to _ or - in code. If we only
	// use _ and - without other identifiers then we cannot
	// recognize what exactly alphabet is represent, so we
	// need some different identifiers to distinguish the code,
	// just like _ represents a, then 2_ represents b instead
	// of __.
	newer := make([]string, AN)
	for i := 0; i < AN; i++ {
		if i < c.cobit {
			newer[i] = specChar[0] + fmt.Sprint(i+1)
		}
		newer[i] = specChar[1] + fmt.Sprint(i%c.cobit) + specChar[0] + fmt.Sprint(i+1)
	}

	var ret string
	for _, v := range code {
		// FIXME: Normally, after coversion, all letter will be lower.
		// And h become 104, 104-97+1=8. But here always indicate n is
		// out range of index, that's so weird.
		//n := (int)(code[k]) - 48 - 97
		n := int(v) - 97 + 1
		ret += newer[n]
	}
	return ret
}

func (c *Confuse) processFileOB(filename string, algoid int) string {
	var newdata string
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("file read failed.")
	}

	name := "co" + fmt.Sprint(algoid) + "_" + filename
	newdata = c.coalgo(algoid, string(data))

	if err := ioutil.WriteFile(name, []byte(newdata), 0644); err != nil {
		log.Fatalf("file write failed.")
	}
	return newdata
}

func (c *Confuse) processDE(code string) (ok bool) {
	if c.isCodeEmpty(code) {
		ok = true
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
	return ok == true
}

// Deobfuscate deobfuscates the code
func (c *Confuse) Deobfuscate(code string, id ...int) {
	if len(id) > 0 {
		c.algoid = id[0]
		if c.checkID(c.algoid) {
			c.dealgo(c.algoid, code)
		}
	}
}

// dealgo defines code deobfuscation algorithms
func (c *Confuse) dealgo(id int, code string) (newdata string) {
	switch id {
	case 1:
		newdata = c.dealgo1(code)
	case 2:
		newdata = c.dealgo2(code)
	case 3:
		newdata = c.dealgo3(code)
	}
	return
}

func (c *Confuse) dealgo1(code string) string {
	var decode string
	var alphabet map[int]rune

	for i := 0; i < AN; i++ {
		for j := 'i'; j <= 'r'; j++ {
			alphabet[i] = j
		}
		for k := 'a'; k <= 'h'; k++ {
			alphabet[i] = k
		}
	}

	for i, _ := range code {
		num := (int)(code[i]) + 48 + 97 - c.cobit + AN
		if num > AN {
			decode = string(alphabet[num-AN])
		}
		decode = string(alphabet[num])
	}

	return decode
}

func (c *Confuse) dealgo2(code string) string {
	var decode string
	var alphabet map[int]rune

	for i := 0; i < AN; i++ {
		for j := 'I'; j <= 'R'; j++ {
			alphabet[i] = j
		}
		for k := 'A'; k <= 'H'; k++ {
			alphabet[i] = k
		}
	}

	for i, _ := range code {
		num := (int)(code[i]) + 48 + 65 - c.cobit + AN
		if num > AN {
			decode = string(alphabet[num-AN])
		}
		decode = string(alphabet[num])
	}

	return decode
}

func (c *Confuse) dealgo3(code string) string {
	// TODO
	return code
}

func (c *Confuse) processFileDE(filename string, algoid int) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("file read failed.")
	}
	name := "de" + fmt.Sprint(algoid) + "_" + filename
	newdata := c.dealgo(algoid, string(data))
	if err := ioutil.WriteFile(name, []byte(newdata), 0644); err != nil {
		log.Fatalf("file write failed.")
	}
	return newdata
}

func main() {
	var (
		input    string
		inputO   string
		filename string
		code     string
	)

	co := &Confuse{input, true, 3, 8}

Again:
	fmt.Printf("Please input OB for obfuscation or DE for deobfuscation: ")
	fmt.Scan(&inputO)
	if inputO != "OB" && inputO != "DE" {
		log.Warnf("none of OB or DE, please input again!")
		goto Again
	}

	switch inputO {
	case "OB":
		fmt.Printf("Which option you want to obfuscate/deobfuscate: [code or file]: ")
		fmt.Scan(&input)

		switch input {
		case "code":
			fmt.Printf("Please input the code string: ")
			fmt.Scan(&code)

			if co.processOB(code) {
				fmt.Println("Code obfuscated!")
			} else {
				log.Warnf("code cannot obfuscated!")
			}
		case "file":
			fmt.Printf("Please input filename: ")
			fmt.Scan(&filename)

			if filename != "" {
				co.processFileOB(filename, co.algoid)
				log.Infof("file obfuscated.")
			} else {
				log.Warnf("invalid filename!")
			}
		}
	case "DE":
		fmt.Printf("Which option you want to obfuscate/deobfuscate: [code or file]: ")
		fmt.Scan(&input)

		switch input {
		case "code":
			fmt.Printf("Please input the code string: ")
			fmt.Scan(&code)

			if co.processDE(code) {
				fmt.Println("Code deobfuscated!")
			} else {
				log.Warnf("code cannot deobfuscated!")
			}
		case "file":
			fmt.Printf("Please input filename: ")
			fmt.Scan(&filename)

			if filename != "" {
				co.processFileDE(filename, co.algoid)
				log.Infof("file deobfuscated.")
			} else {
				log.Warnf("invalid filename!")
			}
		}
	default:

	}
}
