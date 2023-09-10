package main

import "fmt"

type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

// Metode untuk menambah kecepatan mobil
func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
	m.KecepatanPerJam += kecepatanBaru
}

// Metode untuk menggerakkan mobil dengan peningkatan kecepatan
func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}

func main() {
	// Membuat mobil cepat
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	// Membuat mobil lamban
	mobilLamban := Mobil{}
	mobilLamban.Berjalan()

	// Menampilkan kecepatan mobil cepat dan mobil lamban
	fmt.Printf("Kecepatan Mobil Cepat: %d km/jam\n", mobilCepat.KecepatanPerJam)
	fmt.Printf("Kecepatan Mobil Lamban: %d km/jam\n", mobilLamban.KecepatanPerJam)
}
