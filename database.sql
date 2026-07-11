create table users (
    id int auto_increment,
	email varchar(100) not null,
	password varchar(100) not null,
	name varchar(100) not null,
	role ENUM('admin', 'user') not null default 'user',
    createdAt timestamp default current_timestamp,
    updatedAt timestamp default current_timestamp on update current_timestamp,
    PRIMARY KEY (id),
    CONSTRAINT UNIQUE_email UNIQUE (email) 
);


CREATE TABLE refresh_tokens (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    token VARCHAR(255) NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);