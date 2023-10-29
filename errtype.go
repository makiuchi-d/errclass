package errtype

type Type func(error) error

type typedErr struct {
	error
	np *string
}

type dummy struct{}

var (
	_ error = Type(nil)
	_ error = typedErr{}
	_ error = dummy{}
)

func New(name string) Type {
	np := &name
	return func(err error) error {
		if err == nil {
			return nil
		}
		return typedErr{err, np}
	}
}

func (typ Type) Error() string {
	err := typ(dummy{})
	if te, ok := err.(typedErr); ok {
		return *te.np
	}
	return err.Error()
}

func (te typedErr) Unwrap() error {
	return te.error
}

func (te typedErr) Is(target error) bool {
	typ, ok := target.(Type)
	if !ok {
		return false
	}
	tt, ok := typ(dummy{}).(typedErr)
	if !ok {
		return false
	}
	return tt.np == te.np
}

func (dummy) Error() string { return "" }
