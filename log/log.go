package log

import (
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

var SILENT = true
var DEBUG bool
var NO_COLOR bool

func Trace(s ...string) {

	if !SILENT {

		logger := log.New(os.Stderr, "", 0)
		logger.Println(strings.Join(s, " "))
	}
}

func Debug(s ...string) {

	color.NoColor = NO_COLOR

	if DEBUG {

		logger := log.New(os.Stderr, "", 0)
		coloredPrefix := color.New(color.FgGreen).SprintFunc()("[DBG] ")

		logger.SetPrefix(coloredPrefix)
		logger.Println(strings.Join(s, " "))
	}
}

func Info(s ...string) {

	color.NoColor = NO_COLOR

	if !SILENT {

		logger := log.New(os.Stderr, "", 0)
		coloredPrefix := color.New(color.FgBlue).SprintFunc()("[INF] ")

		logger.SetPrefix(coloredPrefix)
		logger.Println(strings.Join(s, " "))
	}
}

func Warn(s ...string) {

	color.NoColor = NO_COLOR

	if !SILENT {

		logger := log.New(os.Stderr, "", 0)
		coloredPrefix := color.New(color.FgYellow).SprintFunc()("[WRN] ")

		logger.SetPrefix(coloredPrefix)
		logger.Println(strings.Join(s, " "))
	}
}

func Error(err error) {

	color.NoColor = NO_COLOR

	if !SILENT && err != nil {

		logger := log.New(os.Stderr, "", 0)
		coloredPrefix := color.New(color.FgRed).SprintFunc()("[ERR] ")

		logger.SetPrefix(coloredPrefix)
		logger.Println(err)
	}
}

func Fatal(err error) {

	color.NoColor = NO_COLOR

	if err != nil {

		logger := log.New(os.Stderr, "", 0)
		coloredPrefix := color.New(color.FgMagenta).SprintFunc()("[FTL] ")

		logger.SetPrefix(coloredPrefix)
		logger.Println(err)

		os.Exit(0)
	}
}
