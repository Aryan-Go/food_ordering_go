use food_go
CREATE TABLE category(
    category_id INT PRIMARY KEY AUTO_INCREMENT,
    category_name VARCHAR(100) NOT NULL,
    category_desc VARCHAR(400) NOT NULL
);
INSERT INTO category (category_name,category_desc) VALUES ("starters","A tempting medley of crispy, spicy, and savory bites that ignite your appetite and set the perfect tone for the feast ahead.");
INSERT INTO category (category_name,category_desc) VALUES ("main course","A hearty and flavorful spread of rich curries, sizzling stir-fries, and comforting classics crafted to satisfy every craving and steal the spotlight.");
INSERT INTO category (category_name,category_desc) VALUES ("desert","An indulgent finale of sweet delights, where every bite melts in your mouth and leaves a lingering taste of bliss.");
