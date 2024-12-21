package youtube_test

import (
	"fmt"
	"net/url"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/roman-kart/gocommons/youtube"
)

var globalString string

func TestGetVideoId(t *testing.T) {
	t.Parallel()

	data := getVideIdDataStable()

	for _, d := range data {
		testMsg := fmt.Sprintf("Test data: %+v", d)

		videoUrl, err := url.Parse(d.Url)
		require.NoError(t, err, testMsg)

		videoId, err := youtube.GetVideoId(videoUrl)
		require.ErrorIs(t, err, d.Err, testMsg)
		require.Equal(t, d.VideoId, videoId, testMsg)
	}
}

func TestGetVideoIdRandomVideoId(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Log("Skip because short mode")
		return
	}

	data := getTestVideoIdDataRandomValid(10_000)

	for _, d := range data {
		testMsg := fmt.Sprintf("Test data: %+v", d)

		videoUrl, err := url.Parse(d.Url)
		require.NoError(t, err)

		videoId, err := youtube.GetVideoId(videoUrl)
		require.NoError(t, err, testMsg)
		require.Equal(t, d.VideoId, videoId, testMsg)
	}
}

func BenchmarkGetVideoId(b *testing.B) {
	var local string

	b.StopTimer()

	videoUrl, err := url.Parse("https://www.youtube.com/watch?v=jNQXAC9IVRw")
	if err != nil {
		b.Fatal(err)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		videoId, err := youtube.GetVideoId(videoUrl)
		if err != nil {
			b.Fatal(err)
		}

		local = videoId
	}

	globalString = local
}

func TestGetThumbnailRequestUrlNotPanics(t *testing.T) {
	t.Parallel()

	require.NotPanics(t, func() {
		u := youtube.GetThumbnailRequestUrl("test", "test")
		runtime.KeepAlive(u)
	})
}

func TestGetThumbnailRequestUrlStableData(t *testing.T) {
	t.Parallel()

	data := getTestGetThumbnailRequestUrlStable()

	for _, d := range data {
		u := youtube.GetThumbnailRequestUrl(d.VideoId, d.Quality)
		require.Equal(t, d.Result, u.String())
	}
}
