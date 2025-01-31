package twitter

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/graniet/go-pretty/table"
	"github.com/Oni-kuki/operative-framework/session"
)

type TwitterRetweet struct {
	session.SessionModule
	Sess *session.Session `json:"-"`
}

func PushTwitterRetweetModule(s *session.Session) *TwitterRetweet {
	mod := TwitterRetweet{
		Sess: s,
	}

	mod.CreateNewParam("TARGET", "TWITTER USER SCREEN NAME", "", true, session.STRING)
	mod.CreateNewParam("WITHRETWEET", "Include RT ?", "true", false, session.STRING)
	mod.CreateNewParam("COUNT", "Number of tweets", "100", false, session.STRING)
	return &mod
}

func (module *TwitterRetweet) Name() string {
	return "twitter.tweets"
}

func (module *TwitterRetweet) Description() string {
	return "Get (re)tweets from target user twitter account"
}

func (module *TwitterRetweet) Author() string {
	return "Tristan Granier"
}

func (module *TwitterRetweet) GetType() []string {
	return []string{
		session.T_TARGET_TWITTER,
	}
}

func (module *TwitterRetweet) GetInformation() session.ModuleInformation {
	information := session.ModuleInformation{
		Name:        module.Name(),
		Description: module.Description(),
		Author:      module.Author(),
		Type:        module.GetType(),
		Parameters:  module.Parameters,
	}
	return information
}

func (module *TwitterRetweet) Start() {

	trg, err := module.GetParameter("TARGET")
	if err != nil {
		module.Sess.Stream.Error(err.Error())
		return
	}

	target, err := module.Sess.GetTarget(trg.Value)
	if err != nil {
		module.Sess.Stream.Error(err.Error())
		return
	}

	argumentRT, err := module.GetParameter("WITHRETWEET")
	if err != nil {
		module.Sess.Stream.Error(err.Error())
		return
	}

	if argumentRT.Value != "false" && argumentRT.Value != "true" {
		module.Sess.Stream.Error("Please set correct value for 'WithRetweet' argument.")
		return
	}

	argumentCount, err := module.GetParameter("COUNT")
	if err != nil {
		module.Sess.Stream.Error(err.Error())
		return
	}

	api := anaconda.NewTwitterApiWithCredentials(module.Sess.Config.Twitter.Password, module.Sess.Config.Twitter.Api.SKey, module.Sess.Config.Twitter.Login, module.Sess.Config.Twitter.Api.Key)
	v := url.Values{}
	user, err := api.GetUserSearch(target.Name, v)
	if err != nil {
		module.Sess.Stream.Error(err.Error())
		return
	}
	u := user[0]
	v.Set("screen_name", u.ScreenName)
	v.Set("include_rts", argumentRT.Value)
	v.Set("count", argumentCount.Value)
	retweets, err := api.GetUserTimeline(v)

	t := module.Sess.Stream.GenerateTable()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"text",
		"user",
		"type",
		"latitude",
		"longitude",
		"Date",
	})
	t.SetAllowedColumnLengths([]int{40, 0})

	for _, tweet := range retweets {
		var text string
		var user string
		var tweetType string
		if strings.Contains(tweet.Text, "RT @") {
			text = strings.TrimSpace(strings.Split(tweet.FullText, ":")[1])
			t := text[:len(text)-3]
			text = t
			if len(tweet.Entities.User_mentions) > 0 {
				user = tweet.Entities.User_mentions[0].Screen_name
			} else {
				user = tweet.User.ScreenName
			}
			tweetType = "RT"
		} else {
			user = tweet.User.ScreenName
			text = strings.TrimSpace(tweet.FullText)
			tweetType = "T"
		}
		latitudeFloat, err := tweet.Latitude()
		latitude := ""
		if err != nil {
			latitude = "-"
		} else {
			latitude = fmt.Sprintf("%f", latitudeFloat)
		}

		longitudeFloat, err := tweet.Longitude()
		longitude := ""
		if err != nil {
			longitude = "-"
		} else {
			longitude = fmt.Sprintf("%f", longitudeFloat)
		}
		t.AppendRow(table.Row{
			text,
			user,
			tweetType,
			latitude,
			longitude,
			tweet.CreatedAt,
		})

		result := target.NewResult()
		result.Set("tweet", text)
		result.Set("user", user)
		result.Set("date", tweet.CreatedAt)
		result.Set("type", tweetType)
		result.Set("latitude", latitude)
		result.Set("longitude", longitude)
		result.Save(module, target)
	}
	module.Sess.Stream.Render(t)

}
