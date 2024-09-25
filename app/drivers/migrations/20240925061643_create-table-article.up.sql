CREATE TABLE articles (
    id VARCHAR(100) PRIMARY KEY,
    writer VARCHAR(100) NOT NULL,
    title VARCHAR(50) NOT NULL,
    content VARCHAR(100) NOT NULL,
    image VARCHAR(255) NOT NULL,
    date DATE,
    status INT,
    id_user VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (id_user) REFERENCES users(id)
);
