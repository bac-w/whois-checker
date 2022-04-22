package log

import "fmt"

type Level int

const (
	ERROR = iota
	WARNING
	INFO
	DEBUG
)

var logLevel Level = INFO

func SetLevel(level Level) {
	logLevel = level
}

func GetLevel(intG int) (Level, error) {
	switch intG {
	case 2:
		return ERROR, nil
	case 1:
		return WARNING, nil
	case 0:
		return INFO, nil
	case 3:
		return DEBUG, nil
	default:
		return 0, fmt.Errorf("invalid verbosity level %s", intG)
	}
}
