package clients

import (
	mysqlDto "github.com/Mitmadhu/mysqlDB/dto"
)

type MysqlClient interface {
	Login(url string, method string, reqDto mysqlDto.ValidateUserRequest) (*mysqlDto.ValidateUserResponse, error)
}
