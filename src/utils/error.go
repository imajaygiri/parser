package utils

import (
	"fmt"
	"runtime"
)

func Error(msg string) error {
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)

	return fmt.Errorf("%s:%d (%s): %s",
		file,
		line,
		fn.Name(),
		msg,
	)
}
