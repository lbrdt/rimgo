package api

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"codeberg.org/video-prize-ranch/rimgo/types"
	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/dustin/go-humanize"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func FetchComments(galleryID string) ([]types.Comment, error) {
	// https://api.imgur.com/comment/v1/comments?client_id=546c25a59c58ad7&filter[post]=eq:g1bk7CB&include=account&per_page=30&sort=best

	res, err := http.Get("https://api.imgur.com/comment/v1/comments?client_id=" + viper.GetString("RIMGU_IMGUR_CLIENT_ID") + "&filter[post]=eq:" + galleryID + "&include=account,adconfig&per_page=30&sort=best")
	if err != nil {
		return []types.Comment{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []types.Comment{}, err
	}

	data := gjson.Parse(string(body))

	wg := sync.WaitGroup{}
	comments := make([]types.Comment, 0)
	data.Get("data").ForEach(
		func(key, value gjson.Result) bool {
			wg.Add(1)

			go func() {
				defer wg.Done()
				comments = append(comments, ParseComment(value))
			}()

			return true
		},
	)
	wg.Wait()

	return comments, nil
}

func ParseComment(data gjson.Result) types.Comment {
	createdTime, _ := time.Parse("2006-01-02T15:04:05Z", data.Get("created_at").String())
	createdAt := createdTime.Format("January 2, 2006 3:04 PM")
	updatedAt, _ := utils.FormatDate(data.Get("updated_at").String())
	deletedAt, _ := utils.FormatDate(data.Get("deleted_at").String())

	userAvatar := strings.ReplaceAll(data.Get("account.avatar").String(), "https://i.imgur.com", "")
	if viper.GetBool("CF_ALL_MEDIA") {
		userAvatar = viper.GetString("CF_MEDIA_DISTRIBUTION") + userAvatar
	}

	wg := sync.WaitGroup{}
	comments := make([]types.Comment, 0)
	data.Get("comments").ForEach(
		func(key, value gjson.Result) bool {
			wg.Add(1)

			go func() {
				defer wg.Done()
				comments = append(comments, ParseComment(value))
			}()

			return true
		},
	)
	wg.Wait()

	return types.Comment{
		Comments: comments,
		User: types.User{
			Id:       data.Get("account.id").Int(),
			Username: data.Get("account.username").String(),
			Avatar:   userAvatar,
		},
		Id:        data.Get("id").String(),
		Comment:   data.Get("comment").String(),
		Upvotes:   data.Get("upvote_count").Int(),
		Downvotes: data.Get("downvote_count").Int(),
		Platform:  data.Get("platform").String(),
		CreatedAt: createdAt,
		RelTime:   humanize.Time(createdTime),
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
