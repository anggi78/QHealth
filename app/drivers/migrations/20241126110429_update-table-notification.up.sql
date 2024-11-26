CREATE TABLE notifications (
    id VARCHAR(100) PRIMARY KEY,
    type VARCHAR(50),
    message TEXT,
    is_read BOOLEAN DEFAULT false,
    recipient_type VARCHAR(50),
    recipient_id VARCHAR,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);