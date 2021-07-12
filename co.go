package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"unicode"
)

// code confuse operations
const (
	OB = "obfuscate"
	DE = "deobfuscate"
	NO = ""

    // alphabets number
	AN = 26 
)

// Confuse defines code confuse fileds
type Confuse struct {
	status string
	algoed bool
	algoid int
	cobit  int
	debug  string
}

// Obfuscation defines code confuse interface
type Obfuscation interface {
	checkStatus(status string) bool
	checkID(id int) bool
	caseTransform(code string, mode ...string)
	isCodeEmpty(code string) (ok bool)

	// for code obfuscate
	processOB(code string, debug string)
	Obfuscate(code string, debug string, id ...int)
	coalgo(id int, code string, debug string) (newdata string)
	coalgo1(code string, debug string) string
	coalgo2(code string, debug string) string
	mapCode2Char(code string, len int) []string
	coalgo3(code string, debug string) string
	coalgo4(code string, debug string) string
	parseEncodeIntoFile(code string, algoid int, debug string) bool
	processFileOB(filename string, algoid int, debug string) string

	// for code deobfuscate
	processDE(code string, debug string)
	Deobfuscate(code string, debug string, id ...int)
	dealgo(id int, code string, debug string) (newdata string)
	dealgo1(code string, debug string) string
	dealgo2(code string, debug string) string
	dealgo3(code string, debug string) string
	parseDecodeIntoFile(code string, algoid int, debug string) bool
	processFileDE(filename string, algoid int, debug string) string
}

// New creates a pointer type Confuse
func NewConfuse(defa bool, id int, bit int, debug string) *Confuse {
	co := &Confuse{
		algoed: defa,
		algoid: id,
		cobit:  bit,
		debug:  debug,
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

func (c *Confuse) processOB(code string, debug string) (ok bool) {
	if c.isCodeEmpty(code) {
		ok = true
		if c.algoed && c.checkStatus(c.status) {
			switch c.status {
			case OB:
				if c.checkID(c.algoid) {
					c.Obfuscate(code, debug, c.algoid)
				} else {
					showMsg("wrong encoding number.")
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
func (c *Confuse) Obfuscate(code string, debug string, id ...int) {
	if len(id) > 0 {
		c.algoid = id[0]
		if c.checkID(c.algoid) {
            // TODO: add channel to transport data
			go func(i int) {
				//c.coalgo(c.algoid, code, debug)
				c.coalgo(i, code, debug)
			}(c.algoid)
		}
	}
}

// coalgo defines code obfuscation algorithms
func (c *Confuse) coalgo(id int, code string, debug string) (newdata string) {
	switch id {
	case 1:
		newdata = c.coalgo1(code, debug)
	case 2:
		newdata = c.coalgo2(code, debug)
	case 3:
		newdata = c.coalgo3(code, debug)
	case 4:
		newdata = c.coalgo4(code, debug)
	}
	return
}

// coalgo1 encoding the code string with ordinary offset
// transform, which means map alphabet to next n alphabet
// with all lower case.
func (c *Confuse) coalgo1(code string, debug string) string {
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
func (c *Confuse) coalgo2(code string, debug string) string {
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
func (c *Confuse) coalgo3(code string, debug string) string {
	ret := ""
	newer := c.mapCode2Char(code, len(code))

	for i := range code {
		idx := int(code[i])
		if debug == "true" {
			fmt.Printf("idx ----> %v\n", idx)
		}

		// TODO: parse other special characters rather than letter only
		if 97 <= idx && idx <= 122 {
			n := idx - 97
			ret += newer[n]
			if debug == "true" {
				fmt.Printf("ret ----> %v\n", ret)
			}
		}
	}
	return ret
}

func (c *Confuse) coalgo4(code string, debug string) string {
	ret := ""
	newer := c.mapCode2Char(code, AN)

	for i := range code {
		idx := int(code[i])
		if debug == "true" {
			fmt.Printf("idx ----> %v\n", idx)
		}
		n := idx - 97
		ret += newer[n]
		if debug == "true" {
			fmt.Printf("ret ----> %v\n", ret)
		}
	}
	return ret
}

func (c *Confuse) parseEncodeIntoFile(code string, algoid int, debug string) bool {
	var newdata string
	rand.Seed(1000)
	name := "co" + fmt.Sprint(algoid) + "_" + fmt.Sprint(rand.Intn(100000)) + ".txt"

	// TODO
	//newdata = c.coalgo(algoid, code, debug)
	newdata = c.coalgo4(code, debug)

	if err := ioutil.WriteFile(name, []byte(newdata), 0644); err != nil {
		showMsg("file write failed.")
	}
    fmt.Println(newdata) 
	return true
}

func (c *Confuse) processFileOB(filename string, algoid int, debug string) string {
	var newdata string

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		showMsg("file read failed.")
	}

	name := "co" + fmt.Sprint(algoid) + "_" + strings.TrimRight(filename, "go") + "txt"

	newdata = c.coalgo(algoid, string(data), debug)
	if err := ioutil.WriteFile(name, []byte(newdata), 0644); err != nil {
		showMsg("file write failed.")
	}
	return newdata
}

func (c *Confuse) processDE(code string, debug string) (ok bool) {
	if c.isCodeEmpty(code) {
		ok = true
		if c.algoed && c.checkStatus(c.status) {
			switch c.status {
			case DE:
				if c.checkID(c.algoid) {
					c.Deobfuscate(code, debug, c.algoid)
				} else {
					showMsg("wrong decoding number.")
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
func (c *Confuse) Deobfuscate(code string, debug string, id ...int) {
	if len(id) > 0 {
		c.algoid = id[0]
		if c.checkID(c.algoid) {
            // TODO
			go func(i int) {
				//c.dealgo(c.algoid, code, debug)
				c.dealgo(i, code, debug)
			}(c.algoid)
		}
	}
}

// dealgo defines code deobfuscation algorithms
func (c *Confuse) dealgo(id int, code string, debug string) (newdata string) {
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

func (c *Confuse) dealgo1(code string, debug string) string {
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

func (c *Confuse) dealgo2(code string, debug string) string {
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

func (c *Confuse) dealgo3(code string, debug string) string {
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

			if debug == "true" {
				fmt.Printf("ret ----> %v\n", ret)
			}

			for idx, v := range m {
				if ret == v {
					res += string(idx)
					if debug == "true" {
						fmt.Printf("idx ----> %v\n", idx)
					}
				}
			}
			if debug == "true" {
				fmt.Printf("res ----> %v\n", res)
			}
			sign += 4
			ret = ""
		}
	}
	if debug == "true" {
		fmt.Printf("result ----> %v\n", res)
	}
	return res
}

func (c *Confuse) parseDecodeIntoFile(code string, algoid int, debug string) bool {
	var newdata string
	rand.Seed(1000)
	name := "de" + fmt.Sprint(algoid) + "_" + fmt.Sprint(rand.Intn(100000)) + ".txt"

	newdata = c.dealgo(algoid, code, debug)
	if err := ioutil.WriteFile(name, []byte(newdata), 0644); err != nil {
		showMsg("file write failed.")
	}
    fmt.Println(newdata) 
	return true
}

func (c *Confuse) processFileDE(filename string, algoid int, debug string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		showMsg("file read failed.")
	}

	name := "de" + fmt.Sprint(algoid) + "_" + strings.TrimRight(filename, "txt") + "go"

	newdata := c.dealgo(algoid, string(data), debug)
	if err := ioutil.WriteFile(name, []byte(newdata), 0644); err != nil {
		showMsg("file write failed.")
	}
	return newdata
}

func showMsg(msg string) {
	fmt.Printf("\tInfo -----> %v\n", msg)
}

func doLoop(input, res1, res2, msg string) bool {
	if input == res1 || input == res2 {
		return false
	} else {
		showMsg(msg)
		return true
	}
}

func runIO(input *string, output string) {
	fmt.Printf(output)
	fmt.Scan(input)
}

func main() {
	var (
		input    string
		inputO   string
		filename string
		code     string
		debug    string
	)

	const (
		OBSTR = "OB"
		DESTR = "DE"

		OUTPUT = "Which option do you want to operate? [OB or DE]: "
		OBMSG  = "None of them, please input again!"

		CODE = "code"
		FILE = "file"

		T = "true"
		F = "false"

		OBSRC = "Which option do you want to obfuscate? [code or file]: "
		DESRC = "Which option do you want to deobfuscate? [code or file]: "

		CODEIN = "Please input the code string: "
		DEBUG  = "Would you like to enable the debug mode? [true or false]: "
		FNSTR  = "Please input the filename: "
	)

	// default algorithm is coalgo3() and dealgo3(),
	// and debug mode disable
	co := NewConfuse(true, 3, 8, F)

A1:
	runIO(&inputO, OUTPUT)
	if doLoop(inputO, OBSTR, DESTR, OBMSG) {
		goto A1
	}

	switch inputO {
	case OBSTR:
	A2:
		runIO(&input, OBSRC)
		if doLoop(input, CODE, FILE, OBMSG) {
			goto A2
		}

		switch input {
		case CODE:
			runIO(&code, CODEIN)
		A3:
			runIO(&debug, DEBUG)
			if doLoop(debug, T, F, OBMSG) {
				goto A3
			}

			if co.processOB(code, co.debug) {
				co.parseEncodeIntoFile(code, co.algoid, debug)
				showMsg("code obfuscated to the file!")
			} else {
				showMsg("code cannot obfuscated!")
			}
		case FILE:
			runIO(&filename, FNSTR)
		A4:
			runIO(&debug, DEBUG)
			if doLoop(debug, T, F, OBMSG) {
				goto A4
			}

			if filename != "" {
				co.processFileOB(filename, co.algoid, debug)
				showMsg("file obfuscated.")
			} else {
				showMsg("invalid filename!")
			}
		}
	case DESTR:
	A5:
		runIO(&input, DESRC)
		if doLoop(input, CODE, FILE, OBMSG) {
			goto A5
		}

		switch input {
		case CODE:
			runIO(&code, CODEIN)
		A6:
			runIO(&debug, DEBUG)
			if doLoop(debug, T, F, OBMSG) {
				goto A6
			}

			if co.processDE(code, co.debug) {
				co.parseDecodeIntoFile(code, co.algoid, debug)
				showMsg("code deobfuscated to the file!")
			} else {
				showMsg("code cannot deobfuscated!")
			}
		case FILE:
			runIO(&filename, FNSTR)
		A7:
			runIO(&debug, DEBUG)
			if doLoop(debug, T, F, OBMSG) {
				goto A7
			}

			if filename != "" {
				co.processFileDE(filename, co.algoid, debug)
				showMsg("file deobfuscated.")
			} else {
				showMsg("invalid filename!")
			}
		}
	default:
	}
}
