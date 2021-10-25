package alioss

import "testing"

func TestAlyOssService_FileUpload(t *testing.T) {
	config := OssConfig{
		Bucket:    "yujs",
		Endpoint:  "",
		AccessKey: "",
		SecretKey: "",
	}
	ossService := NewOssService(config)
	filePath := "../../test.png"

	ossFilePath, err := ossService.FileUpload(filePath)
	if err != nil {
		t.Fatal("upload file error, err", err.Error())
	}
	t.Log("upload file success, ossFilePath:", ossFilePath)
}
