package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"mxshop_srvs/user_srv/global"
	"mxshop_srvs/user_srv/model"
	"mxshop_srvs/user_srv/proto"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/anaskhan96/go-password-encoder"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"gorm.io/gorm"
)

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func ModelToResponse(user model.User) *proto.UserInfoResponse {

	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		PassWord: user.PassWord,
		Mobile:   user.Mobile,
		NickName: user.NickName,
		// BirthDay: user.Birthday,
		Gender: user.Gender,
		Role:   uint32(user.Role),
	}
	// 注意：在grpc的message中，字段都有默认值，因此不能随便给赋值为nil，否则在序列化的时候容易报错
	if user.Birthday != nil {
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}
	return &userInfoRsp

}

type UserServer struct{}

func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	// 获取用户列表
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	fmt.Println("用户列表")
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)

	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)
	for _, user := range users {
		userinfoRsp := ModelToResponse(user)
		rsp.Data = append(rsp.Data, userinfoRsp)
	}

	return rsp, nil
}

// GetUserMobile 通过手机号码查询用户
func (s *UserServer) GetUserMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	if result.Error != nil {
		return nil, result.Error
	}
	return ModelToResponse(user), nil
}

// GetUserById 通过Id查询用户
func (s *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	if result.Error != nil {
		return nil, result.Error
	}
	return ModelToResponse(user), nil
}

// CreateUser 创建用户
func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	// 查询用户是否存在
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已经存在")
	}

	user.Mobile = req.Mobile
	user.NickName = req.NickName

	// 密码加密
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(req.PassWord, options)

	user.PassWord = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	return ModelToResponse(user), nil

}

// UpdateUser 个人中心更新用户
func (s *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*empty.Empty, error) {
	// 查询用户
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	// 将 uint64 转换为 time 类型
	birthday := time.Unix(int64(req.BirthDay), 0)
	user.NickName = req.NickName
	user.Birthday = &birthday
	user.Gender = req.Gender

	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil

}

// CheckPassWord 校验密码
func (s *UserServer) CheckPassWord(ctx context.Context, req *proto.PassWordCheckInfo) (*proto.CheckResponse, error) {
	// 密码加密
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	pwdinfo := strings.Split(req.EncryptedPassWord, "$")
	check := password.Verify(req.PassWord, pwdinfo[2], pwdinfo[3], options)
	return &proto.CheckResponse{Success: check}, nil

}
