package alioss

import "testing"

func TestAlyOssService_FileUpload(t *testing.T) {
	config := OssConfig{
		Bucket:    "yujs",
		Endpoint:  "oss-cn-beijing.aliyuncs.com",
		AccessKey: "LTAI5t7ZysrHQ55pc2KzcKJu",
		SecretKey: "AMHuQgEANOUTIp2BPiF9fkC5M8HpHt",
	}
	ossService := NewOssService(config)
	filePath := "../../test.png"

	ossFilePath, err := ossService.FileUpload(filePath)
	if err != nil {
		t.Fatal("upload file error, err", err.Error())
	}
	t.Log("upload file success, ossFilePath:", ossFilePath)
}
