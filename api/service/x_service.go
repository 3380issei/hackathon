package service

import (
	"api/model"
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"

	"github.com/michimani/gotwi"
)

type XService struct {
}

func NewXService() XService {
	return XService{}
}

type XServiceInterface interface {
	Post(schedule model.Schedule) error
}

func (xs *XService) Post(schedule model.Schedule) error {
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

	post_content += "Test投稿です\n"
	post_content += schedule.Destination + "(" + fmt.Sprint(schedule.Latitude) + ", " + fmt.Sprint(schedule.Longitude) + ")\n"
	post_content += "期限は" + fmt.Sprint(schedule.Deadline) + "まで．"

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
