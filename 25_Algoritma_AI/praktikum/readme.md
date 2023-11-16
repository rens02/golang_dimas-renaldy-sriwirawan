# Tugas

- Terdapat sekumpulan data mengenai tulisan dalam bentuk *tweet* mengenai sebuah kebijakan. Sekumpulan data tersebut ingin dikelompokkan berdasarkan sentimen dari *tweet* tersebut yaitu sentimen positif dan negatif. Jelaskan algoritma A.I. yang dapat digunakan untuk mengelompokkan *tweet* tersebut beserta alasannya.

# Jawab

Untuk mengelompokkan tweet berdasarkan sentimen positif dan negatif dapat menggunakan algoritma Natural Language Processing (NLP) dengan metode klasifikasi teks. Ada beberapa algoritma AI yang efektif untuk tugas ini, dan salah satu yang paling umum digunakan adalah algoritma Naive Bayes, Support Vector Machine (SVM), dan Neural Networks seperti Convolutional Neural Networks (CNN) atau Recurrent Neural Networks (RNN).
 
1. Naive Bayes Classifier

- Naive Bayes Classifier adalah salah satu algoritma klasifikasi yang menghitung probabilitas bahwa sebuah tweet termasuk dalam kategori positif atau negatif. Algoritma ini didasarkan pada teorema Bayes dan mengasumsikan bahwa fitur (kata-kata dalam tweet) adalah independen satu sama lain.
- Alasan : Algoritma ini relatif sederhana dan efisien. Selain itu, ini juga cocok untuk analisis sentimen teks karena dapat mengidentifikasi kata-kata kunci yang mendefinisikan sentimen (misalnya, "bahagia" untuk positif, "sedih" untuk negatif).

2. Support Vector Machine (SVM)

- SVM adalah algoritma yang digunakan untuk memisahkan dua kelas dengan mencari "garis" atau "batas keputusan" terbaik di antara data yang diberikan. Dalam konteks analisis sentimen, SVM dapat digunakan untuk memisahkan tweet menjadi kategori positif dan negatif.
- Alasan : SVM cukup kuat untuk menangani data teks yang kompleks. Ini dapat menangani masalah non-linear dan juga memiliki fleksibilitas dalam pemilihan fungsi kernel yang sesuai untuk data Anda.

3. Recurrent Neural Networks (RNN)

- RNN adalah jenis jaringan saraf yang mampu memahami urutan dalam data, yang sangat berguna dalam pemrosesan teks. Mereka dapat memahami konteks dalam tweet dan mengidentifikasi sentimen berdasarkan urutan kata.
- Alasan : RNN lebih canggih dalam memahami konteks teks. Mereka cocok untuk data teks yang lebih kompleks dan dapat menangani analisis sentimen yang lebih mendalam.