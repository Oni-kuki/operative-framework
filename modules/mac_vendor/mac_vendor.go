package mac_vendor

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/graniet/go-pretty/table"
	"github.com/Oni-kuki/operative-framework/session"
)

type MacVendorModule struct {
	session.SessionModule
	Session *session.Session `json:"-"`
	Stream  *session.Stream  `json:"-"`
}

func PushMacVendorModule(s *session.Session) *MacVendorModule {
	mod := MacVendorModule{
		Session: s,
		Stream:  &s.Stream,
	}

	mod.CreateNewParam("TARGET", "MAC Address from target", "", true, session.STRING)
	return &mod
}

func (module *MacVendorModule) Name() string {
	return "mac.get_vendor"
}

func (module *MacVendorModule) Description() string {
	return "Retrieve mac vendor information"
}

func (module *MacVendorModule) Author() string {
	return "Tristan Granier"
}

func (module *MacVendorModule) GetType() []string {
	return []string{
		session.T_TARGET_MAC,
	}
}

func (module *MacVendorModule) GetInformation() session.ModuleInformation {
	information := session.ModuleInformation{
		Name:        module.Name(),
		Description: module.Description(),
		Author:      module.Author(),
		Type:        module.GetType(),
		Parameters:  module.Parameters,
	}
	return information
}

func (module *MacVendorModule) Start() {
	target, err := module.GetParameter("TARGET")
	if err != nil {
		module.Stream.Error(err.Error())
		return
	}
	mac, err := module.Session.GetTarget(target.Value)
	if err != nil {
		module.Stream.Error(err.Error())
		return
	}

	url := "https://api.macvendors.com/" + mac.GetName()
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		module.Stream.Error(err.Error())
		return
	}

	response, err := client.Do(req)
	if err != nil {
		module.Stream.Error(err.Error())
		return
	}

	if response.StatusCode == 200 {
		macString, _ := ioutil.ReadAll(response.Body)

		t := module.Stream.GenerateTable()
		t.AppendHeader(table.Row{
			"MAC",
			"VENDOR",
		})
		t.SetOutputMirror(os.Stdout)
		t.AppendRow(table.Row{
			mac.GetName(),
			string(macString),
		})
		module.Stream.Render(t)

		result := mac.NewResult()
		result.Set("MAC", mac.GetName())
		result.Set("VENDOR", string(macString))
		result.Save(module, mac)
		return
	}

	module.Stream.Warning("No results found")
	return

}
