package vt

import (
    "fmt"
    "io"
    "os"
    "os/exec"
    "strings"
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
    text = strings.Replace(text, "\r\n", "\n", -1)
    oldLines := strings.Split(text, "\n")
    tokens := []string{}
    for _, l := range oldLines {
        tokens = append(tokens, strings.Split(l, " ")...)
    }
    newLines := []string{}
    currentLine := []string{}
    lineLen := indent
    indentStr := strings.Repeat(" ", indent)
    for _, t := range tokens {
        if lineLen + len(t) > width {
            newLine := indentStr + strings.Join(currentLine, " ")
            newLines = append(newLines, newLine)
            currentLine = []string{t}
            lineLen = len(t) + indent
        } else {
            currentLine = append(currentLine, t)
            lineLen += len(t) + 1
        }
    }
    finalLine := indentStr + strings.Join(currentLine, " ")
    return strings.Join(append(newLines, finalLine), "\n")
}
