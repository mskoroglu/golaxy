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

	t, err := template.ParseFiles(views...)
	if err != nil {
		http.Error(v.Writer.(http.ResponseWriter), err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(v.Writer, handlerView.modelMap); err != nil {
		http.Error(v.Writer.(http.ResponseWriter), err.Error(), http.StatusInternalServerError)
	}
}
