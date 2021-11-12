package ntuapi

type DateRangeError struct {
	message string
}

type NullDateError struct {
	message string
}

func (err *DateRangeError) Error() string {
	return err.message
}

func (err *NullDateError) Error() string {
	return err.message
}
