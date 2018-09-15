package service

import (
	"bytes"
	"github.com/dchest/captcha"
	"insur-box/src/models"
	"insur-box/src/config"
	"insur-box/src/utils"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"insur-box/src/db"
	"time"
	"github.com/gomodule/redigo/redis"
)

const ValidTime = 10 * time.Minute

// 获取图片验证码地址与验证id
func ImagesCode(style models.ImageStyle) (string, string) {
	var content bytes.Buffer
	// 指定默认数字个数
	if style.Length == 0 {
		style.Length = 6
	}
	// 指定oss前缀
	if style.Prefix == "" {
		style.Prefix = "captcha_"
	}

	id, code := utils.CreateCode(style.Length)
	db.Redis.Set(id, utils.BytesToString(code), ValidTime)
	captcha.NewImage(id, code, style.Width, style.Height).WriteTo(&content)
	name := style.Prefix + id + ".png"
	UploadToOss(name, content)
	return config.PicHost + name, id
}

// 验证验证码图片是否正确
func VerifyImageCode(input models.VerifyCode) bool {
	code, _ := redis.String(db.Redis.Get(input.Id))
	db.Redis.Delete(input.Id)
	return code == input.Code
}

// 上传文件到oss
func UploadToOss(name string, content bytes.Buffer) error {
	client, err := oss.New(config.Endpoint, config.AccessKeyId, config.AccessKeySecret)
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(config.Bucket)
	if err != nil {
		return err
	}
	return bucket.PutObject(name, bytes.NewReader(content.Bytes()))

}
