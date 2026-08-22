CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ
);
CREATE INDEX idx_users_deleted_at ON users (deleted_at);

CREATE TABLE organizations (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ
);
CREATE INDEX idx_organizations_deleted_at ON organizations (deleted_at);

CREATE TABLE org_members (
    id BIGSERIAL PRIMARY KEY,
    organization_id BIGINT NOT NULL REFERENCES organizations(id),
    user_id BIGINT NOT NULL REFERENCES users(id),
    role VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    UNIQUE (organization_id, user_id)
);
CREATE INDEX idx_org_members_org_id ON org_members (organization_id);
CREATE INDEX idx_org_members_user_id ON org_members (user_id);