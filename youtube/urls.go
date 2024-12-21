package youtube

import (
	"errors"
	"net/url"
	"strings"
)

var (
	ErrHostIsNotYoutube = errors.New("host is not youtube")
	ErrWrongUrlFormat   = errors.New("wrong url format")
)

// GetVideoId returns youtube video id from videoUrl or error if occurred.
func GetVideoId(videoUrl *url.URL) (string, error) {
	// trim first `/` for correct string.Split working
	pathParts := strings.Split(strings.TrimPrefix(videoUrl.Path, "/"), "/")

	// remove empty path elements
	pathPartsTmp := make([]string, 0)

	for _, v := range pathParts {
		if v == "" {
			continue
		}

		pathPartsTmp = append(pathPartsTmp, v)
	}

	pathParts = pathPartsTmp

	switch strings.TrimPrefix(videoUrl.Host, "www.") {
	case "youtu.be":
		if len(pathParts) < 1 {
			return "", ErrWrongUrlFormat
		}

		return pathParts[0], nil
	case "youtube.com":
		vQueryParam := videoUrl.Query().Get("v")

		if vQueryParam != "" {
			return vQueryParam, nil
		}

		if len(pathParts) < 2 || pathParts[0] != "v" {
			return "", ErrWrongUrlFormat
		}

		return pathParts[1], nil
	default:
		return "", ErrHostIsNotYoutube
	}
}

type ThumbnailQuality string

const (
	ThumbnailQualityHQ     ThumbnailQuality = "hqdefault.jpg"
	ThumbnailQualityMQ     ThumbnailQuality = "mqdefault.jpg"
	ThumbnailQualitySD     ThumbnailQuality = "sddefault.jpg"
	ThumbnailQualityMaxRes ThumbnailQuality = "maxresdefault.jpg"
)

// GetThumbnailRequestUrl generate url for get thumbnail HTTP-request.
func GetThumbnailRequestUrl(videoId string, quality ThumbnailQuality) *url.URL {
	requestUrl, err := url.Parse("https://img.youtube.com/vi/")
	// panic because url must be parsed
	if err != nil {
		panic(err)
	}

	requestUrl = requestUrl.JoinPath(videoId, string(quality))

	return requestUrl
}
