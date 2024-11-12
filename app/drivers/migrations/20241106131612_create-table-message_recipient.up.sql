CREATE TABLE message_recipients (
    id VARCHAR(100) PRIMARY KEY,
    is_read BOOLEAN,
    id_message VARCHAR,
    id_user VARCHAR,
    id_doctor VARCHAR,
    CONSTRAINT fk_message FOREIGN KEY (id_message) REFERENCES messages(id) ON DELETE CASCADE,
    CONSTRAINT fk_recipient_user FOREIGN KEY (id_user) REFERENCES users(id),
    CONSTRAINT fk_doctor FOREIGN KEY (id_doctor) REFERENCES doctors(id)
)