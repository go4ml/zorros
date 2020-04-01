package zorros

import (
	"golang.org/x/xerrors"
	"strings"
)

type zpanic struct{ err error }

func Panic(err error) interface{} {
	return zpanic{err}
}

func (x zpanic) stringify(indepth bool) string {
	s, e := stringifyError(x.err)
	ns := []string{s}
	for e != nil && indepth {
		s, e = stringifyError(e)
		ns = append(ns, s)
	}
	return strings.Join(ns, "\n")
}

func (x zpanic) Error() string {
	return x.stringify(false)
}

func (x zpanic) String() string {
	return x.stringify(true)
}

func (x zpanic) Unwrap() error {
	if w, ok := x.err.(xerrors.Wrapper); ok {
		return w.Unwrap()
	}
	return x.err
}


