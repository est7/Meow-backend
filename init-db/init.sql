/*-- 注意：postgres 用户通常已经存在，是默认的超级用户
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

-- 创建用户表
CREATE TABLE users
(
    id            UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    username      VARCHAR(255) NOT NULL UNIQUE,
    email         VARCHAR(255) UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 创建索引
CREATE INDEX idx_username ON users (username);
CREATE INDEX idx_email ON users (email);

-- 创建触发器函数来自动更新 updated_at
CREATE OR REPLACE FUNCTION update_modified_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 为 users 表创建触发器
CREATE TRIGGER update_users_modtime
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

-- 授予权限
GRANT ALL PRIVILEGES ON DATABASE meow_user TO postgres;
GRANT ALL PRIVILEGES ON DATABASE meow_user TO test_user;

-- 切换到 meow_user 数据库（如果还没有切换的话）
\c meow_user

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

-- 添加一些测试数据（可选）
INSERT INTO users (username, email, password_hash)
VALUES ('test_user1', 'user1@example.com', 'hashed_password_1'),
       ('test_user2', 'user2@example.com', 'hashed_password_2');

-- 验证数据（可选）
SELECT *
FROM users;*/