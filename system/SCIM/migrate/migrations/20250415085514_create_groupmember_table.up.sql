CREATE TABLE group_member (
    _id BIGINT NOT NULL AUTO_INCREMENT,
    _uuid VARCHAR(36) NOT NULL,
    user_id BIGINT,
    group_id BIGINT,
    created_at DATETIME,
    updated_at DATETIME,
    created_by BIGINT,
    updated_by BIGINT,
    PRIMARY KEY (_id),
    FOREIGN KEY (user_id) REFERENCES users(_id) ON DELETE CASCADE,
    FOREIGN KEY (group_id) REFERENCES group_table(_id) ON DELETE CASCADE
);