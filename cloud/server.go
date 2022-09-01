package cloud

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func BestServer() string {
	resp, _ := http.Get("https://api.gofile.io/getServer")

	body, _ := ioutil.ReadAll(resp.Body)

	var serverResponse ServerResponse
	_ = json.Unmarshal(body, &serverResponse)
	return serverResponse.Data.Server
}
