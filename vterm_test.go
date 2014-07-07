package vterm

import (
    "fmt"
    "testing"
)

func TestForegroundColors(t *testing.T) {
    fmt.Printf("%s\n", Colorize("Default", 0, Default))
    fmt.Printf("%s\n", Colorize("Black", 0, Black))
    fmt.Printf("%s\n", Colorize("Red", 0, Red))
    fmt.Printf("%s\n", Colorize("Green", 0, Green))
    fmt.Printf("%s\n", Colorize("Yellow", 0, Yellow))
    fmt.Printf("%s\n", Colorize("Blue", 0, Blue))
    fmt.Printf("%s\n", Colorize("Magenta", 0, Magenta))
    fmt.Printf("%s\n", Colorize("Cyan", 0, Cyan))
    fmt.Printf("%s\n", Colorize("LightGray", 0, LightGray))
    fmt.Printf("%s\n", Colorize("DarkGray", 0, DarkGray))
    fmt.Printf("%s\n", Colorize("LightRed", 0, LightRed))
    fmt.Printf("%s\n", Colorize("LightGreen", 0, LightGreen))
    fmt.Printf("%s\n", Colorize("LightYellow", 0, LightYellow))
    fmt.Printf("%s\n", Colorize("LightBlue", 0, LightBlue))
    fmt.Printf("%s\n", Colorize("LightMagenta", 0, LightMagenta))
    fmt.Printf("%s\n", Colorize("LightCyan", 0, LightCyan))
    fmt.Printf("%s\n", Colorize("White", 0, White))
}

func TestBackgroundColors(t *testing.T) {
    fmt.Printf("%s\n", Colorize("Default", 0, BgDefault))
    fmt.Printf("%s\n", Colorize("Black", 0, BgBlack))
    fmt.Printf("%s\n", Colorize("Red", 0, BgRed))
    fmt.Printf("%s\n", Colorize("Green", 0, BgGreen))
    fmt.Printf("%s\n", Colorize("Yellow", 0, BgYellow))
    fmt.Printf("%s\n", Colorize("Blue", 0, BgBlue))
    fmt.Printf("%s\n", Colorize("Magenta", 0, BgMagenta))
    fmt.Printf("%s\n", Colorize("Cyan", 0, BgCyan))
    fmt.Printf("%s\n", Colorize("LightGray", 0, BgLightGray))
    fmt.Printf("%s\n", Colorize("DarkGray", 0, BgDarkGray))
    fmt.Printf("%s\n", Colorize("LightRed", 0, BgLightRed))
    fmt.Printf("%s\n", Colorize("LightGreen", 0, BgLightGreen))
    fmt.Printf("%s\n", Colorize("LightYellow", 0, BgLightYellow))
    fmt.Printf("%s\n", Colorize("LightBlue", 0, BgLightBlue))
    fmt.Printf("%s\n", Colorize("LightMagenta", 0, BgLightMagenta))
    fmt.Printf("%s\n", Colorize("LightCyan", 0, BgLightCyan))
    fmt.Printf("%s\n", Colorize("White", 0, BgWhite))
}
