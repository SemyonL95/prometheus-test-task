package api

type Сache interface {
	Set(string) error
}

type Metrics interface {
	Inc()
}

type ErrValExists struct {
	Msg string
}

func (err ErrValExists) Error() string {
	return err.Msg
}
