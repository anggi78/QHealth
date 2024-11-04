CREATE TABLE queue (
    id varchar(100) PRIMARY KEY,
    queue_number VARCHAR(50) NOT NULL,
    rest_of_the_queue VARCHAR(100) NOT NULL,
    id_user VARCHAR(100),
    id_doctor VARCHAR(100),
    id_queue_status VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (id_user) REFERENCES users(id),
    CONSTRAINT fk_doctor FOREIGN KEY (id_doctor) REFERENCES doctors(id),
    CONSTRAINT fk_queue_status FOREIGN KEY (id_queue_status) REFERENCES queue_status(id)
)