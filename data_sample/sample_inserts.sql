-- Inserts for the users table
INSERT INTO users (username, email, password_hash) VALUES
('alice_smith', 'alice.smith@example.com', 'hashed_password_alice'),
('bob_jones', 'bob.jones@example.com', 'hashed_password_bob'),
('emma_doe', 'emma.doe@example.com', 'hashed_password_emma');

-- Inserts for the orders table
INSERT INTO orders (user_id, total, created_at) VALUES
(1, 99.99, '2024-04-25 09:00:00'),
(2, 49.50, '2024-04-24 14:30:00'),
(3, 199.95, '2024-04-23 11:45:00');

-- Inserts for the items table
INSERT INTO items (name, price, description) VALUES
('T-shirt', 19.99, 'Cotton T-shirt with logo'),
('Jeans', 39.99, 'Blue denim jeans'),
('Sneakers', 59.95, 'White canvas sneakers');

-- Inserts for the order_items table
INSERT INTO order_items (order_id, item_id, quantity) VALUES
(1, 1, 2),
(1, 3, 1),
(2, 2, 1),
(3, 1, 3),
(3, 2, 2),
(3, 3, 1);
