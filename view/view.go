package view

type View struct {
	layout   string
	name     string
	modelMap map[string]interface{}
}

func (v *View) SetLayout(name string) *View {
	v.layout = name
	return v
}

func (v *View) GetLayout() string {
	return v.layout
}

func (v *View) SetView(name string) *View {
	v.name = name
	return v
}

func (v *View) GetView() string {
	return v.name
}

func (v *View) Put(key string, value interface{}) *View {
	if v.modelMap == nil {
		v.modelMap = make(map[string]interface{})
	}
	v.modelMap[key] = value
	return v
}
