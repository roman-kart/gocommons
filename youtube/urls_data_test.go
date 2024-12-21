package youtube_test

import (
	"fmt"

	"github.com/roman-kart/gocommons/random"
	"github.com/roman-kart/gocommons/youtube"
)

type testGetVideoIdData struct {
	Url     string
	VideoId string
	Err     error
}

type testGetThumbnailRequestUrl struct {
	VideoId string
	Quality youtube.ThumbnailQuality
	Result  string
}

func getVideIdDataStable() []testGetVideoIdData {
	return []testGetVideoIdData{
		{
			Url:     "https://www.youtube.com/watch?v=jNQXAC9IVRw",
			VideoId: "jNQXAC9IVRw",
		},
		{
			Url:     "https://youtu.be/jNQXAC9IVRw",
			VideoId: "jNQXAC9IVRw",
		},
		{
			Url:     "https://www.youtube.com/v/jNQXAC9IVRw",
			VideoId: "jNQXAC9IVRw",
		},

		{
			Url: "https://en.wikipedia.org/wiki/Me_at_the_zoo",
			Err: youtube.ErrHostIsNotYoutube,
		},
		{
			Url: "https://www.youtube.com/",
			Err: youtube.ErrWrongUrlFormat,
		},
		{
			Url: "https://www.youtube.com/x/jNQXAC9IVRw",
			Err: youtube.ErrWrongUrlFormat,
		},
		{
			Url: "https://youtu.be/",
			Err: youtube.ErrWrongUrlFormat,
		},
	}
}

var youtubeHosts = []string{"youtu.be", "www.youtube.com", "youtube.com"}

func getRandomYoutubeHost() string {
	return youtubeHosts[random.Int(0, len(youtubeHosts))]
}

func getTestVideoIdDataRandomValid(cnt int) []testGetVideoIdData {
	data := make([]testGetVideoIdData, 0, cnt)

	for range cnt {
		urlStr, videoId := getRandomUrlStrAndVideoId()

		data = append(data, testGetVideoIdData{
			Url:     urlStr,
			VideoId: videoId,
		})
	}

	return data
}

func getRandomUrlStrAndVideoId() (string, string) {
	var urlStr string

	videoId := random.StringFromBytes(16, random.ASCIILetters+"-")
	host := getRandomYoutubeHost()

	switch host {
	case "youtu.be":
		urlStr = fmt.Sprintf("https://%s/%s", host, videoId)
	case "www.youtube.com", "youtube.com":
		if random.Int(0, 1000)%2 == 0 {
			urlStr = fmt.Sprintf("https://%s/watch?v=%s", host, videoId)
		} else {
			urlStr = fmt.Sprintf("https://%s/v/%s", host, videoId)
		}
	}

	return urlStr, videoId
}

func getTestGetThumbnailRequestUrlStable() []testGetThumbnailRequestUrl {
	return []testGetThumbnailRequestUrl{
		{
			VideoId: "test",
			Quality: youtube.ThumbnailQualityHQ,
			Result:  "https://img.youtube.com/vi/test/hqdefault.jpg",
		},
		{
			VideoId: "test",
			Quality: youtube.ThumbnailQualityMQ,
			Result:  "https://img.youtube.com/vi/test/mqdefault.jpg",
		},
		{
			VideoId: "test",
			Quality: youtube.ThumbnailQualitySD,
			Result:  "https://img.youtube.com/vi/test/sddefault.jpg",
		},
		{
			VideoId: "test",
			Quality: youtube.ThumbnailQualityMaxRes,
			Result:  "https://img.youtube.com/vi/test/maxresdefault.jpg",
		},
	}
}
