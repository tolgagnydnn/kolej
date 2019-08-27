package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func OkulListe(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	var o Okul
	var liste = o.Liste()

	json.NewEncoder(w).Encode(liste)
}

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
			</table>
		</body></html>`

	fmt.Fprintf(w, "%s", html)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/okul/liste", OkulListe).Methods("GET")

	log.Println("API Service is starting...(8090)")
	log.Fatal(http.ListenAndServe(":8090", router))
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Content-Type", "application/json")
}
