CREATE TABLE users (
    _id BIGINT NOT NULL AUTO_INCREMENT,
    _uuid VARCHAR(36) NOT NULL,
    user_name VARCHAR(256) UNIQUE NOT NULL,
    first_name VARCHAR(256) NOT NULL,
    last_name VARCHAR(256),
    email VARCHAR(256) NOT NULL,
    phone_number BIGINT NOT NULL,
    active BOOLEAN NOT NULL,
    created_at DATETIME,
    updated_at DATETIME,
    created_by BIGINT,
    updated_by BIGINT,
    PRIMARY KEY (_id),

    INDEX idx_uuid (_uuid),
    INDEX idx_email (email)
);
