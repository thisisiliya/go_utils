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
var OUT = os.Stderr

var DBG_COLOR = color.FgGreen
var INF_COLOR = color.FgBlue
var WRN_COLOR = color.FgYellow
var ERR_COLOR = color.FgRed
var FTL_COLOR = color.FgMagenta

func Trace(s ...string) {

	if !SILENT {

		logger := log.New(OUT, "", 0)
		logger.Println(strings.Join(s, " "))
	}
}

func Debug(s ...string) {

	color.NoColor = NO_COLOR

	if DEBUG {

		logger := log.New(OUT, "", 0)
		coloredPrefix := "[" + color.New(DBG_COLOR).SprintFunc()("DBG") + "] "

		logger.SetPrefix(coloredPrefix)
		logger.Println(strings.Join(s, " "))
	}
}

func Info(s ...string) {

	color.NoColor = NO_COLOR

	if !SILENT {

		logger := log.New(OUT, "", 0)
		coloredPrefix := "[" + color.New(INF_COLOR).SprintFunc()("INF") + "] "

		logger.SetPrefix(coloredPrefix)
		logger.Println(strings.Join(s, " "))
	}
}

func Warn(s ...string) {

	color.NoColor = NO_COLOR

	if !SILENT {

		logger := log.New(OUT, "", 0)
		coloredPrefix := "[" + color.New(WRN_COLOR).SprintFunc()("WRN") + "] "

		logger.SetPrefix(coloredPrefix)
		logger.Println(strings.Join(s, " "))
	}
}

func Error(err error) {

	color.NoColor = NO_COLOR

	if !SILENT && err != nil {

		logger := log.New(OUT, "", 0)
		coloredPrefix := "[" + color.New(ERR_COLOR).SprintFunc()("ERR") + "] "

		logger.SetPrefix(coloredPrefix)
		logger.Println(err)
	}
}

func Fatal(err error) {

	color.NoColor = NO_COLOR

	if err != nil {

		logger := log.New(OUT, "", 0)
		coloredPrefix := "[" + color.New(FTL_COLOR).SprintFunc()("FTL") + "] "

		logger.SetPrefix(coloredPrefix)
		logger.Println(err)

		os.Exit(0)
	}
}
