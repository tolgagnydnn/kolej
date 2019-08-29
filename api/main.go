package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// OkulListe okulların tüm listesini verir.
func OkulListe(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	var o Okul
	var liste = o.Liste()

	json.NewEncoder(w).Encode(liste)
}

// OkulEkle sisteme yeni bir okul bilgisi ekler
func OkulEkle(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	var o1 = Okul{
		ID:      1,
		Adi:     "Özel Cent Ortaokulu",
		Adres:   "Tarabya Şalçıkır Cad. No:136 İstanbul/Sarıyer",
		Telefon: "0212 2233133",
		Website: "www.centkoleji.k12.tr",
		Sehir:   "İstanbul",
		Semt:    "Sarıyer",
	}
	var o2 = Okul{
		ID:      2,
		Adi:     "İstanbul Erkek Liseliler Eğitim Vakfı Özel 125.Yıl Ortaokulu",
		Adres:   "Nişantepe Mah. Ensar Cad. No: 4/2 İstanbul/Çekmeköy",
		Telefon: "0216 3046909",
		Website: "http://www.ielev.k12.tr/",
		Sehir:   "İstanbul",
		Semt:    "Çekmeköy",
	}
	var o3 = Okul{
		ID:      3,
		Adi:     "Özel Batı Koleji Ortaokulu",
		Adres:   "İnönü Mah. 1750 Sokak No:13-17 Ankara/Yenimahalle",
		Telefon: "0312 2775454",
		Website: "www.batikoleji.k12",
		Sehir:   "Ankara",
		Semt:    "Yenimahalle",
	}
	var o4 = Okul{
		ID:      4,
		Adi:     "Özel Konya Nesibe Aydın Ortaokulu",
		Adres:   "Beyhekim Mah. Darül Hilafet Sok. No:1 Konya/Selçuklu",
		Telefon: "0332 3208511",
		Website: "www.nesibeaydin.k12.tr/konya/",
		Sehir:   "Konya",
		Semt:    "Selçuklu",
	}
	var o5 = Okul{
		ID:      5,
		Adi:     "Özel İnegöl Doğa Ortaokulu",
		Adres:   "Akhisar Mah. Bursa Karayolu No:4 İnegöl/ Bursa Bursa/İnegöl",
		Telefon: "0224 7114100",
		Website: "www.dogakoleji.k12.tr",
		Sehir:   "Bursa",
		Semt:    "İnegöl",
	}

	var o Okul
	o.Ekle(o1)
	o.Ekle(o2)
	o.Ekle(o3)
	o.Ekle(o4)
	o.Ekle(o5)

	var cevap = `{"status": "ok"}`
	fmt.Fprintf(w, "%s", cevap)
}

// Index api indeks sayfasını verir.
func Index(w http.ResponseWriter, r *http.Request) {
	var html = `<!doctype html><html lang="tr"><head>
		<meta charset="utf-8">
		<title>Kolej Projesi API index sayfası</title>
		<style type="text/css">
			.tablo { background-color:#eee;border-collapse:collapse; }
			.tablo th { background-color:#000;color:white;}
			.tablo td, .tablo th { padding:12px;border:1px solid #000; }
		</style>
		</head>
		<body>
			<h2>Metod Listesi</h2>
			<table class="tablo">
			<tr>
				<th>Adres</th>
				<th>Tanım</th>
			</tr>
			<tr>
				<td><a href="/">/</a></td>
				<td>API metodlarını listeleyen indeks sayfasını gösterir.</td>
			</tr>
			<tr>
				<td><a href="okul/liste">/okul/liste</a></td>
				<td>Sistemde kayıtlı tüm okulları bir JSON listesi olarar verir.</td>
			</tr>
			<tr>
				<td><a href="okul/ekle">/okul/ekle</a></td>
				<td>Sisteme yeni bir okul bilgisi ekler. Sonucu JSON formatında verir.</td>
			</tr>
			</table>
		</body></html>`

	fmt.Fprintf(w, "%s", html)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/okul/liste", OkulListe).Methods("GET")
	router.HandleFunc("/okul/ekle", OkulEkle).Methods("GET")

	log.Println("API Service is starting...(8090)")
	log.Fatal(http.ListenAndServe(":8090", router))
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Content-Type", "application/json")
}
