package image

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	size                = "200px,200px" // width,height in pixels (e.g. 1024px,768px or 3,3)
	preserveAspectRatio = false         // preserve aspect ratio
)

// Print image in terminal from url
func DisplayNetPicture(filename string) {
	res, err := http.Get(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if err := display(res.Body); err != nil {
		log.Fatal(err)
	}
}

func display(r io.Reader) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	width, height := widthAndHeight(size)

	fmt.Print("\033]1337;")
	fmt.Printf("File=inline=1")
	if width != "" || height != "" {
		if width != "" {
			fmt.Printf(";width=%s", width)
		}
		if height != "" {
			fmt.Printf(";height=%s", height)
		}
	}
	if preserveAspectRatio {
		fmt.Print("preserveAspectRatio=1")
	}
	fmt.Print(":")
	fmt.Printf("%s", base64.StdEncoding.EncodeToString(data))
	fmt.Print("\a\n")

	return nil
}

func widthAndHeight(size string) (w, h string) {
	if size != "" {
		sp := strings.SplitN(size, ",", -1)
		if len(sp) == 2 {
			w = sp[0]
			h = sp[1]
		}
	}
	return
}
