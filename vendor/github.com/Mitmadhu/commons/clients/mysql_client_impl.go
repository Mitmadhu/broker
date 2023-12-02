package clients

import (
	"encoding/json"
	"net/http"

	"github.com/Mitmadhu/commons/handle_http"
	mysqlDto "github.com/Mitmadhu/mysqlDB/dto"
)

type MysqlClientImpl struct {
}

func (m *MysqlClientImpl) Login(url string, method string, req mysqlDto.ValidateUserRequest) (*mysqlDto.ValidateUserResponse, error) {

	byteResp, err := handle_http.Call(url+"/user-exists", http.MethodPost, mysqlDto.ValidateUserRequest{
		BaseRequest: mysqlDto.BaseRequest{MsgID: req.MsgID},
		Username:    req.Username,
		Password:    req.Password,
	})
	if err != nil {
		return nil, err
	}
	genericResp := &mysqlDto.ValidateUserResponse{}
	err = json.Unmarshal(byteResp, genericResp)
	return genericResp, err
}
