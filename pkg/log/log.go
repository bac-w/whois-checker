package log

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

func Fatal(v ...interface{}) {
	if logLevel >= ERROR {
		log.Fatalf(color.RedString("[FATAL] ")+"%v", v...)
		os.Exit(1)
	}
}

func Fatalf(format string, v ...interface{}) {
	if logLevel >= ERROR {
		log.Fatalf(color.RedString("[FATAL] ")+"%v", fmt.Sprintf(format, v...))
		os.Exit(1)
	}
}

func Error(v ...interface{}) {
	if logLevel >= ERROR {
		log.Printf(color.MagentaString("[ERROR] ")+"%v", v...)
	}
}

func Errorf(format string, v ...interface{}) {
	if logLevel >= ERROR {
		log.Printf(color.MagentaString("[ERROR] ")+"%v", fmt.Sprintf(format, v...))
	}
}

func Warning(v ...interface{}) {
	if logLevel >= WARNING {
		log.Printf(color.MagentaString("[WARNING] ")+"%v", v...)
	}
}

func Warningf(format string, v ...interface{}) {
	if logLevel >= WARNING {
		log.Printf(color.MagentaString("[WARNING] ")+"%v", fmt.Sprintf(format, v...))
	}
}

func Info(v ...interface{}) {
	if logLevel >= INFO {
		log.Printf(color.GreenString("[INFO] ")+"%v", v...)
	}
}

func Infof(format string, v ...interface{}) {
	if logLevel >= INFO {
		log.Printf(color.GreenString("[INFO] ")+"%v", fmt.Sprintf(format, v...))
	}
}
func InfoI(messageName string) {
	if logLevel >= INFO {
		fmt.Println(color.GreenString("%s", messageName))
	}
}
func InfoIp(messageName string) {
	if logLevel >= INFO {
		fmt.Print(color.GreenString("%s", messageName))
	}
}
func InfoIf(messageName string, v ...interface{}) {
	if logLevel >= INFO {
		fmt.Printf(color.GreenString("%s", messageName)+"%v", v...)
	}
}
func Debug(v ...interface{}) {
	if logLevel >= DEBUG {
		log.Printf(color.YellowString("[DEBUG] ")+"%v", v...)
	}
}

func Debugf(format string, v ...interface{}) {
	if logLevel >= DEBUG {
		log.Printf(color.YellowString("[DEBUG] ")+"%v", fmt.Sprintf(format, v...))
	}
}
