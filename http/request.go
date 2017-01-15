package http

import (
	"golaxy/http/request/path"
	"golaxy/http/request"
)

type Request struct {
	URL          string
	Param        request.Param
	Header       Header
	PathVariable path.Variable
}
