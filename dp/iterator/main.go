package iterator

import (
	"fmt"

	"github.com/OdaDaisuke/go-algorithm/dp/iterator/internal"
)

func Run() {
	musics := internal.NewMusics()
	musics.Append(&internal.Music{
		Id:     1,
		Title:  "MusicA",
		Artist: "ArtistA",
	})
	musics.Append(&internal.Music{
		Id:     2,
		Title:  "MusicB",
		Artist: "ArtistB",
	})

	for musics.HasNext() {
		m := musics.Scan()
		fmt.Println(m)
	}
}
