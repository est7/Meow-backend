package service

import (
	"Meow-backend/pkg/errcode"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// VerifyEmailCode 验证邮箱验证码
func (s *UserServiceImpl) VerifyEmailCode(ctx context.Context, email, code string) error {
	// 从 Redis 获取存储的验证码
	storedCode, err := s.GetRedis().Get(ctx, fmt.Sprintf("email_code:%s", email)).Result()
	if err != nil {
		if err == redis.Nil {
			return errcode.ErrVerificationCodeExpired
		}
		return err
	}

	// 验证码比对
	if storedCode != code {
		return errcode.ErrVerificationCodeInvalid
	}

	// 验证成功后删除验证码
	s.redis.Del(ctx, fmt.Sprintf("email_code:%s", email))

	return nil
}

// VerifyPhoneCode 验证手机验证码
func (s *UserServiceImpl) VerifyPhoneCode(ctx context.Context, phone, code string) error {
	// 从 Redis 获取存储的验证码
	storedCode, err := s.redis.Get(ctx, fmt.Sprintf("phone_code:%s", phone)).Result()
	if err != nil {
		if err == redis.Nil {
			return errcode.ErrVerificationCodeExpired
		}
		return err
	}

	// 验证码比对
	if storedCode != code {
		return errcode.ErrVerificationCodeInvalid
	}

	// 验证成功后删除验证码
	s.redis.Del(ctx, fmt.Sprintf("phone_code:%s", phone))

	return nil
}
