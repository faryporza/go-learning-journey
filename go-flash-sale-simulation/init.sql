CREATE TABLE products (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    stock INT
);

CREATE TABLE orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_id INT,
    user_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- จำลองสินค้า iPhone 16 ลดราคาเหลือ 1 บาท มีแค่ 10 เครื่อง!
INSERT INTO products (id, name, stock) VALUES (1, 'iPhone 16 Pro Max (1 THB)', 10);