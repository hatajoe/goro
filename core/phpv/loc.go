package phpv

import (
	"fmt"
	"io"
)

type Loc struct {
	Filename   string
	Line, Char int
}

func (l *Loc) Loc() *Loc {
	return l
}

func (l *Loc) Run(ctx Context) (*ZVal, error) {
	// just a checkpoint, do nothing
	return nil, nil
}

func (l *Loc) Dump(w io.Writer) error {
	return nil
}

func (l *Loc) Error(e error) *PhpError {
	// fill location if missing
	switch err := e.(type) {
	case *PhpError:
		if err.l == nil {
			err.l = l
		}
		return err
	default:
		return &PhpError{Err: e, Code: E_ERROR, l: l}
	}
}

func (l *Loc) Errorf(ctx Context, code PhpErrorType, f string, arg ...interface{}) *PhpError {
	return &PhpError{Err: fmt.Errorf(f, arg...), l: l}
}

func (l *Loc) String() string {
	return fmt.Sprintf("at %s on line %d", l.Filename, l.Line)
}
