# Clean and Hexagonal Architecture

`Clean Architecture` dan `Hexagonal Architecture` adalah metodologi untuk merancang software yang fokus pada `pengorganisasian` kode menjadi `layer-layer` yang independen dan terpisah sesuai dengan fungsi dan tujuannya masing-masing. Kedua arsitektur ini memungkinkan untuk membuat aplikasi yang lebih modular, scalable, dan maintainable.

Clean Architecture memiliki prinsip-prinsip sebagai berikut:

- Menggunakan prinsip `Dependency Inversion Principle` (DIP), yaitu dependency hanya tergantung pada abstraksi bukan implementasi.
- Mengorganisasikan kode menjadi beberapa `layer`, seperti domain, use case, dan delivery.
- Layer-layer harus terisolasi dari layer lainnya dan harus memiliki aturan dan `ketergantungan` yang jelas.
- `Domain layer` berisi `representasi` dari business rules, entitas, dan model-model data.
- `Use case layer` berisi implementasi dari use case dan business logic.
- `Delivery layer` berisi implementasi dari delivery mechanism, seperti HTTP, gRPC, atau message queue.
- `Infrastructure` layer berisi implementasi dari database, external service, atau framework.

Sedangkan Hexagonal Architecture memiliki prinsip-prinsip sebagai berikut:

- Menggunakan `prinsip Dependency Inversion Principle` (DIP), yaitu dependency hanya tergantung pada abstraksi bukan implementasi.
- Memisahkan kode menjadi beberapa layer, seperti domain, application, dan infrastructure.
- Menggunakan port dan adapter untuk menghubungkan antara layer-layer tersebut.
- Domain layer berisi representasi dari business rules, entitas, dan model-model data.
- Application layer berisi implementasi dari use case dan business logic.
- Infrastructure layer berisi implementasi dari database, external service, atau framework.
