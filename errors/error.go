package errors

type Error struct {
	msg string
}

func (err Error) Error() string {
	return err.msg
}

func New(msg string) error {
	return &Error{msg: msg}
}

var Nil = New("object is nil")
