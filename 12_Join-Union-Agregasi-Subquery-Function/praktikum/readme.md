# Join, Union, Agregasi, Subquery, Function - DBMS
## Database

```
CREATE DATABASE alta_online_shop;

CREATE TABLE users (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  alamat VARCHAR(255),
  dob DATE,
  status SMALLINT,
  gender CHAR(1),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE operators (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE product_types (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE products (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  product_type_id INT(11),
  operator_id INT(11),
  code VARCHAR(50),
  name VARCHAR(100),
  status SMALLINT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (product_type_id) REFERENCES product_types(id),
  FOREIGN KEY (operator_id) REFERENCES operators(id)
);

CREATE TABLE product_descriptions (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  description TEXT, 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE payment_methods (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  status SMALLINT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT(11),
  payment_method_id INT(11),
  status VARCHAR(10),
  total_qty INT(11),
  total_price NUMERIC(25,2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

CREATE TABLE transaction_details (
  transaction_id INT(11),
  product_id INT(11),
  status VARCHAR(10),
  qty INT(11),
  price NUMERIC(25,2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (transaction_id) REFERENCES transactions(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);
```

## Data Manipulation Language (DML)

-- 1. INSERT
--    a. Insert 5 operators pada tabel operators
```
INSERT INTO operators (name) VALUES
    ('Operator A'),
    ('Operator B'),
    ('Operator C'),
    ('Operator D'),
    ('Operator E');
```

![dml1a](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml1a.JPG)

--    b. Insert 3 product types
```
INSERT INTO product_types (name) VALUES
    ('Mobil'),
    ('Motor'),
    ('Sepeda');
```

![dml1b](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml1b.JPG)

--    c. Insert 2 product dengan product type id = 1, dan operators id = 3
```
INSERT INTO products (product_type_id, operator_id, name, status) VALUES
    (1, 3, 'Mobil hitam', 1),
    (1, 3, 'Mobil putih', 1);
```

![dml1c](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml1c.JPG)

--    d. Insert 3 product dengan product type id = 2, dan operators id = 1
```
INSERT INTO products (product_type_id, operator_id, name, status) VALUES
    (2, 1, 'Motor hitam', 1),
    (2, 1, 'Motor putih', 1),
    (2, 1, 'Motor merah', 1);
```

![dml1d](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml1d.JPG)

--    e. Insert 3 product dengan product type id = 3, dan operators id = 4
```
INSERT INTO products (product_type_id, operator_id, name, status) VALUES
    (3, 4, 'Sepeda hitam', 1),
    (3, 4, 'Sepeda putih', 1),
    (3, 4, 'Sepeda merah', 1);
```

![dml1e](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml1e.JPG)

--    f. Insert product description pada setiap product
```
INSERT INTO product_descriptions (description) VALUES
    ('Mobil berwarna hitam'),
    ('Mobil berwarna putih'),
    ('Motor berwarna hitam'),
    ('Motor berwarna putih'),
    ('Motor berwarna merah'),
    ('Sepeda berwarna hitam'),
    ('Sepeda berwarna putih'),
    ('Sepeda berwarna merah');
```

![dml1f](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml1f.JPG)

--    g. Insert 3 payment methods
```
INSERT INTO payment_methods (name, status) VALUES
    ('E-Wallet', 1),
    ('Virtual Account', 1),
    ('Bank', 1);
```

![dml1g](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml1g.JPG)

--    h. Insert 5 user pada tabel user
```
INSERT INTO users (name, alamat, status, dob, gender) VALUES
    ('Aditya', 'Jakarta', 1, '2002-03-15', 'M'),
    ('Jhonathan', 'Depok', 1, '1999-04-14', 'M'),
    ('Vera', 'Tangerang', 1, '1990-01-30', 'F'),
    ('Luna', 'Jakarta', 1, '2001-11-10', 'F'),
    ('Tomo', 'Karawaci', 1, '1995-09-20', 'M');
```

![dml1h](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml1h.JPG)

--    i. Insert 3 transaksi di masing-masing user
```
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price) VALUES
    -- Transaksi untuk User 1
    (1, 1, 'Completed', 3, 300000.00),
    (1, 1, 'Pending', 3, 300000.00),
    (1, 3, 'Completed', 3, 300000.00),

    -- Transaksi untuk User 2
    (2, 3, 'Completed', 3, 300000.00),
    (2, 2, 'Completed', 3, 300000.00),
    (2, 3, 'Pending', 3, 300000.00),

    -- Transaksi untuk User 3
    (3, 1, 'Pending', 3, 300000.00),
    (3, 2, 'Pending', 3, 300000.00),
    (3, 3, 'Pending', 3, 300000.00),

    -- Transaksi untuk User 4
    (4, 1, 'Pending', 3, 300000.00),
    (4, 3, 'Completed', 3, 300000.00),
    (4, 3, 'Completed', 3, 300000.00),

    -- Transaksi untuk User 5
    (5, 2, 'Completed', 3, 300000.00),
    (5, 2, 'Completed', 3, 300000.00),
    (5, 3, 'Completed', 3, 300000.00);
```

![dml1i](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml1i.JPG)

--    j. Insert 3 produk di masing-masing transaksi:
```
INSERT INTO transaction_details (transaction_id, product_id, status, qty, price) VALUES
    -- Produk untuk Transaksi User 1, Transaksi 1
    (1, 1, 'Completed', 1, 100000.00),
    (1, 1, 'Completed', 1, 100000.00),
    (1, 3, 'Completed', 1, 100000.00),

    -- Produk untuk Transaksi User 1, Transaksi 2
    (2, 5, 'Pending', 1, 100000.00),
    (2, 2, 'Pending', 1, 100000.00),
    (2, 8, 'Pending', 1, 100000.00),

    -- Produk untuk Transaksi User 1, Transaksi 3
    (3, 4, 'Pending', 1, 100000.00),
    (3, 2, 'Completed', 1, 100000.00),
    (3, 3, 'Completed', 1, 100000.00),

    -- Produk untuk Transaksi User 2, Transaksi 1
    (4, 8, 'Completed', 1, 100000.00),
    (4, 3, 'Pending', 1, 100000.00),
    (4, 2, 'Pending', 1, 100000.00),

    -- Produk untuk Transaksi User 2, Transaksi 2
    (5, 1, 'Completed', 1, 100000.00),
    (5, 4, 'Completed', 1, 100000.00),
    (5, 3, 'Completed', 1, 100000.00),

    -- Produk untuk Transaksi User 2, Transaksi 3
    (6, 2, 'Pending', 1, 100000.00),
    (6, 5, 'Pending', 1, 100000.00),
    (6, 6, 'Pending', 1, 100000.00),

    -- Produk untuk Transaksi User 3, Transaksi 1
    (7, 1, 'Pending', 1, 100000.00),
    (7, 8, 'Pending', 1, 100000.00),
    (7, 7, 'Completed', 1, 100000.00),

-- Produk untuk Transaksi User 3, Transaksi 2
    (8, 7, 'Completed', 1, 100000.00),
    (8, 8, 'Completed', 1, 100000.00),
    (8, 1, 'Pending', 1, 100000.00),

-- Produk untuk Transaksi User 3, Transaksi 3
    (9, 4, 'Pending', 1, 100000.00),
    (9, 3, 'Completed', 1, 100000.00),
    (9, 2, 'Pending', 1, 100000.00),

-- Produk untuk Transaksi User 4, Transaksi 1
    (10, 5, 'Pending', 1, 100000.00),
    (10, 2, 'Pending', 1, 100000.00),
    (10, 3, 'Pending', 1, 100000.00),

-- Produk untuk Transaksi User 4, Transaksi 2
    (11, 4, 'Completed', 1, 100000.00),
    (11, 1, 'Completed', 1, 100000.00),
    (11, 3, 'Pending', 1, 100000.00),

-- Produk untuk Transaksi User 4, Transaksi 3
    (12, 1, 'Completed', 1, 100000.00),
    (12, 2, 'Completed', 1, 100000.00),
    (12, 3, 'Completed', 1, 100000.00),

-- Produk untuk Transaksi User 5, Transaksi 1
    (13, 6, 'Completed', 1, 100000.00),
    (13, 7, 'Pending', 1, 100000.00),
    (13, 8, 'Completed', 1, 100000.00),

-- Produk untuk Transaksi User 5, Transaksi 2
    (14, 5, 'Pending', 1, 100000.00),
    (14, 3, 'Completed', 1, 100000.00),
    (14, 4, 'Completed', 1, 100000.00),

-- Produk untuk Transaksi User 5, Transaksi 3
    (15, 5, 'Completed', 1, 100000.00),
    (15, 6, 'Completed', 1, 100000.00),
    (15, 7, 'Completed', 1, 100000.00);
```

![dml1j](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml1j.JPG)

-- 2. SELECT
-- a. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
```
SELECT name, gender
FROM users
WHERE gender = 'M';
```

![dml2a](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml2a.JPG)

-- b. Tampilkan product dengan id = 3.
```
SELECT *
FROM products
WHERE id = 3;
```

![dml2b](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml2b.JPG)

-- c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata 'a'.
```
SELECT *
FROM users
WHERE created_at >= NOW() - INTERVAL 7 DAY AND name LIKE '%a%';
```

![dml2c](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml2c.JPG)

-- d. Hitung jumlah user / pelanggan dengan status gender Perempuan.
```
SELECT COUNT(*)
FROM users
WHERE gender = 'F';
```

![dml2d](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml2d.JPG)

-- e. Tampilkan data pelanggan dengan urutan sesuai nama abjad.
```
SELECT *
FROM users
ORDER BY name;
```

![dml2e](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml2e.JPG)

-- f. Tampilkan 5 data pada tabel product.
```
SELECT *
FROM products
LIMIT 5;
```

![dml2f](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml2f.JPG)

-- 3. UPDATE
-- a. Ubah data product id 1 dengan nama 'product dummy':
```
UPDATE products
SET name = 'product dummy'
WHERE id = 1;
```

![dml3a](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml3a.JPG)

-- b. Update qty = 3 pada transaction detail dengan product id = 1:
```
UPDATE transaction_details
SET qty = 3
WHERE product_id = 1;
```

![dml3b](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml3b.JPG)

-- 4. DELETE
-- a. Delete data pada tabel product dengan id = 1:
```
DELETE FROM products
WHERE id = 1;
```

![dml4a](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml4a.JPG)

-- b. Delete pada tabel product dengan product type id = 1:
```
DELETE FROM products
WHERE product_type_id = 1;
```

![dml4b](/12_Join-Union-Agregasi-Subquery-Function/screenshots/dml4b.JPG)

## Join, Union, Sub query, Function

-- 1. Gabungkan data transaksi dari user id 1 dan user id 2:
```
SELECT *
FROM transactions
WHERE user_id IN (1, 2);
```

![join1](/12_Join-Union-Agregasi-Subquery-Function/screenshots/join1.JPG)

-- 2. Tampilkan jumlah harga transaksi user id 1:
```
SELECT user_id, SUM(total_price) AS total_harga
FROM transactions
WHERE user_id = 1
GROUP BY user_id;
```

![join2](/12_Join-Union-Agregasi-Subquery-Function/screenshots/join2.JPG)

-- 3. Tampilkan total transaksi dengan product type 2:
```
SELECT pt.name AS product_type, SUM(t.price) AS total_harga
FROM transaction_details t
JOIN products p ON t.product_id = p.id
JOIN product_types pt ON p.product_type_id = pt.id
WHERE pt.id = 2
GROUP BY pt.name;
```

![join3](/12_Join-Union-Agregasi-Subquery-Function/screenshots/join3.JPG)

-- 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan:
```
SELECT p.*, pt.name AS product_type_name
FROM products p
JOIN product_types pt ON p.product_type_id = pt.id;

CREATE OR REPLACE VIEW transaction_with_product_id AS
SELECT t.*, td.product_id
FROM transactions t
JOIN transaction_details td ON t.id = td.transaction_id;
```

![join4](/12_Join-Union-Agregasi-Subquery-Function/screenshots/join4.JPG)

-- 5. Tampilkan semua field table transaction, field name table product dan field name table user:
```
SELECT t.*, p.name AS products_name, u.name AS user_name
FROM transaction_with_product_id t
JOIN products p ON t.product_id = p.id
JOIN users u ON t.user_id = u.id;
```

![join5a](/12_Join-Union-Agregasi-Subquery-Function/screenshots/join5a.JPG)

![join5b](/12_Join-Union-Agregasi-Subquery-Function/screenshots/join5b.JPG)

-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud:
```
DELIMITER //

CREATE TRIGGER after_transaction_delete
AFTER DELETE ON transactions
FOR EACH ROW
BEGIN
	DELETE FROM transaction_details
	WHERE transaction_id = OLD.id;
END;
//

DELIMITER ;
```

![join6](/12_Join-Union-Agregasi-Subquery-Function/screenshots/join6.JPG)

-- 7.Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus:
```
DELIMITER //

CREATE TRIGGER after_transaction_detail_delete
AFTER DELETE ON transaction_details
FOR EACH ROW
BEGIN
	DECLARE trans_id INT(11);
	SET trans_id = OLD.transaction_id;
	
	UPDATE transactions
	SET total_qty = total_qty - OLD.qty
	WHERE id = trans_id;
END;
//

DELIMITER ;
```

![join7](/12_Join-Union-Agregasi-Subquery-Function/screenshots/join7.JPG)

-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query:
```
SELECT *
FROM products
WHERE id NOT IN (
    SELECT DISTINCT product_id
    FROM transaction_details
);
```

![join8](/12_Join-Union-Agregasi-Subquery-Function/screenshots/join8.JPG)