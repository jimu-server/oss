package oss

import (
	"github.com/jimu-server/config"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

var Tencent *cos.Client

func tencent() {
	bucketurl := config.Evn.App.Tencent.BucketURL
	u, _ := url.Parse(bucketurl)
	serviceurl := config.Evn.App.Tencent.ServiceURL
	su, _ := url.Parse(serviceurl)
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	Tencent = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.Evn.App.Tencent.SecretID,
			SecretKey: config.Evn.App.Tencent.SecretKey,
		},
	})
}
