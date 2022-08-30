package cloudinary

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func Upload(file interface{}) (string, error) {
	cld, _ := cloudinary.NewFromParams("", "", "")
	ctx := context.Background()

	name := time.Now().Unix()
	temp := fmt.Sprint(name)

	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: temp})
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return resp.SecureURL, nil
}
