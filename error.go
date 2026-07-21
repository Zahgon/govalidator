package govalidator

type Errors []error

func (es Errors) Errors() []error { _ = "STUB: not implemented"; return nil }

func (es Errors) Error() string { _ = "STUB: not implemented"; return "" }

type Error struct {
	Name                     string
	Err                      error
	CustomErrorMessageExists bool

	Validator string
	Path      []string
}

func (e Error) Error() string { _ = "STUB: not implemented"; return "" }
