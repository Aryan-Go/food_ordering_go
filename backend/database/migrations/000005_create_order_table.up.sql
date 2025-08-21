use food_go;
CREATE TABLE IF NOT EXISTS order_table(
	order_id INT PRIMARY KEY auto_increment,
    customer_id INT NOT NULL,
    food_status ENUM('completed','left') NOT NULL,
    chef_id INT NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES user(user_id),
    FOREIGN KEY (chef_id) REFERENCES user(user_id)
);