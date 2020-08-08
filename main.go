package main

import (
	"time"

	"github.com/iMro0t/go-pvr/v18/pvr"
)

// Do not remove this
func init() { pvr.Call = main }

func main() {
	pvr.AddChannel(pvr.Channel{
		ID:        1,
		IsRadio:   false,
		Number:    1,
		SubNumber: 1,
		Name:      "Big Buck Bunny",
		IconPath:  "https://peach.blender.org/wp-content/uploads/poster_bunny_small.jpg",
		IsHidden:  false,
		Live: pvr.Stream{
			URL: "http://distribution.bbb3d.renderfarming.net/video/mp4/bbb_sunflower_1080p_30fps_normal.mp4",
		},
	})

	pvr.AddChannelGroup(pvr.ChannelGroup{
		IsRadio:  false,
		Name:     "Kids",
		Position: 1,
		Members:  []int{1},
	})

	pvr.AddEPG(pvr.EPG{
		BroadcastID:      100,
		ChannelID:        1,
		Title:            "Big Buck Bunny",
		StartTime:        time.Now().Unix(),
		EndTime:          time.Now().Add(time.Hour).Unix(),
		PlotOutline:      "A large and lovable rabbit deals with three tiny bullies, led by a flying squirrel, who are determined to squelch his happiness.",
		Plot:             "A recently awoken enormous and utterly adorable fluffy rabbit is heartlessly harassed by a flying squirrel's gang of rodents who are determined to squash his happiness.",
		OriginalTitle:    "Big Buck Bunny",
		Cast:             []string{"Bunny"},
		Director:         []string{"Sacha Goedegebure"},
		Writer:           []string{"Sacha Goedegebure"},
		Year:             2008,
		IMDBNumber:       "tt1254207",
		IconPath:         "https://peach.blender.org/wp-content/uploads/poster_bunny_small.jpg",
		GenreType:        pvr.EPGGenre.Custom,
		GenreSubType:     pvr.EPGGenre.Unknown,
		GenreDescription: []string{"Fun"},
		FirstAired:       time.Date(2008, 04, 10, 0, 0, 0, 0, time.Local).Unix(),
		StarRating:       3,
	})
}
