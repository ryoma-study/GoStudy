package main

import (
	"context"
	proto "v1/api/user"
	"v1/internal/user/biz"
)

func (u *UserService) GetUser(ctx context.Context, r *proto.UserRequest) (*proto.UserResponse, error) {
	return biz.GetUser(r)
}
