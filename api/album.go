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

	album := types.Album{}
	if data.Get("shared_with_community").Bool() {
		album, err = FetchPosts(albumID)
	} else {
		album, err = ParseAlbum(data)
	}

	return album, err
}

func FetchPosts(albumID string) (types.Album, error) {
	res, err := http.Get("https://api.imgur.com/post/v1/posts/" + albumID + "?client_id=" + viper.GetString("RIMGU_IMGUR_CLIENT_ID") + "&include=media%2Caccount")
	if err != nil {
		return types.Album{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return types.Album{}, err
	}

	data := gjson.Parse(string(body))

	return ParseAlbum(data)
}

func ParseAlbum(data gjson.Result) (types.Album, error) {
	media := make([]types.Media, 0)
	data.Get("media").ForEach(
		func(key gjson.Result, value gjson.Result) bool {
			url := value.Get("url").String()
			url = strings.ReplaceAll(url, "https://i.imgur.com", "")

			if strings.HasSuffix(url, "mp4") || viper.GetBool("CF_ALL_MEDIA") {
				url = viper.GetString("CF_MEDIA_DISTRIBUTION") + url
			}

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
		Id:                  data.Get("id").String(),
		Title:               data.Get("title").String(),
		SharedWithCommunity: data.Get("shared_with_community").Bool(),
		Views:               data.Get("view_count").Int(),
		Upvotes:             data.Get("upvote_count").Int(),
		Downvotes:           data.Get("downvote_count").Int(),
		Comments:            data.Get("comment_count").Int(),
		CreatedAt:           createdAt,
		Media:               media,
	}, nil
}
