CREATE TABLE messages (
    id VARCHAR(100) PRIMARY KEY,
    message_body TEXT,
    create_date DATE DEFAULT CURRENT_DATE,
    id_parent_message VARCHAR,
    id_user VARCHAR,
    id_doctor VARCHAR,
    CONSTRAINT fk_user FOREIGN KEY (id_user) REFERENCES users(id),
    CONSTRAINT fk_doctor FOREIGN KEY (id_doctor) REFERENCES doctors(id),
    CONSTRAINT fk_parent_message FOREIGN KEY (id_parent_message) REFERENCES messages(id)
)