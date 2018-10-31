package vterm

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"unicode/utf8"
)

// Attributes
const (
	Bold      = 1
	Dim       = 2
	Underline = 4
	Blink     = 5
	Reverse   = 7
	Hidden    = 8
	None      = 0
)

// Foreground Colors
const (
	Default = 39
	Black   = iota + 29
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	LightGray
)

const (
	DarkGray = iota + 90
	LightRed
	LightGreen
	LightYellow
	LightBlue
	LightMagenta
	LightCyan
	White
)

// Background Colors
const (
	BgDefault = 49
	BgBlack   = iota + 39
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgLightGray
)

const (
	BgDarkGray = iota + 100
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgWhite
)

func Colorize(arg interface{}, attribute, color int) string {
	return fmt.Sprintf("\x1b[%d;%dm%v\x1b[0m", attribute, color, arg)
}

func PipeToLess(str string) error {
	less := exec.Command("less", "-r")
	less.Stdout = os.Stdout
	less.Stderr = os.Stderr
	w, err := less.StdinPipe()
	if err != nil {
		return err
	}
	less.Start()
	r := strings.NewReader(str)
	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}
	w.Close()
	return less.Wait()
}

func Reflow(text string, width, indent int) string {
	text = strings.Replace(text, "\t", " ", -1)
	text = strings.Replace(text, "\r", " ", -1)
	text = strings.Replace(text, "\n", " ", -1)
	tokens := strings.Split(text, " ")

	newLines := []string{}
	currentLine := []string{}
	lineLen := indent
	indentStr := strings.Repeat(" ", indent)
	for _, t := range tokens {
		length := utf8.RuneCountInString(t)
		if lineLen+length > width {
			newLine := indentStr + strings.Join(currentLine, " ")
			newLines = append(newLines, newLine)
			currentLine = []string{t}
			lineLen = length + indent
		} else {
			currentLine = append(currentLine, t)
			lineLen += length + 1
		}
	}
	lastLine := indentStr + strings.Join(currentLine, " ")
	return strings.Join(append(newLines, lastLine), "\n")
}
