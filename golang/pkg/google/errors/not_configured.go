package errors

type ErrNotConfigured struct {
	error
}

func NewNotConfigured(err error) *ErrNotConfigured {
	return &ErrNotConfigured{
		error: err,
	}
}

func (*ErrNotConfigured) NotConfigured() bool {
	return true
}
