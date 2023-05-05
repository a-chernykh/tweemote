package errors

type withMessage struct {
	cause error
	msg   string
}

func (w *withMessage) Error() string {
	return w.msg + ": " + w.cause.Error()
}

func Wrap(e error, msg string) error {
	return &withMessage{
		cause: e,
		msg:   msg,
	}
}
