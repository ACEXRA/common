CREATE TABLE group_table (
    _id BIGINT NOT NULL AUTO_INCREMENT,
    _uuid VARCHAR(36) NOT NULL,
    group_name VARCHAR(256) UNIQUE NOT NULL,
    active BOOLEAN NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    created_by BIGINT,
    updated_by BIGINT,
    PRIMARY KEY (_id),

    INDEX idx_uuid (_uuid),
    INDEX idx_group_name (group_name)
);