-- DATA MANIPULATION LANGUAGE (DML)
-- 1. INSERT
--    a. Insert 5 operators pada tabel operators
INSERT INTO operators (name) VALUES
    ('Operator A'),
    ('Operator B'),
    ('Operator C'),
    ('Operator D'),
    ('Operator E');

--    b. Insert 3 product types
INSERT INTO product_types (name) VALUES
    ('Mobil'),
    ('Motor'),
    ('Sepeda');

--    c. Insert 2 product dengan product type id = 1, dan operators id = 3
INSERT INTO products (product_type_id, operator_id, name, status) VALUES
    (1, 3, 'Mobil hitam', 1),
    (1, 3, 'Mobil putih', 1);

--    d. Insert 3 product dengan product type id = 2, dan operators id = 1
INSERT INTO products (product_type_id, operator_id, name, status) VALUES
    (2, 1, 'Motor hitam', 1),
    (2, 1, 'Motor putih', 1),
    (2, 1, 'Motor merah', 1);

--    e. Insert 3 product dengan product type id = 3, dan operators id = 4
INSERT INTO products (product_type_id, operator_id, name, status) VALUES
    (3, 4, 'Sepeda hitam', 1),
    (3, 4, 'Sepeda putih', 1),
    (3, 4, 'Sepeda merah', 1);

--    f. Insert product description pada setiap product
INSERT INTO product_descriptions (description) VALUES
    ('Mobil berwarna hitam'),
    ('Mobil berwarna putih'),
    ('Motor berwarna hitam'),
    ('Motor berwarna putih'),
    ('Motor berwarna merah'),
    ('Sepeda berwarna hitam'),
    ('Sepeda berwarna putih'),
    ('Sepeda berwarna merah');

--    g. Insert 3 payment methods
INSERT INTO payment_methods (name, status) VALUES
    ('E-Wallet', 1),
    ('Virtual Account', 1),
    ('Bank', 1);

--    h. Insert 5 user pada tabel user
INSERT INTO users (name, alamat, status, dob, gender) VALUES
    ('Aditya', 'Jakarta', 1, '2002-03-15', 'M'),
    ('Dimas', 'Depok', 1, '1999-04-14', 'M'),
    ('Vera', 'Tangerang', 1, '1990-01-30', 'F'),
    ('Luna', 'Jakarta', 1, '2001-11-10', 'F'),
    ('Tomo', 'Karawaci', 1, '1995-09-20', 'M');

--    i. Insert 3 transaksi di masing-masing user
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

--    j. Insert 3 produk di masing-masing transaksi:
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

-- 2. SELECT
-- a. Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT name, gender
FROM users
WHERE gender = 'M';

-- b. Tampilkan product dengan id = 3.
SELECT *
FROM products
WHERE id = 3;

-- c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata 'a'.
SELECT *
FROM users
WHERE created_at >= NOW() - INTERVAL 7 DAY AND name LIKE '%a%';

-- d. Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*)
FROM users
WHERE gender = 'F';

-- e. Tampilkan data pelanggan dengan urutan sesuai nama abjad.
SELECT *
FROM users
ORDER BY name;

-- f. Tampilkan 5 data pada tabel product.
SELECT *
FROM products
LIMIT 5;

-- 3. UPDATE
-- a. Ubah data product id 1 dengan nama 'product dummy':
UPDATE products
SET name = 'product dummy'
WHERE id = 1;

-- b. Update qty = 3 pada transaction detail dengan product id = 1:
UPDATE transaction_details
SET qty = 3
WHERE product_id = 1;

-- 4. DELETE
-- a. Delete data pada tabel product dengan id = 1:
DELETE FROM products
WHERE id = 1;

-- b. Delete pada tabel product dengan product type id = 1:
DELETE FROM products
WHERE product_type_id = 1;

-- Join, Union, Sub query, Function
-- 1. Gabungkan data transaksi dari user id 1 dan user id 2:
SELECT *
FROM transactions
WHERE user_id IN (1, 2);

-- 2. Tampilkan jumlah harga transaksi user id 1:
SELECT user_id, SUM(total_price) AS total_harga
FROM transactions
WHERE user_id = 1
GROUP BY user_id;

-- 3. Tampilkan total transaksi dengan product type 2:
SELECT pt.name AS product_type, SUM(t.price) AS total_harga
FROM transaction_details t
JOIN products p ON t.product_id = p.id
JOIN product_types pt ON p.product_type_id = pt.id
WHERE pt.id = 2
GROUP BY pt.name;

-- 4. Tampilkan semua field table product dan field name table product type yang saling berhubungan:
SELECT p.*, pt.name AS product_type_name
FROM products p
JOIN product_types pt ON p.product_type_id = pt.id;

CREATE OR REPLACE VIEW transaction_with_product_id AS
SELECT t.*, td.product_id
FROM transactions t
JOIN transaction_details td ON t.id = td.transaction_id;

-- 5. Tampilkan semua field table transaction, field name table product dan field name table user:
SELECT t.*, p.name AS products_name, u.name AS user_name
FROM transaction_with_product_id t
JOIN products p ON t.product_id = p.id
JOIN users u ON t.user_id = u.id;

-- 6. Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud:
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

-- 7.Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus:
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


-- 8. Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query:
SELECT *
FROM products
WHERE id NOT IN (
    SELECT DISTINCT product_id
    FROM transaction_details
);
