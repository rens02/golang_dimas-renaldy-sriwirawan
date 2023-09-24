# Intro RESTful API

API RESTful adalah sebuah pendekatan atau gaya arsitektur perangkat lunak yang didesain khusus untuk memudahkan `interaksi` antar sistem, terutama melalui `jaringan internet`. Pendekatan ini bertujuan untuk mencapai skalabilitas yang tinggi, penggunaan sumber daya yang efisien, dan dapat diimplementasikan dengan mudah.

API RESTful didasarkan pada beberapa prinsip dasar, seperti:

- `Stateless`: setiap permintaan yang diterima oleh server harus mencantumkan semua informasi yang diperlukan untuk memenuhi permintaan tersebut. Server tidak harus menyimpan informasi atau state dari permintaan sebelumnya.
- `Uniform Interface`: API harus memiliki antarmuka yang seragam dan mudah dipahami, seperti penggunaan URI untuk mengidentifikasi sumber daya, penggunaan metode HTTP yang tepat untuk berinteraksi dengan sumber daya, dan penggunaan format data yang standar (seperti JSON atau XML) untuk mengirim dan menerima data.
- `Cacheable`: server harus memberikan informasi yang memungkinkan klien untuk melakukan caching data yang diterima dari server.
- `Layered System`: arsitektur API harus mendukung lapisan yang berbeda, seperti reverse proxy atau cache, sehingga dapat meningkatkan skalabilitas dan keamanan sistem.

Berikut ini adalah penjelasan singkat dari masing-masing metode HTTP yang digunakan pada API RESTful:

- `GET`: digunakan untuk `mengambil` informasi dari server. Permintaan GET tidak merubah status sumber daya di server.

- `POST`: digunakan untuk `mengirim` data baru ke server untuk dibuat atau diolah. Permintaan POST mengubah status sumber daya di server.

- `PUT`: digunakan untuk `mengirim` data yang `menggantikan` data yang sudah ada pada server. Permintaan PUT juga dapat digunakan untuk `membuat data baru` jika data yang diidentifikasi `tidak ada` di server.

- `DELETE`: digunakan untuk `menghapus` data dari server.

- `PATCH`: digunakan untuk mengirimkan data yang hanya `sebagian` saja yang akan dirubah dari sumber daya yang sudah ada pada server. Permintaan PATCH tidak merubah status sumber daya di server, namun hanya mengubah data yang `diinginkan`.