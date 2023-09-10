 # Soal Concurrent Programming (eksplorasi)

## Soal Prioritas 1 
1. Buatlah sebuah fungsi yang mencetak angka kelipatan x, dimana x merupakan parameter bilangan bulat positif. lalu jalankan fungsi tersebut dengan menerapkan goroutine, dengan interval waktu 3 detik! 
   <br>********************************Jawab :  [Source Code](Praktikum/multiplex.go)********************************   
   Output 
   <br>![Alt Text](Screenshot/multiplex.png)<br>

2. Buatlah program yang mencetak bilangan kelipatan 3 dengan menerapkan goroutine dan channel!   
   <br>********************************Jawab :  [Source Code](Praktikum/multiplechanel.go)********************************   
   Output 
   <br>![Alt Text](Screenshot/multiplechannel.png)<br>

3. Buatlah program yang mencetak bilangan kelipatan 3 dengan menerapkan goroutine dan buffer channel!
   <br>********************************Jawab :  [Source Code](Praktikum/multiplesbuferchanel.go)********************************   
   Output 
   <br>![Alt Text](Screenshot/multiplechannel.png)<br>

4. Buatlah  program yang yang menerapkan mutex! Jenis program yang dibuat bebas (contoh: perhitungan faktorial)
    <br>********************************Jawab :  [Source Code](Praktikum/incrementCounter.go)********************************   
   Output 
   <br>![Alt Text](Screenshot/mutex.png)<br>



### Resume
Concurrent programming di Golang merupakan kemampuan untuk menjalankan beberapa Praktikum secara simultan atau paralel pada waktu yang sama. Dalam Golang, concurrent programming dapat dilakukan dengan menggunakan goroutines, channels, dan select.

Goroutines adalah lightweight threads yang dijalankan oleh Go runtime, yang memungkinkan aplikasi untuk menjalankan banyak Praktikum dalam waktu yang bersamaan tanpa perlu mengalami overhead yang besar seperti pada thread konvensional. Goroutines juga dapat berkomunikasi satu sama lain melalui channels, yang berfungsi sebagai media untuk mengirim dan menerima data antar goroutines.

Pada Golang, penggunaan channels sangat penting dalam melakukan concurrent programming. Channels memungkinkan pengiriman data secara aman antara goroutines yang berjalan secara simultan. Pengiriman data antar goroutines melalui channels diatur dengan menggunakan konsep select.

Dalam praktiknya, concurrent programming di Golang digunakan untuk meningkatkan kinerja dan efisiensi dari aplikasi yang dijalankan. Dalam mengimplementasikan concurrent programming di Golang, diperlukan pemahaman yang baik tentang penggunaan goroutines, channels, dan select agar dapat mengoptimalkan performa dan meminimalkan masalah deadlock atau race condition yang sering terjadi pada program yang dijalankan secara paralel.
