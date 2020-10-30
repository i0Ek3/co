package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
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
	debug  bool
}

// confuser defines code confuse interface
type confuser interface {
	New() *Confuse
	checkStatus(status string) bool
	checkID(id int) bool
	caseTransform(code string, mode ...string)
	isCodeEmpty(code string) (ok bool)

	// for code obfuscate
	processOB(code string, debug bool)
	Obfuscate(code string, debug bool, id ...int)
	coalgo(id int, code string, debug bool) (newdata string)
	coalgo1(code string, debug bool) string
	coalgo2(code string, debug bool) string
	mapCode2Char(code string, len int) []string
	coalgo3(code string, debug bool) string
	coalgo4String(code string, debug bool) string
	parseEncodeIntoFile(code string, algoid int, debug bool) bool
	processFileOB(filename string, algoid int, debug bool) string

	// for code deobfuscate
	processDE(code string, debug bool)
	Deobfuscate(code string, debug bool, id ...int)
	dealgo(id int, code string, debug bool) (newdata string)
	dealgo1(code string, debug bool) string
	dealgo2(code string, debug bool) string
	dealgo3(code string, debug bool) string
	parseDecodeIntoFile(code string, algoid int, debug bool) bool
	processFileDE(filename string, algoid int, debug bool) string
}

// New creates new code confuse instance
func (c *Confuse) New() *Confuse {
	co := &Confuse{
		status: c.status,
		algoed: c.algoed,
		algoid: c.algoid,
		cobit:  c.cobit,
		debug:  c.debug,
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

func (c *Confuse) isCodeEmpty(code string) bool {
	if code == "" {
		return false
	}
	return true
}

func (c *Confuse) processOB(code string, debug bool) (ok bool) {
	if c.isCodeEmpty(code) {
		ok = true
		if c.algoed && c.checkStatus(c.status) {
			switch c.status {
			case OB:
				if c.checkID(c.algoid) {
					c.Obfuscate(code, debug, c.algoid)
				} else {
					log.Fatalf("wrong encoding number.")
				}
			}
		} else {
			switch c.status {
			case OB:
				c.Obfuscate(code, debug)
			}
		}
	}
	return ok == true
}

// Obfuscate obfuscates the code
func (c *Confuse) Obfuscate(code string, debug bool, id ...int) {
	if len(id) > 0 {
		c.algoid = id[0]
		if c.checkID(c.algoid) {
			go func() {
				c.coalgo(c.algoid, code, debug)
			}()
		}
	}
}

// coalgo defines code obfuscation algorithms
func (c *Confuse) coalgo(id int, code string, debug bool) (newdata string) {
	switch id {
	case 1:
		newdata = c.coalgo1(code, debug)
	case 2:
		newdata = c.coalgo2(code, debug)
	case 3:
		newdata = c.coalgo3(code, debug)
	}
	return
}

// coalgo1 encoding the code string with ordinary offset
// transform, which means map alphabet to next n alphabet
// with all lower case.
func (c *Confuse) coalgo1(code string, debug bool) string {
	var encode string
	alphabet := make(map[int]rune)

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
func (c *Confuse) coalgo2(code string, debug bool) string {
	var encode string
	alphabet := make(map[int]rune)

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

func (c *Confuse) mapCode2Char(code string, len int) []string {
	mode := "lower"
	c.caseTransform(code, mode)

	specChar := []string{"_", "-"}
	newer := make([]string, len)

	for i := 0; i < AN; i++ {
		newer[i] = specChar[1] + fmt.Sprint(i/c.cobit) + specChar[0] + fmt.Sprint(i%c.cobit)
	}
	return newer
}

// coalgo3 encoding the code string with ordinary offset
// transform, but mapped with special characters like _.
func (c *Confuse) coalgo3(code string, debug bool) string {
	ret := ""
	newer := c.mapCode2Char(code, len(code))

	for i := range code {
		idx := int(code[i])
		// TODO: parse other special characters rather than letter only
		if 97 <= idx && idx <= 122 {
			n := idx - 97
			ret += newer[n]
		}
	}
	return ret
}

func (c *Confuse) coalgo4String(code string, debug bool) string {
	ret := ""
	newer := c.mapCode2Char(code, AN)

	for i := range code {
		idx := int(code[i])
		n := idx - 97
		ret += newer[n]
	}
	return ret
}

func (c *Confuse) parseEncodeIntoFile(code string, algoid int, debug bool) bool {
	var newdata string
	rand.Seed(1000)
	name := "co" + fmt.Sprint(algoid) + "_" + fmt.Sprint(rand.Intn(100000)) + ".txt"

	newdata = c.coalgo4String(code, debug)
	if err := ioutil.WriteFile(name, []byte(newdata), 0644); err != nil {
		log.Fatalf("file write failed.")
	}
	return true
}

func (c *Confuse) processFileOB(filename string, algoid int, debug bool) string {
	var newdata string

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("file read failed.")
	}

	name := "co" + fmt.Sprint(algoid) + "_" + strings.TrimRight(filename, "go") + "txt"

	newdata = c.coalgo(algoid, string(data), debug)
	if err := ioutil.WriteFile(name, []byte(newdata), 0644); err != nil {
		log.Fatalf("file write failed.")
	}
	return newdata
}

func (c *Confuse) processDE(code string, debug bool) (ok bool) {
	if c.isCodeEmpty(code) {
		ok = true
		if c.algoed && c.checkStatus(c.status) {
			switch c.status {
			case DE:
				if c.checkID(c.algoid) {
					c.Deobfuscate(code, debug, c.algoid)
				} else {
					log.Fatalf("wrong decoding number.")
				}
			}
		} else {
			switch c.status {
			case DE:
				c.Deobfuscate(code, debug)
			}
		}
	}
	return ok == true
}

// Deobfuscate deobfuscates the code
func (c *Confuse) Deobfuscate(code string, debug bool, id ...int) {
	if len(id) > 0 {
		c.algoid = id[0]
		if c.checkID(c.algoid) {
			go func() {
				c.dealgo(c.algoid, code, debug)
			}()
		}
	}
}

// dealgo defines code deobfuscation algorithms
func (c *Confuse) dealgo(id int, code string, debug bool) (newdata string) {
	switch id {
	case 1:
		newdata = c.dealgo1(code, debug)
	case 2:
		newdata = c.dealgo2(code, debug)
	case 3:
		newdata = c.dealgo3(code, debug)
	}
	return
}

func (c *Confuse) dealgo1(code string, debug bool) string {
	var decode string
	alphabet := make(map[int]rune)

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

func (c *Confuse) dealgo2(code string, debug bool) string {
	var decode string
	alphabet := make(map[int]rune)

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

func (c *Confuse) dealgo3(code string, debug bool) string {
	m := map[rune]string{
		'a': "-0_0",
		'b': "-0_1",
		'c': "-0_2",
		'd': "-0_3",
		'e': "-0_4",
		'f': "-0_5",
		'g': "-0_6",
		'h': "-0_7",
		'i': "-1_0",
		'j': "-1_1",
		'k': "-1_2",
		'l': "-1_3",
		'm': "-1_4",
		'n': "-1_5",
		'o': "-1_6",
		'p': "-1_7",
		'q': "-2_0",
		'r': "-2_1",
		's': "-2_2",
		't': "-2_3",
		'u': "-2_4",
		'v': "-2_5",
		'w': "-2_6",
		'x': "-2_7",
		'y': "-3_0",
		'z': "-3_1",
	}

	ret := ""
	res := ""
	sign := 0

	for i := 0; i < len(code); i++ {
		if (i+1)%4 == 0 {
			for j := sign; j < i+1; j++ {
				ret += string(code[j])
			}

			if debug {
				fmt.Printf("ret ----> %v\n", ret)
			}

			for idx, v := range m {
				if ret == v {
					res += string(idx)
					if debug {
						fmt.Printf("idx ----> %v\n", idx)
					}
				}
			}
			if debug {
				fmt.Printf("res ----> %v\n", res)
			}
			sign += 4
			ret = ""
		}
	}
	if debug {
		fmt.Printf("result ----> %v\n", res)
	}
	return res
}

func (c *Confuse) parseDecodeIntoFile(code string, algoid int, debug bool) bool {
	var newdata string
	rand.Seed(1000)
	name := "de" + fmt.Sprint(algoid) + "_" + fmt.Sprint(rand.Intn(100000)) + ".txt"

	newdata = c.dealgo(algoid, code, debug)
	if err := ioutil.WriteFile(name, []byte(newdata), 0644); err != nil {
		log.Fatalf("file write failed.")
	}
	return true
}

func (c *Confuse) processFileDE(filename string, algoid int, debug bool) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("file read failed.")
	}

	name := "de" + fmt.Sprint(algoid) + "_" + strings.TrimRight(filename, "txt") + "go"

	newdata := c.dealgo(algoid, string(data), debug)
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
		debug    bool
	)

	co := &Confuse{input, true, 3, 8, debug}

Again:
	fmt.Printf("Please input OB for obfuscation or DE for deobfuscation: ")
	fmt.Scan(&inputO)
	if inputO != "OB" && inputO != "DE" {
		log.Warnf("none of OB or DE, please input again!")
		goto Again
	}

	switch inputO {
	case "OB":
		fmt.Printf("Which option you want to obfuscate: [code or file]: ")
		fmt.Scan(&input)

		switch input {
		case "code":
			fmt.Printf("Please input the code string: ")
			fmt.Scan(&code)

			fmt.Printf("Would you like enable debug mode? [true or false]: ")
			fmt.Scan(&debug)

			if co.processOB(code, debug) {
				co.parseEncodeIntoFile(code, co.algoid, debug)
				log.Infof("code obfuscated to the file!")
			} else {
				log.Warnf("code cannot obfuscated!")
			}
		case "file":
			fmt.Printf("Please input filename: ")
			fmt.Scan(&filename)

			fmt.Printf("Would you like enable debug mode? [ture or false]: ")
			fmt.Scan(&debug)

			if filename != "" {
				co.processFileOB(filename, co.algoid, debug)
				log.Infof("file obfuscated.")
			} else {
				log.Warnf("invalid filename!")
			}
		}
	case "DE":
		fmt.Printf("Which option you want to deobfuscate: [code or file]: ")
		fmt.Scan(&input)

		switch input {
		case "code":
			fmt.Printf("Please input the code string: ")
			fmt.Scan(&code)

			fmt.Printf("Would you like enable debug mode? [ture or false]: ")
			fmt.Scan(&debug)

			if co.processDE(code, debug) {
				co.parseDecodeIntoFile(code, co.algoid, debug)
				log.Infof("code deobfuscated to the file!")
			} else {
				log.Warnf("code cannot deobfuscated!")
			}
		case "file":
			fmt.Printf("Please input filename: ")
			fmt.Scan(&filename)

			fmt.Printf("Would you like enable debug mode? [ture or false]: ")
			fmt.Scan(&debug)

			if filename != "" {
				co.processFileDE(filename, co.algoid, debug)
				log.Infof("file deobfuscated.")
			} else {
				log.Warnf("invalid filename!")
			}
		}
	default:

	}
}
