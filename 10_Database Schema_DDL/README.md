# # Praktikum Soal Database Schema, DDL, DML


### Merancang Skema Database

1. Sistem dapat menyimpan data mengenai detail item product, yaitu : product, product_type, product_description, operator, payment_methods.
2. Sistem juga harus menyimpan data mengenai pelanggan yang akan membeli product tsb diantaranya : nama, alamat, tanggal lahir, status_user, gender, created_at, updated_at.
3. Sistem dapat mencatat transaksi pembelian dari pelanggan.
4. Sistem dapat mencatat detail transaksi pembelian dari pelanggan.<br>
 Outuput 
    <br>![Alt Text](Screenshot/schema.jpeg)<br>

### Data Definition Language (DDL)

1. Create database alta_online_shop.
<br>![Alt Text](Screenshot/1.png)<br>
2. Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.
    1. Create table user.
    2. Create table product, product type, operators, product description, payment_method.
    3. Create table transaction, transaction detail.
    <br>![Alt Text](Screenshot/2.png)<br>
3. Create tabel kurir dengan field id, name, created_at, updated_at.
<br>![Alt Text](Screenshot/3.png)<br>
4. Tambahkan ongkos_dasar column di tabel kurir.
<br>![Alt Text](Screenshot/4.png)<br>
5. Rename tabel kurir menjadi shipping.
<br>![Alt Text](Screenshot/5.png)<br>
6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
<br>![Alt Text](Screenshot/6.png)<br>
7. Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many. Seperti:
    1. 1-to-1: payment method description.
    2. 1-to-many: user dengan alamat.
    3. many-to-many: user dengan payment method menjadi user_payment_method_detail.
    <br>![Alt Text](Screenshot/7.png)<br>
<br>********************************Jawab :  [Source Code](Praktikum/alta_online_shop.sql)********************************   
    Outuput dari erd query di atas
    <br>![Alt Text](Screenshot/hasil.png)<br>