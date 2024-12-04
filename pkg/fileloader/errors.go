package fileloader

type ErrorType int

type LoaderError struct {
	Message string
	Err     error
}

func (e LoaderError) Error() string { return e.Message + ": " + e.Err.Error() }

type FileError struct{ LoaderError }

type LineError struct{ LoaderError }
