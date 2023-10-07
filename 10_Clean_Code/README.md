# Soal Clean Code (Eksplorasi)
## Soal Prioritas 1
1. Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan.
    - Berapa banyak kekurangan dalam penulisan kode tersebut?
    - Bagian mana saja terjadi kekurangan tersebut?
    - Tuliskan alasan dari setiap kekurangan tersebut. Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan berikut.
<br>Jawab<br>
```
package main

// Kekurangan 1: tipe data yang salah pada user
type user struct {
    id       int    
    username int    // Alasan 1: Username dan password biasanya berupa string, bukan integer.
    password int    // Alasan 2: Sebaiknya menggunakan tipe data yang lebih aman untuk menyimpan password, seperti string hashed.
}

// Kekurangan 2: Struktur "userservice" memiliki field yang tidak dijelaskan dengan baik.
//    Alasan: Field "t" seharusnya memiliki dokumentasi atau nama yang lebih deskriptif untuk menjelaskan maksudnya.
type userservice struct {
    t []user
}

// Kekurangan 3: Fungsi "getallusers" menggunakan receiver by value, sehingga membuat salinan besar dari slice "t".
//    Alasan: Ini bisa menjadi masalah kinerja karena meng-copy seluruh data ketika mengembalikannya. 
// Seharusnya menggunakan pointer receiver (&userservice) untuk menghindari masalah ini.
func (u userservice) getallusers() []user {
    return u.t
}

// Kekurangan 4: Fungsi "getuserbyid" mengembalikan nilai default jika id tidak ditemukan, tanpa memberikan informasi/error yang jelas.
func (u userservice) getuserbyid(id int) user {
    for _, r := range u.t {
        if id == r.id {
            return r
        }
    }

    return user{} // Mengembalikan user kosong tanpa memberikan pesan/error yang berguna.
}
```

2. Kode berikut ini dituliskan tanpa mengikuti kebiasaan penulisan yang disarankan. Ubahlah penulisan kode berikut menjadi lebih terbaca dan lebih rapi!
<br>********************************Jawab :  [Source Code](Praktikum/kendaraan.go)******************************** <br>

# Summary
Clean code merupakan menulis kode yang mudah `dibaca`, `dijaga`, dan `dipahami`. Berikut adalah beberapa konsep utama clean code:

- `Konvensi penamaan`: Gunakan nama yang bermakna untuk `variabel`, `fungsi`, dan `package`. Gunakan nama yang pendek dan jelas untuk variabel yang memiliki cakupan yang terbatas. serta menerapkan `camelcase` untuk variabel dan huruf `kapital` diawal pada nama struct, fungsi, dan method
- `Fungsi`: Tulis fungsi kecil dan terfokus yang melakukan satu hal dan melakukannya dengan baik. Gunakan nama fungsi yang `deskriptif` yang menjelaskan `apa yang dilakukan` oleh fungsi tersebut.
- `Komentar`: Gunakan komentar untuk menjelaskan mengapa kode ditulis, bukan apa yang dilakukan oleh kode tersebut. Komentar harus digunakan dengan `hemat` dan hanya jika `diperlukan`.
- `Organisasi` kode: Susun kode ke dalam paket kecil dan terfokus yang `mudah dipahami`.