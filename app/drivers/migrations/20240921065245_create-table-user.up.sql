CREATE TABLE users (
    id varchar(100) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    image VARCHAR(255),
    address VARCHAR(255),
    phone VARCHAR(255) NOT NULL,
    birth DATE,
    jk VARCHAR(255),
    nik VARCHAR(255),
    image_ktp VARCHAR(255),
    id_role VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
)