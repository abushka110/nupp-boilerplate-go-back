CREATE TABLE organizations (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    city VARCHAR(255),
    address VARCHAR(255),
    lat FLOAT8,
    lon FLOAT8,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_date TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_organizations_user_id ON organizations(user_id);
CREATE INDEX idx_organizations_deleted_date ON organizations(deleted_date);
