# golaxy
Go dilini öğrenmek amacıyla başladığım bir çalışmadır. Kararlı çalıştığını iddia etmiyor, uygulamalarınızda kullanmanız sonucunda alabileceğiniz hatalar için sorumluluk kabul etmiyorum :)

Kullanmaktan keyif aldığım Spring Framework'ten esinlendim. Java'cı arkadaşlar dejavu yaşadıklarını sanmasınlar :)

Repoya koyacak ad bulamadım ve Gopher geleneğini takip edip golaxy gibi tuhaf bir şey uydurabildim. Talep olursa Godaman olarak değiştiririz. Ve yine talep olursa bundan sonra asla isim arayışlarına girmeyebilirim de :)

### Kullanım
```
go get github.com/mskoroglu/golaxy
```
komutu ile bağımlılık GOPATH'e indirildikten sonra
```
var _ = http.Get(`/hello/(?P<name>\w+)`, func(pv *path.Variable) string {
	return "Hello " + pv.GetString("name")
})
```
fonksiyonunu yazıp ve tarayıcıda ```/hello/Gopher``` adresini açın.

Örnek uygulama: https://github.com/mskoroglu/golaxy-example

### Şimdilik Hedefler
- Kullanımı kolay
- Controller mekanizması basit ama güçlü
- Mümkün olduğunca üçüncü parti paketlerden uzak
- Çıktı boyutu küçük
- Yeniden derlemeye gerek kalmadan yapılabilecek ayarlar
- Registry işleminin otomatikleştirilmesi

### Sonraki Hedefler
- Dependency Injection
- Popüler ORM araçları için JPA tadında bir çatı
- Mikroservis mimarisine uygun pratik bir geliştirme ortamı
- View katmanına da el atıp MVC'ye uygun bir zemin oluşturulması
- URL Mapping için Spring Framework tarzı named regexp kullanılması

Belirttiğim gibi; amacım dili öğrenmekle birlikte, ilerleyen zamanlarda lazım olduğunda hızlı ve basit bir şekilde işimi görebilecek bir araç oluşturmaktır.

#### Son olarak
Eleştirilere ve katkılara açıktır. Ama sert vurmayın lütfen :)