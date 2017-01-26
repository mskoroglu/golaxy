package golaxy

import (
	"github.com/mskoroglu/golaxy/http"
)

// Web sunucusu başlatılırken (şimdilik) çağırılması gereken fonksiyondur.
var Run = func(args ...interface{}) {
	http.StartHttpServer()
}
