package upload

type FileInterface interface {
	FileUpload(objectKey, localFile string) error
	FileExists(objectKey string) (bool, error)
}
