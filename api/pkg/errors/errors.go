package errors

type ErrorType string

func (e ErrorType) Error() string {
	return string(e)
}

const (
	QueryError ErrorType = "cannot execute select query"
)
