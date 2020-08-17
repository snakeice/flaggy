package utils

import (
	"os"

	"github.com/ztrue/tracerr"
)

func ErrCheck(err error) {
	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
		os.Exit(1)
	}
}
