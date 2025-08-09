package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

func downloadFromUrl(url string, saveto string) error {
	client := grab.NewClient()
	req, err := grab.NewRequest(saveto, url)

	if err != nil {
		return err
	}

	fmt.Println("Downloading", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %s\n", resp.HTTPResponse.Status)

	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf(
				"  transferred %d / %d bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size(),
				100*resp.Progress(),
			)

		case <-resp.Done:
			fmt.Printf(
				"  transferred %d / %d bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size(),
				100*resp.Progress(),
			)
			break Loop
		}
	}

	if err := resp.Err(); err != nil {
		return err
	}

	fmt.Println("Download saved to", resp.Filename)
	return nil
}

func downloadAndGetContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
