CREATE TABLE IF NOT EXISTS orders(
    Id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    Customer_id INT NOT NULL,
    Order_date Date,
    Status ENUM('pending','in_progress','completed','cancelled') DEFAULT 'pending' NOT NULL,
    Created_at TIMESTAMP DEFAULT NOW(),
    Updated_at TIMESTAMP DEFAULT NOW() ON UPDATE NOW(),
    FOREIGN KEY(Customer_id) REFERENCES customer(Id) ON DELETE CASCADE
);