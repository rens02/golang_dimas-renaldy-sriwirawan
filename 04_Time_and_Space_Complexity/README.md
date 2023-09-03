# Soal Time and Space Complexity

## Tugas
 1. Dalam matematika, bilangan prima adalah bilangan asli yang lebih besar dari angka 1, yang faktor pembaginya adalah 1 dan bilangan itu sendiri. 2 dan 3 adalah bilangan prima. 4 bukan bilangan prima karena 4 bisa dibagi 2. Kamu diminta untuk membuat fungsi untuk menentukan bahwa sebuah bilangan termasuk bilangan prima atau tidak.
    
    Buatlah solusi yang lebih optimal, dengan kompleksitas lebih cepat dari O(n) / O(n/2).
    
    ********************************Sample Test Case********************************
    
    Input: 1000000007
    
    Output: Bilangan Prima
    
    Input: 1500450271
    
    Output: Bilangan Prima
     <br>********************************Jawab********************************<br> 
    dioptimalkan menggunakan complexity O(sqrt(n)) karena lebih cepat dari dari O(n) atau O(n/2) dimana untuk pengecekan bilangan prima terdapat beberapa handling, lalu dilakukan looping lagi untuk cek proses penemuan akar kuadrat dari angka yang diberikan [Sours Code](tugas/bilanganprima.go#L22)

    untuk output nya sebagai berikut ini 
    <br>![Alt Text](Screenshot/prima.png)<br>

2. Terdapat dua bilangan integer yaitu x dan n. Buatlah sebuah fungsi untuk melakukan perhitungan perpangkatan (x^n dibaca x pangkat n). Time complexity dari sebuah fungsi perpangkatan harus lebih cepat dari O(n). Contoh time complexity yang optimal adalah logaritmik.
    <br>********************************Jawab********************************<br>  
    [Sours Code](tugas/bilanganinteger.go)
    Output:
    <br>![Alt Text](Screenshot/integer.png)<br>

## Resume
Space complexity pada Golang mengacu pada jumlah memori yang digunakan oleh program ketika dieksekusi. Hal ini tergantung pada banyak faktor, seperti jumlah variabel, tipe data, penggunaan array atau slice, dan alokasi memori secara dinamis. Oleh karena itu, sangat penting untuk mempertimbangkan space complexity saat menulis program Golang untuk memastikan penggunaan memori yang optimal.
