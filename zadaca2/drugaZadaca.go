package main

import "fmt"

// 1.
type Pravokutnik struct {
	Sirina float64
	Visina float64
}

func (p Pravokutnik) Povrsina() float64 {
	return p.Sirina * p.Visina
}

func (p Pravokutnik) Opseg() float64 {
	return 2 * (p.Sirina + p.Visina)
}

// 2.

type Adresa struct {
	Grad  string
	Ulica string
}

type Osoba struct {
	Ime     string
	Prezime string
	Godine  int
	Adresa  Adresa
}

func (o Osoba) Ispisi() {
	fmt.Printf("%s %s %d godina, živi u %s , %s\n", o.Ime, o.Prezime, o.Godine, o.Adresa.Grad, o.Adresa.Ulica)
}

func main() {

	/* 1. Napiši strukturu koja predstavlja pravokutnik i sadrži polja za širinu i visinu.
	   Kreiraj metode za izračunavanje površine i opsega pravokutnika,metode moraju biti vezane za novo kreiranu strukturu Pravokutnika.
	   U main funkciji inicijalizirajte jedan pravokutnik, te pozovite iznad kreirane metode za računanje površine i opsega */

	p := Pravokutnik{Sirina: 5, Visina: 10}

	fmt.Println("Površina pravokutnika je:", p.Povrsina(), "Opseg pravokutnika je:", p.Opseg())

	/*2. Napiši strukturu koja predstavlja adresu, koja sadrži polja za grad i ulicu. Zatim napiši strukturu koja predstavlja osobu,
		koja sadrži ime, godine i adresu.
	    Kreiraj metodu koja ispisuje puni opis osobe, uključujući njezinu adresu.
	    (Sva polja) u formatu:
	    Ime Prezime, 20 godina, živi u Grad, Ulica.*/

	osoba := Osoba{Ime: "Antonio", Prezime: "Sego", Godine: 25, Adresa: Adresa{Grad: "Posusje", Ulica: "Ulica 5"}}
	osoba.Ispisi()

}
