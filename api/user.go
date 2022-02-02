package api

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"codeberg.org/video-prize-ranch/rimgo/types"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func FetchUser(username string) (types.User, error) {
	res, err := http.Get("https://api.imgur.com/account/v1/accounts/" + username + "?client_id=" + viper.GetString("RIMGU_IMGUR_CLIENT_ID"))
	if err != nil {
		return types.User{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return types.User{}, err
	}

	data := gjson.Parse(string(body))

	createdTime, _ := time.Parse(time.RFC3339, data.Get("created_at").String())

	return types.User{
		Id:        data.Get("id").Int(),
		Bio:       data.Get("bio").String(),
		Username:  data.Get("username").String(),
		Points:    data.Get("reputation_count").Int(),
		Cover:     strings.ReplaceAll(data.Get("cover_url").String(), "https://imgur.com", ""),
		Avatar:    strings.ReplaceAll(data.Get("avatar_url").String(), "https://i.imgur.com", ""),
		CreatedAt: createdTime.Format("January 2, 2006"),
	}, nil
}

func FetchSubmissions(username string, sort string, page string) ([]types.Submission, error) {
	res, err := http.Get("https://api.imgur.com/3/account/" + username + "/submissions/" + page + "/" + sort + "?album_previews=1&client_id=" + viper.GetString("RIMGU_IMGUR_CLIENT_ID"))
	if err != nil {
		return []types.Submission{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []types.Submission{}, err
	}

	data := gjson.Parse(string(body))

	submissions := []types.Submission{}

	wg := sync.WaitGroup{}
	data.Get("data").ForEach(
		func(key, value gjson.Result) bool {
			wg.Add(1)

			go func() {
				defer wg.Done()

				cover := value.Get("images.#(id==\"" + value.Get("cover").String() + "\")")

				submissions = append(submissions, types.Submission{
					Id:    value.Get("id").String(),
					Link:  strings.ReplaceAll(value.Get("link").String(), "https://imgur.com", ""),
					Title: value.Get("title").String(),
					Cover: types.Media{
						Id:          cover.Get("id").String(),
						Description: cover.Get("description").String(),
						Type:        strings.Split(cover.Get("type").String(), "/")[0],
						Url:         strings.ReplaceAll(cover.Get("link").String(), "https://i.imgur.com", ""),
					},
					Points:    cover.Get("points").Int(),
					Upvotes:   cover.Get("ups").Int(),
					Downvotes: cover.Get("downs").Int(),
					Comments:  cover.Get("comment_count").Int(),
					Views:     cover.Get("views").Int(),
					IsAlbum:   cover.Get("is_album").Bool(),
				})
			}()

			return true
		},
	)
	wg.Wait()
	return submissions, nil
}
