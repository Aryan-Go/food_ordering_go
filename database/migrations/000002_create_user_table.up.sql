use food_go
CREATE TABLE user(
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(200) NOT NULL ,
    username VARCHAR(300),
    password VARCHAR(200) NOT NULL,
    role ENUM('chef','customer','admin')
);
INSERT INTO user (email,username,password,role) VALUES("admin@gmail.com","admin","$2a$10$szHaFxbefg73iCmJmxm3J.wiQvhLKPtsUUSgTILi7lZraS8y1y72y","admin");
INSERT INTO user (email,username,password,role) VALUES("chef@gmail.com","chef","$2a$10$CjFpMrh7fguKKVLqeB.a6uZ1dacrq.vmMTFH6MG8CG0WLh8PQc4IG","chef");
