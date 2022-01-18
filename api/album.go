package api

import (
	"io"
	"net/http"
	"strings"

	"codeberg.org/video-prize-ranch/rimgo/types"
	"codeberg.org/video-prize-ranch/rimgo/utils"
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

	media := make([]types.Media, 0)
	data.Get("media").ForEach(
		func(key gjson.Result, value gjson.Result) bool {
			url := value.Get("url").String()
			url = strings.ReplaceAll(url, "https://i.imgur.com", "")

			media = append(media, types.Media{
				Id:          value.Get("id").String(),
				Name:        value.Get("name").String(),
				MimeType:    value.Get("mime_type").String(),
				Type:        value.Get("type").String(),
				Title:       value.Get("metadata.title").String(),
				Description: value.Get("metadata.description").String(),
				Url:         url,
			})

			return true
		},
	)

	createdAt, err := utils.FormatDate(data.Get("created_at").String())
	if err != nil {
		return types.Album{}, err
	}

	return types.Album{
		Id:        data.Get("id").String(),
		Title:     data.Get("title").String(),
		Views:     data.Get("view_count").Int(),
		Upvotes:   data.Get("upvote_count").Int(),
		Downvotes: data.Get("downvote_count").Int(),
		Comments:  data.Get("comment_count").Int(),
		CreatedAt: createdAt,
		Media:     media,
	}, nil
}
