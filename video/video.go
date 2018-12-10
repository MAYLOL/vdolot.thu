// h 20181210
//
// Video Thumb Service

package video

import (
	"encoding/json"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/opennota/screengen"
)

func Thumb(w io.Writer, r *ThumbRequest) (statusCode int) {
	statusCode = http.StatusInternalServerError
	for {
		//log.Printf("req: r='%s', t='%s', w='%s', h='%s'\n", r.R, r.T, r.W, r.H)
		if _, err := url.Parse(r.R); err != nil {
			statusCode = http.StatusBadRequest
			break
		}
		// Parse resource
		gen, err := screengen.NewGenerator(r.R)
		if err != nil {
			break
		}
		defer gen.Close()
		// Parse timetick
		t, _ := time.ParseDuration(r.T)
		if t == 0 {
			t = time.Second
		}
		// Parse width & height
		_w, _ := strconv.Atoi(r.W)
		_h, _ := strconv.Atoi(r.H)
		if _w == 0 && _h != 0 {
			_w = gen.Width() * _h / gen.Height()
		}
		if _h == 0 && _w != 0 {
			_h = gen.Height() * _w / gen.Width()
		}
		// Take screenshot
		var img image.Image
		if _w != 0 || _h != 0 {
			img, err = gen.ImageWxH(int64(t/time.Millisecond), _w, _h)
		} else {
			img, err = gen.Image(int64(t / time.Millisecond))
		}
		if err != nil {
			break
		}
		// Encode JPEG 4:2:0
		err = jpeg.Encode(w, img, &jpeg.Options{Quality: 80})
		if err != nil {
			break
		}
		statusCode = http.StatusOK
		//
		// Finally
		if true {
			break
		}
	}
	return statusCode
}

func Meta(w io.Writer, r *MetaRequest) (statusCode int) {
	statusCode = http.StatusInternalServerError
	for {
		//log.Printf("req: r='%s'\n", r.R)
		if _, err := url.Parse(r.R); err != nil {
			statusCode = http.StatusBadRequest
			break
		}
		// Parse resource
		gen, err := screengen.NewGenerator(r.R)
		if err != nil {
			break
		}
		defer gen.Close()
		// Parse meta-data
		b, err := json.Marshal(struct {
			*screengen.Generator
			VideoCodecLongName string `json:",omitempty"`
			AudioCodecLongName string `json:",omitempty"`
			Width              int    `json:"width"`
			Height             int    `json:"height"`
		}{
			Generator: gen,
			Width:     gen.Width(),
			Height:    gen.Height(),
		})
		if err != nil {
			break
		}
		w.Write(b)
		statusCode = http.StatusOK
		//
		// Finally
		if true {
			break
		}
	}
	return statusCode
}
