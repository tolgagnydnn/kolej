package main

type Okul struct {
	ID      int    `json:"id,omitempty"`
	Adi     string `json:"adi,omitempty"`
	Adres   string `json:"adres,omitempty"`
	Telefon string `json:"telefon,omitempty"`
	Website string `json:"website,omitempty"`
	Sehir   string `json:"sehir,omitempty"`
	Semt    string `json:"semt,omitempty"`
}

func (o *Okul) Liste() []Okul {
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

	var liste = make([]Okul, 0)
	liste = append(liste, o1)
	liste = append(liste, o2)
	liste = append(liste, o3)
	liste = append(liste, o4)
	liste = append(liste, o5)

	return liste
}
