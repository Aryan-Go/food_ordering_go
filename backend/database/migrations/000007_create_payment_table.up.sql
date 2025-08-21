use food_go;
CREATE TABLE IF NOT EXISTS payment_table(
	total_price DECIMAL,
    tip DECIMAL,
    payment_status ENUM('completed','left') NOT NULL,
    order_id INT NOT NULL,
    customer_id INT NOT NULL,
    payment_id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    FOREIGN KEY (customer_id) REFERENCES user(user_id),
    FOREIGN KEY (order_id) REFERENCES order_table(order_id)
);