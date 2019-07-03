# golaxy
Go dilini öğrenmek amacıyla başladığım bir çalışmadır. Kararlı çalıştığını iddia etmiyor, uygulamalarınızda kullanmanız sonucunda alabileceğiniz hatalar için sorumluluk kabul etmiyorum :)

### Kullanım
```
go get github.com/mskoroglu/golaxy
```
komutu ile bağımlılık çekildikten sonra
```
var _ = http.Get(`/hello/(?P<name>\w+)`, func(pv *path.Variable) string {
	return "Hello " + pv.GetString("name")
})
```
fonksiyonunu yazıp tarayıcıda ```http://localhost:8080/hello/Gopher``` adresini açın.
