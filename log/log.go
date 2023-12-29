package log

import (
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

var SILENT bool

func Trace(s ...string) {

	if !SILENT {

		logger := log.New(os.Stderr, "", 0)
		logger.Println(strings.Join(s, " "))
	}
}

func Info(s ...string) {

	if !SILENT {

		logger := log.New(os.Stderr, "", 0)
		coloredPrefix := color.New(color.FgBlue).SprintFunc()("[INF] ")

		logger.SetPrefix(coloredPrefix)
		logger.Println(strings.Join(s, " "))
	}
}

func Warn(s ...string) {

	if !SILENT {

		logger := log.New(os.Stderr, "", 0)
		coloredPrefix := color.New(color.FgYellow).SprintFunc()("[WRN] ")

		logger.SetPrefix(coloredPrefix)
		logger.Println(strings.Join(s, " "))
	}
}

func Error(s ...string) {

	if !SILENT {

		logger := log.New(os.Stderr, "", 0)
		coloredPrefix := color.New(color.FgRed).SprintFunc()("[ERR] ")

		logger.SetPrefix(coloredPrefix)
		logger.Println(strings.Join(s, " "))
	}
}
