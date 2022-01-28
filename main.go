package brainfuck

// #include "bfstdin.h"
import "C"
import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type ByteCellsInterface interface {
	incrementPosition() error
	decrementPosition() error
	consoleInput() error
	consoleOutput() error
	incrementPointer()
	decrementPointer()
	generateStrip() (int, []rune)
	processLoop()
}

type ByteCells struct {
	ByteArray [30_000]rune
	position  int
}

var availCommands = []string{
	"<", //60
	">", //62
	"+", //43
	"-", //45
	".", //46
	",", //44
	"]", //93
	"[", //91
}

func isValidCmd(r rune) bool {
	for _, item := range availCommands {
		if string(r) == item {
			return true
		}
	}
	return false
}

func tidyString(inputStrip string) []rune {
	runes := []rune{}
	for _, item := range inputStrip {
		if isValidCmd(item) {
			runes = append(runes, item)
		}
	}
	return runes
}

func (Bc *ByteCells) consoleInput() error {
	s, err := C.inputPswrd()
	if err != nil {
		return errors.New("an error occured with input: " + err.Error())
	}
	Bc.ByteArray[Bc.position] = rune(s)
	return nil
}

func (Bc *ByteCells) consoleOutput() error {
	fmt.Printf("%s", string(Bc.ByteArray[Bc.position]))
	return nil
}

func (Bc *ByteCells) incrementPosition() error {
	if Bc.position < 30_000 {
		Bc.position += 1
		return nil
	} else if Bc.position == 30_000 {
		Bc.position = 0
		return nil
	}
	return errors.New("an error occured while incrementing the data pointer")
}

func (Bc *ByteCells) decrementPosition() error {
	if Bc.position > 0 {
		Bc.position -= 1
		return nil
	} else if Bc.position == 0 {
		Bc.position = 30_000
	}
	return errors.New("an error occured while decrementing the data pointer")
}

func (Bc *ByteCells) incrementPointer() {
	if Bc.ByteArray[Bc.position] == 255 {
		Bc.ByteArray[Bc.position] = 0
	} else {
		Bc.ByteArray[Bc.position] += 1
	}
}

func (Bc *ByteCells) decrementPointer() {
	if Bc.ByteArray[Bc.position] == 0 {
		Bc.ByteArray[Bc.position] = 255
	} else {
		Bc.ByteArray[Bc.position] -= 1
	}
}

func (Bc *ByteCells) processLoop(cmds []rune) {
	ind := -1
	for ind < len(cmds)-1 {
		ind += 1
		item := cmds[ind]
		switch item {
		case '>':
			Bc.incrementPosition()
			continue
		case '<':
			Bc.decrementPosition()
			continue
		case '+':
			Bc.incrementPointer()
			continue
		case '-':
			Bc.decrementPointer()
			continue
		case '.':
			Bc.consoleOutput()
			continue
		case ',':
			Bc.consoleInput()
			continue
		case '[':
			if Bc.ByteArray[Bc.position] == 0 {
				i := ind + 1
				for cmds[i] != 93 {
					i += 1
				}
				ind = i
			}
			continue
		case ']':
			if Bc.ByteArray[Bc.position] != 0 {
				i := ind
				for cmds[i] != 91 {
					i -= 1
				}
				ind = i
			}
			continue
		}
	}
}
func initByteCells() ByteCells {
	return ByteCells{[30_000]rune{}, 0}
}

func runCode(Bc ByteCells, s string) {
	Bc.processLoop(tidyString(string(s)))
}

func runFromFile(fpath string) error {
	if path.Ext(fpath) == ".bf" {
		f, err := ioutil.ReadFile(fpath)
		if err != nil {
			return errors.New("an error occured while reading the file: " + err.Error())
		}
		Bc := ByteCells{[30_000]rune{}, 0}
		runCode(Bc, string(f))
	}
	return nil
}

func main() {
	if os.Args[1] != "" {
		err := runFromFile(os.Args[1])
		panic(errors.New("an error occured while opening the given file path: " + err.Error()))
	} else {
		panic(errors.New("an error occured while opening the file path: " + os.Args[1]))
	}
}
