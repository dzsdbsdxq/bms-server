package upload

import (
	"bms-server/global"
	"bms-server/utils"
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

type SzTvLocalResponse struct {
	Data []struct {
		Aid             int    `json:"aid"`
		Filename        string `json:"filename"`
		Url             string `json:"url"`
		CompressUrl     string `json:"compressUrl"`
		OriginUrl       string `json:"originUrl"`
		HttpUrl         string `json:"httpUrl"`
		CompressHttpUrl string `json:"compressHttpUrl"`
		OriginHttpUrl   string `json:"originHttpUrl"`
		Width           int    `json:"width"`
		Height          int    `json:"height"`
	} `json:"data"`
	Message string `json:"message"`
	State   int    `json:"state"`
	Success bool   `json:"success"`
}

type SzTvLocal struct{}

func (st *SzTvLocal) UploadFile(file *multipart.FileHeader) (string, string, error) {
	date, nonce, sign, err := st.Authorized("POST", "/api/revelation/uploadImg")
	if err != nil {
		return "", "", err
	}
	f, openError := file.Open()
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() failed", zap.Any("err", openError.Error()))

		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close()

	requestBody := &bytes.Buffer{}

	multipartWriter := multipart.NewWriter(requestBody)

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	_ = multipartWriter.WriteField("name", fileKey)
	formFile, err := multipartWriter.CreateFormFile("file", file.Filename)
	if err != nil {
		return "", "", err
	}

	_, err = io.Copy(formFile, f)

	err = multipartWriter.Close()
	if err != nil {
		return "", "", err
	}

	req, err := http.NewRequest("POST", "https://api.scms.sztv.com.cn/api/revelation/uploadImg?TenantID=ysz", requestBody)
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Set("X-Date", date)
	req.Header.Set("Nonce", nonce)
	req.Header.Set("Authorization", sign)

	// 发送GET请求并获取响应
	client := &http.Client{}
	postResponse, err := client.Do(req)
	if err != nil {
		fmt.Println("发送POST请求失败:", err)
		return "", "", err
	}
	defer postResponse.Body.Close()

	// 读取POST响应的内容
	postResponseBody, err := ioutil.ReadAll(postResponse.Body)
	if err != nil {
		fmt.Println("读取POST响应失败:", err)
		return "", "", err
	}
	var resData *SzTvLocalResponse

	err = jsoniter.Unmarshal(postResponseBody, &resData)
	if err != nil {
		return "", "", err
	}
	if resData.State == 200 {
		if len(resData.Data) > 0 {
			return resData.Data[0].OriginHttpUrl, fileKey, nil
		}
		return "", "", nil
	}

	return "", "", errors.New(resData.Message)
}

func (st *SzTvLocal) DeleteFile(key string) error {
	return nil
}

func (st *SzTvLocal) Authorized(method string, path string) (gmtDate string, nonce string, sign string, err error) {
	nonce = utils.GetRandomId(32)
	// 将UTC时间格式化为易读的字符串
	utcString := time.Now().UTC().Format(time.RFC1123)
	signingString := fmt.Sprintf("x-date: %s \n@request-target: %s %s\nhost: apix.scms.sztv.com.cn\nnonce: %s", utcString, strings.ToLower(method), path, nonce)
	mac := hmac.New(sha512.New, []byte("xUJ7Gls45St0CTnatnwZwsH4UyYj0rpX"))
	_, err = mac.Write([]byte(signingString))
	if err != nil {
		return "", "", "", err
	}
	enc := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	auth := fmt.Sprintf("hmac username=\"onesz\", algorithm=\"hmac-sha512\", headers=\"x-date @request-target host nonce\", signature=\"%s\"", enc)

	return utcString, nonce, auth, nil
}
