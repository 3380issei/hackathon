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

// 時間表示のstring化と成型
func time_to_str(t time.Time) string {
	return t.Format("2006年1月2日 15:04")
}

func (xs *XService) Post(user model.User, schedule model.Schedule) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//環境変数回りを理解，実行環境にenvファイルが置け次第即変更すること
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

	post_content += user.Name + "さんは目的地に宣言通り到着できませんでした．\n"
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
