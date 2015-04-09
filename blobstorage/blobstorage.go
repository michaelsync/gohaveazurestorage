package blobstorage

import "gohaveazurestorage/gohaveazurestoragecommon"

type BlobStorage struct {
	http             *gohaveazurestoragecommon.HTTP
	baseUrl          string
	secondaryBaseUrl string
}

func New(http *gohaveazurestoragecommon.HTTP) *BlobStorage {
	return &BlobStorage{http: http}
}

func (blobStorage *BlobStorage) CreateContainer(containerName string) (httpStatusCode int) {
	_, httpStatusCode = blobStorage.http.Request("PUT", containerName, "?restype=container", nil, map[string]string{}, false, false, false)
	return httpStatusCode
}

func (blobStorage *BlobStorage) DeleteContainer(containerName string) (httpStatusCode int) {
	_, httpStatusCode = blobStorage.http.Request("DELETE", containerName, "?restype=container", nil, map[string]string{}, false, false, false)
	return httpStatusCode
}

func (blobStorage *BlobStorage) ListContainers() (body []byte, httpStatusCode int) {
	return blobStorage.http.Request("GET", "?comp=list", "", nil, map[string]string{}, true, false, false)
}

func (blobStorage *BlobStorage) GetContainerProperties(containerName string) (body []byte, httpStatusCode int) {
	return blobStorage.http.Request("GET", containerName, "?restype=container", nil, map[string]string{}, true, false, false)
}

func (blobStorage *BlobStorage) GetContainerMetadata(containerName string) (body []byte, httpStatusCode int) {
	return blobStorage.http.Request("GET", containerName+"?comp=metadata", "&restype=container", nil, map[string]string{}, true, false, false)
}

func (blobStorage *BlobStorage) ListBlobs(containerName string) (body []byte, httpStatusCode int) {
	return blobStorage.http.Request("GET", containerName+"?comp=list", "&restype=container", nil, map[string]string{}, true, false, false)
}
