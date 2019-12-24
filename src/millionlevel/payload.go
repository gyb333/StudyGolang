package millionlevel

type PayloadCollection struct {
	WindowsVersion  string    `json:"version"`
	Token           string    `json:"token"`
	Payloads        []Payload `json:"data"`
}

//type Payload struct {
//	// [redacted]
//}

func (p *Payload) UploadToS3() error {
	//// the storageFolder method ensures that there are no name collision in
	//// case we get same timestamp in the key name
	//storage_path := fmt.Sprintf("%v/%v", p.storageFolder, time.Now().UnixNano())
	//
	//bucket := S3Bucket
	//
	//b := new(bytes.Buffer)
	//encodeErr := json.NewEncoder(b).Encode(payload)
	//if encodeErr != nil {
	//	return encodeErr
	//}
	//
	//// Everything we post to the S3 bucket should be marked 'private'
	//var acl = s3.Private
	//var contentType = "application/octet-stream"
	//
	//return bucket.PutReader(storage_path, b, int64(b.Len()), contentType, acl, s3.Options{})
	return nil
}
