use food_go;
CREATE TABLE IF NOT EXISTS ordered_items(
	food_id INT NOT NULL,
    quantity INT NOT NULL,
    special_instructions VARCHAR(400),
    order_id INT NOT NULL,
    food_status ENUM('completed','left') NOT NULL,
    PRIMARY KEY (food_id,order_id),
    FOREIGN KEY (food_id) REFERENCES food_menu(food_id),
    FOREIGN KEY (order_id) REFERENCES order_table(order_id)
);
