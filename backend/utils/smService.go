/**
 * @date: 2022/2/17
 * @desc: ...
 */

package utils

import (
	"DataCompare/global"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//
// SmService
// @Description: SmService
//
type SmService struct {
	SendSmBaseUrl   string
	AppSecret       string
	AppKey          string
	SMTemplateCode  int
	CodeLen         int
	VerifySmBaseUrl string
}

//
// NewSmService
// @Description: 构建Sm服务
// @return SmService:
//
func NewSmService() *SmService {
	smService := &SmService{
		global.ServerConfig.SmInfo.SendSmBaseUrl,
		global.ServerConfig.SmInfo.AppSecret,
		global.ServerConfig.SmInfo.AppKey,
		global.ServerConfig.SmInfo.SMTemplateCode,
		global.ServerConfig.SmInfo.CodeLen,
		global.ServerConfig.SmInfo.VerifySmBaseUrl,
	}
	return smService
}

//
// SendSmCode
// @Description: 发送验证码
// @receiver s
//
func (s *SmService) SendSmCode(mobile string) (resObj string, err error) {
	rand.Seed(time.Now().UnixNano())
	nonce := strconv.FormatInt(int64(rand.Intn(10)), 10)
	curTime := strconv.FormatInt(time.Now().Unix(), 10)
	checkSum, err := buildCheckSum(curTime, nonce)
	if err != nil {
		return "", err
	}
	client := http.Client{}
	body := strings.NewReader(fmt.Sprintf(`codeLen=%d&templateid=%d&mobile=%s`, s.CodeLen, s.SMTemplateCode, mobile))
	request, err := http.NewRequest("POST", s.SendSmBaseUrl, body)
	// 增加 head 选项
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	request.Header.Add("AppKey", s.AppKey)
	request.Header.Add("CurTime", curTime)
	request.Header.Add("Nonce", nonce)
	request.Header.Add("CheckSum", checkSum)
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	resMsg, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(resMsg))
	return string(resMsg), err
}

//
// VerifySmCode
// @Description: 校验验证码
//
func (s *SmService) VerifySmCode(mobile, code string) (res bool, err error) {
	// todo 生产环境删除此处return
	//return true, nil
	rand.Seed(time.Now().UnixNano())
	nonce := strconv.FormatInt(int64(rand.Intn(10)), 10)
	curTime := strconv.FormatInt(time.Now().Unix(), 10)
	checkSum, err := buildCheckSum(curTime, nonce)
	if err != nil {
		return false, err
	}
	client := http.Client{}
	body := strings.NewReader(fmt.Sprintf(`mobile=%s&code=%s`, mobile, code))
	request, err := http.NewRequest("POST", s.VerifySmBaseUrl, body)
	// 增加 head 选项
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	request.Header.Add("AppKey", s.AppKey)
	request.Header.Add("CurTime", curTime)
	request.Header.Add("Nonce", nonce)
	request.Header.Add("CheckSum", checkSum)
	response, err := client.Do(request)
	if err != nil {
		global.Logger.Error("验证码服务异常:", zap.String("error", err.Error()))
		return false, err
	}
	defer response.Body.Close()
	resMsg, err := ioutil.ReadAll(response.Body)
	if err != nil {
		global.Logger.Error("验证码服务异常:", zap.String("error", err.Error()))
		return false, err
	}
	type MsgCode struct {
		Code int
		Msg  string
		Obj  string
	}
	msgCodeObj := &MsgCode{}
	_ = json.Unmarshal(resMsg, msgCodeObj)
	if msgCodeObj.Code != 200 {
		return false, nil
	}
	return true, nil
}

//
// BuildCheckSum
// @Description: 构建 BuildCheckSum
//
func buildCheckSum(curTime, nonce string) (checkSum string, err error) {
	resStr := global.ServerConfig.SmInfo.AppSecret + nonce + curTime
	t := sha1.New()
	_, err = io.WriteString(t, resStr)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", t.Sum(nil)), nil
}
