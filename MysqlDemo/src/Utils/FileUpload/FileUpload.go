package FileUpload

import (
	"log"
	"net/http"
	"bytes"
	"fmt"
	"os"
	"mime/multipart"
	"path/filepath"
	"io"
	"time"
)

func GoroutineTest() {
	fmt.Println("start!!!!!")
	start := time.Now()
	defer func() {
		cost := time.Since(start)
		fmt.Println("sum cost=", cost)
	}()
	channels := make(chan bool)
	for i := 0; i < 40; i++ {
		go updateFileSignle(i, channels)
	}

	for i := 0; i < 40; i++ {
		<-channels
	}
}

func updateFileSignle(i int, c chan bool) {
	fmt.Println(i)
	extraParams := map[string]string{
		"ModuleType": "SaasOAAdmin",
	}
	request, err := newMultipartRequest("http://59.110.225.214:8093/api/V1/OSS/FileSingle", extraParams, "file", "D:/Test/testFU.docx")
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(body)
	}
	c <- true
}

// Creates a new file upload http request with optional extra params
func newMultipartRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Token", "B62ED9F0B636411D1849E9A4DB144EE2")
	request.Header.Add("EnterpriseId", "3")
	request.Header.Add("AuthCode", "777")
	request.Header.Add("TerminalType", "1")

	return request, err
}
