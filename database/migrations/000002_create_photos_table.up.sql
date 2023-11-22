CREATE TABLE  IF NOT EXISTS photos(
    id SERIAL, 
    title VARCHAR(64) NOT NULL,
    caption VARCHAR(64) UNIQUE NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);