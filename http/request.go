package http

import (
	"github.com/mskoroglu/golaxy/http/request/path"
	"github.com/mskoroglu/golaxy/http/request"
)

type Request struct {
	URL          string
	Param        request.Param
	Header       Header
	PathVariable path.Variable
}
