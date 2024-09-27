CREATE TABLE article_views (
    id VARCHAR(100) PRIMARY KEY,
    id_user VARCHAR(100),
    id_article VARCHAR(100),
    viewed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY (id_user) REFERENCES users(id),
    CONSTRAINT fk_article FOREIGN KEY (id_article) REFERENCES articles(id)
);
