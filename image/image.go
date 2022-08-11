package image

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	width               = flag.String("width", "200px", "width (e.g. 100px, 10%, or auto)")
	height              = flag.String("height", "200px", "height (e.g. 100px, 10%, or auto)")
	size                = flag.String("size", "", "width,height in pixels (e.g. 1024px,768px or 3,3)")
	preserveAspectRatio = flag.Bool("p", false, "preserve aspect ratio")
)

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

func displayLocalPicture(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := display(f); err != nil {
		log.Fatal(err)
	}
}

func display(r io.Reader) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	width, height := widthAndHeight()

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
	if *preserveAspectRatio {
		fmt.Print("preserveAspectRatio=1")
	}
	fmt.Print(":")
	fmt.Printf("%s", base64.StdEncoding.EncodeToString(data))
	fmt.Print("\a\n")

	return nil
}

func widthAndHeight() (w, h string) {
	if *width != "" {
		w = *width
	}
	if *height != "" {
		h = *height
	}
	if *size != "" {
		sp := strings.SplitN(*size, ",", -1)
		if len(sp) == 2 {
			w = sp[0]
			h = sp[1]
		}
	}
	return
}
