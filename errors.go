package pikpakapi

import "errors"

var (
	ErrNotFoundFolder = errors.New("Not found PikPak folder")
	ErrNotFoundFile   = errors.New("Not found PikPak file")
)
