#ORM DAN CODE STRUCTURE (MVC)

- ORM (Object-Relational Mapping) adalah teknik dalam pengembangan perangkat lunak untuk `memetakan data` dari database relasional ke dalam objek-objek yang bisa diakses oleh bahasa pemrograman yang digunakan. Dalam bahasa Go, terdapat beberapa ORM yang populer, seperti `GORM`, `XORM`, dan `QBS`.

- Sementara itu, code structure dalam bahasa Go biasanya mengikuti pola `MVC` (Model-View-Controller). Dalam MVC, `Model` mewakili data atau objek, `View` bertanggung jawab atas tampilan, dan `Controller` sebagai pengontrol atau penghubung antara Model dan View. Penerapan MVC pada Go dapat membantu `mengorganisir` kode menjadi lebih `terstruktur` dan mudah dipelihara.

- Dalam MVC, `Model` biasanya berisi definisi `struktur data` atau tabel dalam database, seperti user atau blog, serta fungsi-fungsi untuk membaca atau mengubah data dalam database. `View` berisi `tampilan` atau halaman-halaman yang menampilkan data dari Model. Sedangkan `Controller` berfungsi sebagai `penghubung` antara Model dan View, yang bertanggung jawab atas `logika` bisnis dan aksi-aksi yang terjadi pada data.