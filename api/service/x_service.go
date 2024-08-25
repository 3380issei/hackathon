package service

import (
	"api/model"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"

	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
)

type XService struct {
}

func NewXService() XService {
	return XService{}
}

type XServiceInterface interface {
	Post(schedule model.Schedule) error
}

func (xs *XService) Post(user model.User, schedule model.Schedule) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	in := &gotwi.NewClientInput{
		AuthenticationMethod: gotwi.AuthenMethodOAuth1UserContext,
		OAuthToken:           "1820304527171743744-EhQgkpLo00G5O4TqwXBVJ0Bh3q7CUn",
		OAuthTokenSecret:     "eVyCRhIhhgOwlUaXRvB09XxyldYXXLtyNd3kruBhpuv4R",
	}

	c, err := gotwi.NewClient(in)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var post_content string

	post_content += user.Name + "さんは目的地に予定通り到着することができませんでした．\n"
	post_content += "到着予定場所は " + schedule.Destination + "(" + fmt.Sprint(schedule.Latitude) + ", " + fmt.Sprint(schedule.Longitude) + ")\n"
	post_content += "期限は " + time_to_str(schedule.Deadline) + " まででした．\n"
	post_content += "自分で設定しておいて何故時間通りにたどり着けないのです？"

	p := &types.CreateInput{
		Text: gotwi.String(post_content),
	}

	_, err = managetweet.Create(context.Background(), c, p)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func time_to_str(t time.Time) string {
	return t.Format("2006年1月2日 15:04")
}
