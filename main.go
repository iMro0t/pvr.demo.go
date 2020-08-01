package main

import (
	"github.com/iMro0t/pvr.demo.go/pvr"
)

// Do not remove this
func init() { pvr.Call = main }

func main() {
	pvr.AddChannel(pvr.Channel{
		ID:        1,
		IsRadio:   false,
		Number:    1,
		SubNumber: 1,
		Name:      "Demo TV Channel 1",
		IconPath:  "",
		IsHidden:  false,
		Live: pvr.Stream{
			URL: "http://distribution.bbb3d.renderfarming.net/video/mp4/bbb_sunflower_1080p_30fps_normal.mp4",
		},
	})
}
