package view

import (
	"html/template"
	"io"
	"net/http"
)

var (
	VIEW_PREFIX = "view/"
	VIEW_SUFFIX = ".html"
)

type ViewResolver struct {
	View   *View
	Writer io.Writer
}

func (v *ViewResolver) Resolve() {
	handlerView := *v.View
	layoutName := handlerView.GetLayout()
	viewName := handlerView.GetView()
	var views = make([]string, 0)
	if layoutName != "" {
		views = append(views, VIEW_PREFIX+layoutName+VIEW_SUFFIX)
	}
	views = append(views, VIEW_PREFIX+viewName+VIEW_SUFFIX)

	tpl := template.Must(template.New(viewName).Funcs(v.GetFuncMap()).ParseFiles(views...))
	err := tpl.ExecuteTemplate(v.Writer, layoutName+VIEW_SUFFIX, handlerView.modelMap)
	if err != nil {
		http.Error(v.Writer.(http.ResponseWriter), http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (v *ViewResolver) GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"html": func(str string) template.HTML {
			return template.HTML(str)
		},
	}
}
