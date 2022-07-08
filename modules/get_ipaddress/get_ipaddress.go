package get_ipaddress

import (
	"net"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	"github.com/Oni-kuki/operative-framework/session"
)

type GetIpAddressModule struct {
	session.SessionModule
	sess   *session.Session `json:"-"`
	Stream *session.Stream  `json:"-"`
}

func PushGetIpAddressModule(s *session.Session) *GetIpAddressModule {
	mod := GetIpAddressModule{
		sess:   s,
		Stream: &s.Stream,
	}

	mod.CreateNewParam("TARGET", "Website domain target", "", true, session.STRING)
	return &mod
}

func (module *GetIpAddressModule) Name() string {
	return "get.ip_address"
}

func (module *GetIpAddressModule) Author() string {
	return "Tristan Granier"
}

func (module *GetIpAddressModule) Description() string {
	return "Get internet protocol address from specific target"
}

func (module *GetIpAddressModule) GetType() []string {
	return []string{
		session.T_TARGET_WEBSITE,
	}
}

func (module *GetIpAddressModule) GetInformation() session.ModuleInformation {
	information := session.ModuleInformation{
		Name:        module.Name(),
		Description: module.Description(),
		Author:      module.Author(),
		Type:        module.GetType(),
		Parameters:  module.Parameters,
	}
	return information
}

func (module *GetIpAddressModule) Start() {

	trg, err := module.GetParameter("TARGET")
	if err != nil {
		module.sess.Stream.Error(err.Error())
		return
	}

	target, err := module.sess.GetTarget(trg.Value)
	if err != nil {
		module.sess.Stream.Error(err.Error())
		return
	}

	if strings.Contains(target.GetName(), "://") {
		expProto := strings.Split(target.GetName(), "://")
		proto := expProto[0]
		expURL := ""
		if strings.Contains(target.GetName(), "/") {
			expURL = strings.Split(expProto[1], "/")[0]
			target.Name = proto + "://" + expURL
		}
	} else {

		if strings.Contains(target.GetName(), "/") {
			expURL := strings.Split(target.GetName(), "/")[0]
			target.Name = "https://" + expURL
		}
	}

	if strings.Contains(target.GetName(), "://") {
		target.Name = strings.Split(target.Name, "://")[1]
	}

	ipAddress, _ := net.LookupIP(target.GetName()) // take from 1st argument
	if len(ipAddress) > 0 {
		t := module.sess.Stream.GenerateTable()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{
			"IP",
		})
		for _, ip := range ipAddress {
			t.AppendRow(table.Row{
				ip.String(),
			})

			result := target.NewResult()
			result.Set("IP", ip.String())
			result.Save(module, target)
			module.Results = append(module.Results, ip.String())
		}
		module.sess.Stream.Render(t)
	} else {
		module.sess.Stream.Error("No result found")
		return
	}
}
