package api

import (
	"io"
	"net/http"
	"strings"
	"time"

	"codeberg.org/video-prize-ranch/go-rimgu/types"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func FetchAlbum(albumID string) (types.Album, error) {
	// https://api.imgur.com/post/v1/albums/zk7mdKH?client_id=${CLIENT_ID}&include=media%2Caccount

	res, err := http.Get("https://api.imgur.com/post/v1/albums/" + albumID + "?client_id=" + viper.GetString("RIMGU_IMGUR_CLIENT_ID") + "&include=media%2Caccount")
	if err != nil {
		return types.Album{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return types.Album{}, err
	}

	data := gjson.Parse(string(body))

	media := make([]string, 0)
	data.Get("media").ForEach(
		func(key gjson.Result, value gjson.Result) bool {
			url := value.Get("url").String()
			url = strings.ReplaceAll(url, "https://i.imgur.com", "/media")
			media = append(media, url)

			return true
		},
	)

	createdAt, err := time.Parse("2006-01-02T15:04:05Z", data.Get("created_at").String())
	if err != nil {
		return types.Album{}, err
	}

	return types.Album{
		Id: data.Get("id").String(),
    Title: data.Get("title").String(),
    Views: data.Get("view_count").Int(),
    Upvotes: data.Get("upvote_count").Int(),
    Downvotes: data.Get("downvote_count").Int(),
    CreatedAt: createdAt.Format("January 2, 2006 3:04 PM"),
		Media: media,
	}, nil
}