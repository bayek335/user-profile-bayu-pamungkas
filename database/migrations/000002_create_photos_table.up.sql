CREATE TABLE  IF NOT EXISTS photos(
    id SERIAL, 
    title VARCHAR(64) NOT NULL,
    caption VARCHAR(64) NOT NULL,
    user_id INT NOT NULL,
    photo_url VARCHAR(128) NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);