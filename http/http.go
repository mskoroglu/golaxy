package http

import (
	"reflect"
	"net/http"
	"fmt"
	"strconv"
	"encoding/json"
	"regexp"
	"github.com/mskoroglu/golaxy/http/request/path"
	"github.com/mskoroglu/golaxy/http/request"
	"github.com/mskoroglu/golaxy/view"
)

type handlerFunc struct {
	path      string
	method    string
	function  reflect.Value
	input     []reflect.Value
	output    []reflect.Value
	request   Request
	goRequest http.Request
	view      view.View
}

var (
	STATIC_DIRECTORY                  = "static"
	STATIC_PATH_PREFIX                = "/static/"
	handlers           []*handlerFunc = make([]*handlerFunc, 0)
)

func Get(path string, function interface{}) error {
	createHandler(path, "GET", function)
	return nil
}

func Post(path string, function interface{}) error {
	createHandler(path, "POST", function)
	return nil
}

func Put(path string, function interface{}) error {
	createHandler(path, "PUT", function)
	return nil
}

func Delete(path string, function interface{}) error {
	createHandler(path, "DELETE", function)
	return nil
}

func createHandler(path, method string, function interface{}) {
	functionValue := reflect.ValueOf(function)
	if functionValue.Kind() != reflect.Func {
		panic("Handler must be function!")
	}
	registerHandler(&handlerFunc{
		path:     path,
		method:   method,
		function: functionValue,
	})
}

func registerHandler(handler *handlerFunc) {
	funcValueType := handler.function.Type()
	var inArgs = make([]reflect.Value, funcValueType.NumIn())
	for i := 0; i < funcValueType.NumIn(); i++ {
		switch funcValueType.In(i).Elem() {
		case reflect.TypeOf(http.Request{}):
			inArgs[i] = reflect.ValueOf(&handler.goRequest)
		case reflect.TypeOf(Request{}):
			inArgs[i] = reflect.ValueOf(&handler.request)
		case reflect.TypeOf(request.Param{}):
			inArgs[i] = reflect.ValueOf(&handler.request.Param)
		case reflect.TypeOf(Header{}):
			inArgs[i] = reflect.ValueOf(&handler.request.Header)
		case reflect.TypeOf(path.Variable{}):
			inArgs[i] = reflect.ValueOf(&handler.request.PathVariable)
		case reflect.TypeOf(view.View{}):
			inArgs[i] = reflect.ValueOf(&handler.view)
		}
	}
	handler.input = inArgs
	handlers = append(handlers, handler)
}

func isMatches(handler *handlerFunc, request *http.Request) bool {
	return regexp.MustCompile(`^` + handler.path + `$`).MatchString(request.URL.Path) &&
		handler.method == request.Method
}

func assignPathVariables(handler *handlerFunc, request *http.Request) {
	handler.request.PathVariable = make(map[string]string)
	exp := regexp.MustCompile(`^` + handler.path + `$`)
	matches := exp.FindStringSubmatch(request.URL.Path)
	if matches != nil {
		for i, name := range exp.SubexpNames() {
			if i != 0 {
				handler.request.PathVariable[name] = matches[i]
			}
		}
	}

}

func assignRequestParams(handler *handlerFunc, request *http.Request) {
	handler.request.Param = make(map[string][]string)
	for key, value := range request.URL.Query() {
		handler.request.Param[key] = value
	}
}

func assignRequestHeaders(handler *handlerFunc, request *http.Request) {
	handler.request.Header = make(map[string][]string)
	for key, value := range request.Header {
		handler.request.Header[key] = value
	}
}

func requestHandler(writer http.ResponseWriter, request *http.Request) {
	var notFound bool = true
	for i := 0; i < len(handlers); i++ {
		handler := handlers[i]
		if isMatches(handler, request) {
			processRequest(handler, writer, request)
			notFound = false
		}
	}
	if notFound {
		http.Error(writer, "Not Found!", http.StatusNotFound)
	}
}

func isView(handler *handlerFunc) bool {
	var isView bool
	for i := 0; i < handler.function.Type().NumIn(); i++ {
		isView = handler.function.Type().In(i).Elem() == reflect.TypeOf(view.View{})
		if isView {
			return isView
		}
	}
	return isView
}

func processRequest(handler *handlerFunc, writer http.ResponseWriter, request *http.Request) {
	handler.goRequest = *request
	handler.request.URL = request.URL.Path
	assignPathVariables(handler, request)
	assignRequestParams(handler, request)
	assignRequestHeaders(handler, request)

	funcOut := handler.function.Call(handler.input)
	if isView(handler) {
		if len(handler.view.GetView()) == 0 {
			handler.view.SetView(funcOut[0].Interface().(string))
		}
		viewResolver := view.ViewResolver{View: &handler.view, Writer: writer}
		viewResolver.Resolve()
	} else if !isView(handler) && len(funcOut) > 0 {
		funcOutObject := funcOut[0].Interface()
		writer.Header().Set("Content-Type", "application/json")
		responseWrite(writer, funcOutObject)
	}
	writer.Header().Set("Server", "Golaxy Web Server")
}

func responseWrite(writer http.ResponseWriter, funcOutObject interface{}) {
	output, err := json.Marshal(funcOutObject)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonString := string(output)
	fmt.Fprint(writer, jsonString)
}

func StartHttpServer(ip string, port int) {
	fs := http.FileServer(http.Dir(STATIC_DIRECTORY))
	http.Handle(STATIC_PATH_PREFIX, http.StripPrefix(STATIC_PATH_PREFIX, fs))
	http.HandleFunc("/", requestHandler)
	socket := ip + ":" + strconv.Itoa(port)
	http.ListenAndServe(socket, nil)
}
