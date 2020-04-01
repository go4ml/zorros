package zorros

import (
	"fmt"
	"golang.org/x/xerrors"
	"strings"
)

func Trace(err error) error {
	if _, ok := err.(xerrors.Formatter); ok {
		return err
	}
	return zerror{err, xerrors.Caller(1)}
}

func Errorf(f string, a ...interface{}) error {
	return zerror{fmt.Errorf(f, a...), xerrors.Caller(1)}
}

type zerror struct {
	error
	frame xerrors.Frame
}

func (e zerror) FormatError(p xerrors.Printer) error {
	p.Print(e.error.Error() + " at ")
	e.frame.Format(p)
	return nil
}

func stringifyError(err error) (string, error) {
	ep := &errorPrinter{details: true}
	if f, ok := err.(xerrors.Formatter); ok {
		err = f.FormatError(ep)
	} else {
		ep.Print(err.Error())
		err = nil
	}
	return strings.Join(strings.Fields(ep.String()), " "), err
}


