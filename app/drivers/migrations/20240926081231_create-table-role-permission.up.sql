CREATE TABLE role_permission (
    id varchar(100) PRIMARY KEY,
    can_create BOOLEAN DEFAULT FALSE, 
    can_read BOOLEAN DEFAULT FALSE,    
    can_edit BOOLEAN DEFAULT FALSE,   
    can_delete BOOLEAN DEFAULT FALSE,
    id_role VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_role FOREIGN KEY (id_role) REFERENCES role(id)
);
