-- 注意：postgres 用户通常已经存在，是默认的超级用户
-- 如果 postgres 用户不存在，可以取消下面这行的注释
-- CREATE USER postgres WITH SUPERUSER PASSWORD 'strong_password_for_postgres';

-- 创建一个新的用户
CREATE USER test_user WITH PASSWORD 'strong_password_for_test_user' CREATEDB;

-- 创建数据库
CREATE DATABASE meow_user WITH OWNER postgres;

-- 连接到新创建的数据库
\c meow_user

-- 设置字符集和时区
SET client_encoding = 'UTF8';
SET timezone = 'UTC';

-- 创建必要的扩展
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- 创建触发器函数来自动更新 created_at 和 updated_at
CREATE OR REPLACE FUNCTION update_timestamp_columns()
    RETURNS TRIGGER AS
$$
BEGIN
    IF TG_OP = 'INSERT' THEN
        NEW.created_at = NOW();
    END IF;
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 创建用户表
CREATE TABLE users
(
    id                 UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    created_at         TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at         TIMESTAMP WITH TIME ZONE,
    is_logic_deleted   BOOLEAN                  DEFAULT FALSE,
    version            INTEGER                  DEFAULT 1,
    username           VARCHAR(50) UNIQUE  NOT NULL,
    email              VARCHAR(100) UNIQUE NOT NULL,
    password           VARCHAR(255)        NOT NULL,
    first_name         VARCHAR(50),
    last_name          VARCHAR(50),
    avatar             VARCHAR(255),
    bio                TEXT,
    date_of_birth      DATE,
    phone_number       VARCHAR(20),
    last_login_at      TIMESTAMP WITH TIME ZONE,
    is_verified        BOOLEAN                  DEFAULT FALSE,
    role               VARCHAR(20)              DEFAULT 'user',
    preferred_language VARCHAR(10)              DEFAULT 'en',
    CONSTRAINT check_role CHECK (role IN ('user', 'admin', 'moderator'))
);

-- 创建Card表
CREATE TABLE cards
(
    id               UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    created_at       TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at       TIMESTAMP WITH TIME ZONE,
    is_logic_deleted BOOLEAN                  DEFAULT FALSE,
    version          INTEGER                  DEFAULT 1,
    user_id          UUID               NOT NULL REFERENCES users (id),
    card_number      VARCHAR(20) UNIQUE NOT NULL,
    cardholder_name  VARCHAR(100)       NOT NULL,
    expiration_date  DATE               NOT NULL,
    cvv              VARCHAR(4)         NOT NULL,
    card_type        VARCHAR(20),
    billing_address  TEXT,
    is_default       BOOLEAN                  DEFAULT FALSE,
    last_four_digits VARCHAR(4)
);

-- 创建Feed表
CREATE TABLE feeds
(
    id               UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    created_at       TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at       TIMESTAMP WITH TIME ZONE,
    is_logic_deleted BOOLEAN                  DEFAULT FALSE,
    version          INTEGER                  DEFAULT 1,
    user_id          UUID NOT NULL REFERENCES users (id),
    content          TEXT NOT NULL,
    media_url        VARCHAR(255),
    media_type       VARCHAR(20),
    like_count       INTEGER                  DEFAULT 0,
    comment_count    INTEGER                  DEFAULT 0,
    share_count      INTEGER                  DEFAULT 0,
    is_public        BOOLEAN                  DEFAULT TRUE,
    location         VARCHAR(100),
    tags             TEXT[],
    topics           TEXT[]
);

-- 创建Comment表
CREATE TABLE comments
(
    id                UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    created_at        TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at        TIMESTAMP WITH TIME ZONE,
    is_logic_deleted  BOOLEAN                  DEFAULT FALSE,
    version           INTEGER                  DEFAULT 1,
    feed_id           UUID NOT NULL REFERENCES feeds (id),
    user_id           UUID NOT NULL REFERENCES users (id),
    content           TEXT NOT NULL,
    like_count        INTEGER                  DEFAULT 0,
    parent_comment_id UUID REFERENCES comments (id),
    is_edited         BOOLEAN                  DEFAULT FALSE
);

-- 创建索引
CREATE INDEX idx_username ON users (username);
CREATE INDEX idx_users_deleted_at ON users (deleted_at);
CREATE INDEX idx_cards_deleted_at ON cards (deleted_at);
CREATE INDEX idx_feeds_deleted_at ON feeds (deleted_at);
CREATE INDEX idx_comments_deleted_at ON comments (deleted_at);
CREATE INDEX idx_feeds_user_id ON feeds (user_id);
CREATE INDEX idx_comments_feed_id ON comments (feed_id);
CREATE INDEX idx_comments_user_id ON comments (user_id);
CREATE INDEX idx_cards_user_id ON cards (user_id);

-- 为所有表创建触发器
CREATE TRIGGER update_users_timestamp
    BEFORE INSERT OR UPDATE
    ON users
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp_columns();

CREATE TRIGGER update_cards_timestamp
    BEFORE INSERT OR UPDATE
    ON cards
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp_columns();

CREATE TRIGGER update_feeds_timestamp
    BEFORE INSERT OR UPDATE
    ON feeds
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp_columns();

CREATE TRIGGER update_comments_timestamp
    BEFORE INSERT OR UPDATE
    ON comments
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp_columns();

-- 授予权限
GRANT ALL PRIVILEGES ON DATABASE meow_user TO postgres;
GRANT ALL PRIVILEGES ON DATABASE meow_user TO test_user;

-- 授予 test_user 在 public schema 上的权限
GRANT ALL PRIVILEGES ON SCHEMA public TO test_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO test_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO test_user;

-- 设置默认权限，使得 test_user 对未来创建的表和序列也有权限
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO test_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO test_user;

-- 创建一个只读用户（可选）
CREATE USER read_only_user WITH PASSWORD 'read_only_password';
GRANT CONNECT ON DATABASE meow_user TO read_only_user;
GRANT USAGE ON SCHEMA public TO read_only_user;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO read_only_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO read_only_user;

-- 添加一些测试数据
INSERT INTO users (username, email, password, role, preferred_language)
VALUES ('test_user1', 'test@example.com', 'hashed_password_1', 'user', 'en'),
       ('test_user2', 'user2@example.com', 'hashed_password_2', 'user', 'zh');

INSERT INTO cards (user_id, card_number, cardholder_name, expiration_date, cvv, card_type, last_four_digits)
VALUES ((SELECT id FROM users WHERE username = 'test_user1'), '1234567890123456', 'Test User 1', '2025-12-31', '123',
        'Visa', '3456'),
       ((SELECT id FROM users WHERE username = 'test_user2'), '9876543210987654', 'Test User 2', '2024-10-31', '456',
        'Mastercard', '7654');

INSERT INTO feeds (user_id, content, is_public)
VALUES ((SELECT id FROM users WHERE username = 'test_user1'), 'This is a test feed from user 1', true),
       ((SELECT id FROM users WHERE username = 'test_user2'), 'Another test feed from user 2', true);

INSERT INTO comments (feed_id, user_id, content)
VALUES ((SELECT id FROM feeds WHERE user_id = (SELECT id FROM users WHERE username = 'test_user1') LIMIT 1),
        (SELECT id FROM users WHERE username = 'test_user2'),
        'Great post!');

-- 验证数据（可选）
SELECT *
FROM users;
SELECT *
FROM cards;
SELECT *
FROM feeds;
SELECT *
FROM comments;