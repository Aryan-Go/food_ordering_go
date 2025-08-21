use food_go;
CREATE TABLE IF NOT EXISTS food_menu(
	food_id INT PRIMARY KEY AUTO_INCREMENT,
    food_name VARCHAR(200) NOT NULL,
    description VARCHAR(400),
    price DECIMAL NOT NULL,
    category_id INT,
    FOREIGN KEY (category_id) REFERENCES category(category_id)
);

INSERT INTO food_menu (food_name,description,price,category_id) VALUES ("Paneer Tikka" , "Paneer Tikka is a popular Indian appetizer made from marinated paneer cubes grilled or roasted to perfection with spices and vegetables.",300.00,1);
INSERT INTO food_menu (food_name,description,price,category_id) VALUES ("Dahi Kebab" , "Dahi Kebab is a soft and creamy North Indian snack made from hung curd, paneer, and mild spices, lightly pan-fried for a crispy outer layer.",250.00,1);
INSERT INTO food_menu (food_name,description,price,category_id) VALUES ("Sushi" , "Sushi is a traditional Japanese dish made with vinegared rice, often paired with raw or cooked seafood, vegetables, and wrapped in seaweed.",450.00,1);

INSERT INTO food_menu (food_name,description,price,category_id) VALUES ("Thai Curry momos" , "Veg Thai curry momos are soft dumplings stuffed with seasoned vegetables, served in a creamy, mildly spiced Thai curry made with coconut milk, lemongrass, and herbs—a perfect fusion of comfort and flavor.",400.00,2);
INSERT INTO food_menu (food_name,description,price,category_id) VALUES ("Fajitas" , "Veg fajitas are a sizzling Tex-Mex dish made with sautéed bell peppers, onions, and spiced vegetables, served with warm tortillas and toppings like salsa, guacamole, and cheese for a flavorful wrap experience.",500.00,2);
INSERT INTO food_menu (food_name,description,price,category_id) VALUES ("Napolean Pizza" , "Neapolitan pizza (also spelled Napoleon pizza) is a traditional Italian pizza known for its thin, soft, and chewy crust, topped with fresh tomato sauce, mozzarella cheese, basil, and a drizzle of olive oil—baked quickly at high heat for a perfect char.",500.00,2);

INSERT INTO food_menu (food_name,description,price,category_id) VALUES ("Hot Chocolate foutain + icecream" , "A Hot Chocolate Fountain with Ice Cream is an indulgent dessert experience where warm, flowing chocolate is paired with scoops of creamy ice cream—perfect for dipping, drizzling, or just enjoying the rich contrast of hot and cold in every bite.",200.00,3);
INSERT INTO food_menu (food_name,description,price,category_id) VALUES ("Gulab Jamun" , "Gulab Jamun is a classic Indian dessert made of soft, deep-fried milk solids soaked in fragrant sugar syrup, known for its rich sweetness and melt-in-the-mouth texture.",200.00,3);
INSERT INTO food_menu (food_name,description,price,category_id) VALUES ("Kulfi Faluda" , "Kulfi Faluda is a traditional Indian dessert that combines creamy, dense kulfi (frozen milk-based treat) with sweet, chewy faluda noodles, basil seeds, rose syrup, and chilled milk — offering a refreshing and indulgent experience.",150.00,3);