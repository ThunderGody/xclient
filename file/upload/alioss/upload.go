package alioss

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	FilePrefix = "xpic"
)

type AlyOssService struct {
	*oss.Client
	Bucket string
}

func NewOssService(config OssConfig) *AlyOssService {
	client, err := oss.New(config.Endpoint, config.AccessKey, config.SecretKey)
	if err != nil {
		return nil
	}

	return &AlyOssService{Client: client, Bucket: config.Bucket}
}

// FileUpload 上传文件
// localFile 本地文件
func (aos *AlyOssService) FileUpload(localFilePath string) (string, error) {
	bucket, err := aos.bucket()
	if err != nil {
		return "", err
	}
	objectKey := GenerateObjectKey(FilePrefix, localFilePath)
	return objectKey, bucket.PutObjectFromFile(objectKey, localFilePath)
}

func (aos *AlyOssService) FileExists(objectKey string) (bool, error) {
	bucket, err := aos.bucket()
	if err != nil {
		return false, err
	}

	exists, err := bucket.IsObjectExist(objectKey)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (aos *AlyOssService) bucket() (*oss.Bucket, error) {
	bucket, err := aos.Client.Bucket(aos.Bucket)
	if err != nil {
		return nil, err
	}
	if bucket == nil {
		return nil, errors.New("bucket error")
	}

	return bucket, nil
}

// 删除本地文件
func (aos *AlyOssService) deleteLocalFile(filePath string) error {
	return os.Remove(filePath)
}

// GenerateObjectKey 生成唯一ID
func GenerateObjectKey(prefix, localFile string) string {
	objectKey := strconv.FormatInt(time.Now().UnixNano(), 10)
	ext := filepath.Ext(localFile)
	objectKey = prefix + RandomStr() + objectKey + ext

	return objectKey
}

// RandomStr 随机生成字符串
func RandomStr() string {
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Intn(899999)
	randInt += 100000
	return strconv.FormatInt(int64(randInt), 10)
}
