package presenters

import (
	"net/http"

	"github.com/go-log/log"
)

//go:generate mockery -name=Logger -inpkg -case=underscore -testonly

// Logger ...
//
// It's used only for mock generating.
type Logger interface {
	log.Logger
}

//go:generate mockery -name=ResponseWriter -inpkg -case=underscore -testonly

// ResponseWriter ...
//
// It's used only for mock generating.
type ResponseWriter interface {
	http.ResponseWriter
}
