package ntuapi

type ErrDateRange struct {
	message string
}

type ErrNullDate struct {
	message string
}

type ErrWeekOutOfRange struct {
	message string
}

func (err *ErrDateRange) Error() string {
	return err.message
}

func (err *ErrNullDate) Error() string {
	return err.message
}

func (err *ErrWeekOutOfRange) Error() string {
	return err.message
}
