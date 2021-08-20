package exception

type Duplicity interface {
	Error() string
	Duplicity() bool
}

type DataDuplicity struct {
	ErrMessage string
}

func (duplicity DataDuplicity) Error() string {
	return duplicity.ErrMessage
}

func (duplicity DataDuplicity) Duplicity() bool {
	return true
}
