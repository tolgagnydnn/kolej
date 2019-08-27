# KOLEJ PROJESİ

## Başlangıç
Kolej projesi VueJS ve Go dili ile geliştirilmekte olan ve iki ayrı proje içeren bir sistemdir. Bu iki farklı ortamı bir arada kullanabilmek adına geliştirme ortamının tamamını docker üzerinde inşa ettik. Kısaca mevcut projeyi kendi bilgisayarınıza indirdikten sonra yapmanız gereken aşağıdaki komutu çalıştırmak.

```
    $ docker-compose up
```

Sonrasında VueJS ve gerekli bağımlılıkları önyüz (frontend) için ve Go ve gerekli paketler arka uç (backend) için yüklenipkodlar derlenecek ve projeler çalışmaya başlayacaktır. Önyüze ulaşmak için tarayıcınıza **http://localhost:8080** ve API bölümüne ulaşmak için **htt://locahost:8090** yazabilirsiniz.

## Amaç
Özel okullar özelinde bir çok seçenek ile karşı karşıya kalan öğrenci velilerinin daha doğru, basit ve detaylı bir seçim yapabilmeleri için mevcut özel okulların farklı kriterlere göre karşılaştırılmasını sağlamak.

## Hedef
Türkiye’deki tüm özel okulların sisteme eklenip ilk etapta en az 3 kriter baz alınarak karşılaştırılabilmesinin sağlanmasıdır.

## Teknoloji: 
Sistem genelinde güncel teknolojiler kullanılacaktır. Önyüz (frontend) bölümü için vuejs ve bootstrap kullanılacak.VueJS metodlarının hepsinin açıklaması üzerinde olacak ayrıca testleri yapılmış olacaktır. Arkauç (backend) için go dili ile geliştirilecek bir api kullanılacak ve tüm metodların açıklaması ve testi yazılmış olacaktır.

### Teknoloji Listesi
- Go
- VueJS
- Bootstrap
- CSS (SASS)
- HTML
- docker
- docker-compose
