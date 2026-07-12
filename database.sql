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

create table accounts (
    id          int      auto_increment,
    user_id     int      not null,
    account_bank varchar(100) not null,
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    account_type enum('savings', 'checking') null,
    createdAt   timestamp default current_timestamp,
    updatedAt timestamp default current_timestamp on update current_timestamp,

    PRIMARY KEY (id),
    CONSTRAINT `Account_users_id_fkey` FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    
)

create table transactions (
    id            int      auto_increment,
    from_account_id int null,
    to_account_id   int null,
    amount        DECIMAL(15, 2) null,
    type          enum('transfer', 'deposit', 'withdrawal') not null,
    description   varchar(100) null,
    createdAt     timestamp default current_timestamp,

    PRIMARY KEY (id),
    CONSTRAINT `Transaction_from_account_id_fkey` FOREIGN KEY (from_account_id) REFERENCES accounts(id)
        ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT `Transaction_to_account_id_fkey` FOREIGN KEY (to_account_id) REFERENCES accounts(id)
        ON DELETE SET NULL ON UPDATE CASCADE
)