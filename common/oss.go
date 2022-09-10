package common

import (
	"net/http"
	"net/url"
	"sheep/demo/config"

	"github.com/tencentyun/cos-go-sdk-v5"
)

var OSSClient *cos.Client

func InitCosClient() (err error) {
	u, _ := url.Parse("https://insurence-1304011999.cos.ap-shanghai.myqcloud.com")
	su, _ := url.Parse("https://cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{
		BucketURL:  u,
		ServiceURL: su,
	}
	conf := config.GetOSSConfig()
	OSSClient = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  conf.SecretID,
			SecretKey: conf.SecretKey,
		},
	})
	return nil
}
