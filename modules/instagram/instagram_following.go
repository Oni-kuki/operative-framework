package instagram

import (
	"fmt"
	"os"

	"github.com/graniet/go-pretty/table"
	"github.com/graniet/operative-framework/session"
	"gopkg.in/ahmdrz/goinsta.v3"
)

type InstagramFollowing struct {
	session.SessionModule
	Sess *session.Session `json:"-"`
}

func PushInstagramFollowingModule(s *session.Session) *InstagramFollowing {
	mod := InstagramFollowing{
		Sess: s,
	}

	mod.CreateNewParam("TARGET", "INSTAGRAM USER ACCOUNT", "", true, session.STRING)
	return &mod
}

func (module *InstagramFollowing) Name() string {
	return "instagram.following"
}

func (module *InstagramFollowing) Description() string {
	return "Get following for target user instagram account"
}

func (module *InstagramFollowing) Author() string {
	return "Tristan Granier"
}

func (module *InstagramFollowing) GetType() []string {
	return []string{
		session.T_TARGET_INSTAGRAM,
	}
}

func (module *InstagramFollowing) GetInformation() session.ModuleInformation {
	information := session.ModuleInformation{
		Name:        module.Name(),
		Description: module.Description(),
		Author:      module.Author(),
		Type:        module.GetType(),
		Parameters:  module.Parameters,
	}
	return information
}

func (module *InstagramFollowing) Start() {

	trg, err := module.GetParameter("TARGET")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	target, err2 := module.Sess.GetTarget(trg.Value)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	insta := goinsta.New(module.Sess.Config.Instagram.Login, module.Sess.Config.Instagram.Password)

	if err := insta.Login(); err != nil {
		fmt.Println(err)
		return
	}

	profil, err := insta.Profiles.ByName(target.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t := module.Sess.Stream.GenerateTable()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Username",
		"Full Name",
	})
	followings := profil.Following()
	for followings.Next() {
		for _, following := range followings.Users {
			t.AppendRow(table.Row{
				following.Username,
				following.FullName,
			})

			result := target.NewResult()
			result.Set("username", following.Username)
			result.Set("full_name", following.FullName)
			result.Save(module, target)
		}
	}
	module.Sess.Stream.Render(t)
}
