package transfer

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const upload = "https://transfer.sh/"

func (b *transfer) DoUpload(name string, size int64, file io.Reader) error {

	body, err := b.newUpload(uploadConfig{
		fileSize:   size,
		fileName:   name,
		fileReader: file,
		debug:      b.Config.DebugMode,
	})
	if err != nil {
		return fmt.Errorf("upload returns error: %s", err)
	}

	b.resp = fmt.Sprintln("Download Link: " + string(body))

	return nil
}

func (b *transfer) PostUpload(string, int64) error {
	fmt.Printf("%s", b.resp)
	return nil
}

func (b transfer) newUpload(config uploadConfig) ([]byte, error) {
	if config.debug {
		log.Printf("start upload")
	}
	client := http.Client{}

	req, err := http.NewRequest("PUT", upload+config.fileName, config.fileReader)
	if err != nil {
		return nil, err
	}
	if config.debug {
		log.Printf("header: %v", req.Header)
	}
	resp, err := client.Do(req)
	if err != nil {
		if config.debug {
			log.Printf("do requests returns error: %v", err)
		}
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if config.debug {
			log.Printf("read response returns: %v", err)
		}
		return nil, err
	}
	_ = resp.Body.Close()
	if config.debug {
		log.Printf("returns: %v", string(body))
	}

	return body, nil
}