package refer

import "github.com/kmou424/ero"

var (
	ErrRefWithDefKeyNotFound = ero.New("can't find ref with default key")
	ErrRefNotFound           = ero.New("can't find ref with key")
)
