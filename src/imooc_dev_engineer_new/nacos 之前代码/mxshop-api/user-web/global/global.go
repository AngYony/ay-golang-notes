package global

import (
	"mxshop-api/user-web/config"
	"mxshop-api/user-web/proto"

	ut "github.com/go-playground/universal-translator"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	Trans        ut.Translator

	UserSrvClient proto.UserClient
)
