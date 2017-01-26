package view

// Şablonlar ile çalışılacağı zaman, handler fonksiyonun bu tipte bir parametre kabul etmesi beklenir.
type View struct {
	layout   string
	name     string
	modelMap map[string]interface{}
	redirect string
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

// View'a bir değer gönderilmek istenildiğinde key-value şeklinde kullanılır.
func (v *View) Put(key string, value interface{}) *View {
	if v.modelMap == nil {
		v.modelMap = make(map[string]interface{})
	}
	v.modelMap[key] = value
	return v
}

// HTTP 301 yönlendirmesi yapmak için kullanılır.
func (v *View) Redirect(url string) {
	v.redirect = url
}

func (v *View) IsRedirected() (bool, string) {
	return v.redirect != "", v.redirect
}
