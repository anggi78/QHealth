CREATE TABLE patients (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(50),
    age INT,
    diagnosis VARCHAR,
    category VARCHAR,
    priority VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);