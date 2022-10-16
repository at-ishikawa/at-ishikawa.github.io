-- Grant a root to dump data
-- CREATE USER 'root'@'%' IDENTIFIED BY 'password';
-- GRANT ALL PRIVILEGES ON *.* TO 'root'@'%';

-- Create a replication user
CREATE USER 'repl'@'%' IDENTIFIED BY 'password';
GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%';

-- Create a monitor user
CREATE USER 'monitor'@'%' IDENTIFIED BY 'monitor_password';
GRANT ALL PRIVILEGES ON *.* TO 'monitor'@'%';

FLUSH PRIVILEGES;

USE test;
CREATE TABLE users (
    id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    PRIMARY KEY (id),
    UNIQUE(name)
);
