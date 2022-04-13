package hourly_upload

import (
    "log"
    "fmt"
    "mime/multipart"
    "bytes"
    "net/http"
    "io/ioutil"
    "encoding/base64"
)

func toBase64(file []byte) string {
    return base64.StdEncoding.EncodeToString(file)
}

func postRequestCall(image_base64encoding string, access_token string) {
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

    client := &http.Client {
    }

    req, err := http.NewRequest(method, url, payload)

    if err != nil {
        fmt.Println(err)
        return
    }

    log.Println(access_token)
    req.Header.Add("Authorization", "Bearer" + " " + access_token)

    req.Header.Set("Content-Type", form.FormDataContentType())
    res, err := client.Do(req)

    if err != nil {
        fmt.Println(err)
        return
    }

    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(string((body)))
}

func HourlyUpload(img_list []string, access_token string) {

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
    postRequestCall(image_base64encoding, access_token)

}
