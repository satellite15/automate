package manifest

import "errors"

var (
	ErrNoSuchManifest = errors.New("unable to locate manifest")
	ErrCannotParse    = errors.New("failed to parse manifest")
	ErrInvalidSchema  = errors.New("unknown schema")
)
