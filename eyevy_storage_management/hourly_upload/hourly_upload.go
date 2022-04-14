package hourly_upload

import (
    "log"
    "fmt"
    "mime/multipart"
    "bytes"
    "strings"
    "net/http"
    "io/ioutil"
    "encoding/base64"
    "encoding/json"
)

func postToWebhook (img_url string, url string) {
    // Find out what the post request looks like.
    // The docs have a curl example that you can reverse-engineer.

	client := &http.Client{}
	var data = strings.NewReader(`{"content": "", 
                                   "embeds": [
                                        {"title": "Hourly Upload"},
                                        {"image": {"url": "`+img_url+`"}}
                                        ]}`)
 

	req, err := http.NewRequest("POST", url, data)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", bodyText)
}

func toBase64(file []byte) string {
    return base64.StdEncoding.EncodeToString(file)
}

func postRequestCall(image_base64encoding string, 
                     access_token string) (post_result string) {

    // Based off Imgur API Docs, with the addition of manually converting
    // the image bytes into base64 string.

    url := "https://api.imgur.com/3/image"
    method := "POST"

    payload := &bytes.Buffer{}
    form := multipart.NewWriter(payload) // We're making a form here.
    _ = form.WriteField("image", image_base64encoding)
    err := form.Close()

    if err != nil {
        fmt.Println(err)
        return
    }

    client := &http.Client {}

    req, err := http.NewRequest(method, url, payload)
    if err != nil {
        fmt.Println(err)
        return
    }

    req.Header.Add("Authorization", "Bearer" + " " + access_token)
    req.Header.Set("Content-Type", form.FormDataContentType())
    res, err := client.Do(req)

    if err != nil {
        fmt.Println(err)
        return
    }

    defer res.Body.Close()

    // The POST request's result body.
    body, err := ioutil.ReadAll(res.Body) 
    if err != nil {
        fmt.Println(err)
    }

    return string(body)
}

func HourlyUpload(img_list []string, 
                  access_token string, 
                  webhook_url string) {

    log.Println("In HourlyUpload function")

    most_recent := img_list[len(img_list)-1]
    log.Println("Most recent image is:" + most_recent)

    // Opening up file to get bytes..
    img_bytes, read_err := ioutil.ReadFile("../captures/" + most_recent)

    if read_err != nil {
        log.Fatal(read_err)
    }

    // Grabbing base64 string encoding of bytes.
    var image_base64encoding string = toBase64(img_bytes)

    // Response of POST request to Imgur API
    imgur_upload_res := postRequestCall(image_base64encoding, access_token)

    // Converting the POST result from string to indexable JSON
    var imgur_upload_res_json map[string]interface{}
    json.Unmarshal([]byte(imgur_upload_res), &imgur_upload_res_json)

    image_upload_data := imgur_upload_res_json["data"].(map[string]interface{})
    log.Println(image_upload_data["link"])
    log.Printf("%T\n",image_upload_data["link"])

    log.Println(webhook_url)
    postToWebhook(image_upload_data["link"].(string), webhook_url)
}
