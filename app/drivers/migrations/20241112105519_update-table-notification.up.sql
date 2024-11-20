CREATE TABLE notifications (
    id VARCHAR(100) PRIMARY KEY,
    type VARCHAR(50),
    message TEXT,
    is_read BOOLEAN DEFAULT false,
    id_user VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (id_user) REFERENCES users(id)
);