package engine

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/Oni-kuki/go-pretty/table"
	"github.com/Oni-kuki/operative-framework/api"
	"github.com/Oni-kuki/operative-framework/session"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/color"
)

// Checking If Input As Default Command
func CommandBase(line string, s *session.Session) bool {

	// Default Command
	if line == "info session" {
		ViewInformation(s)
		return true
	} else if line == "info api" {
		ViewApiInformation(s)
		return true
	} else if line == "env" {
		ViewEnvironment(s)
		return true
	} else if line == "clear" {
		s.ClearScreen()
		return true
	} else if line == "modules" {
		s.ListModules()
		return true
	}
	return false
}

// View Environment File Argument
func ViewEnvironment(s *session.Session) {
	t := s.Stream.GenerateTable()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Name",
		"Value",
	})
	mp, err := godotenv.Read(s.Config.Common.ConfigurationFile)
	if err == nil {
		for name, value := range mp {
			t.AppendRow(table.Row{
				name,
				value,
			})
		}
	}
	s.Stream.Render(t)
	s.Stream.Standard("Environment loaded at '" + s.Config.Common.ConfigurationFile + "'")
}

// View Session Information
func ViewInformation(s *session.Session) {
	t := s.Stream.GenerateTable()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Name",
		"Value",
	})
	apiStatus := color.Red("offline")
	if s.Information.ApiStatus {
		apiStatus = color.Green("online")
	}

	trackerStatus := color.Red("offline")
	if s.Information.TrackerStatus {
		trackerStatus = color.Green("online")
	}

	t.AppendRow(table.Row{
		"OPF",
		s.Config.Common.BaseDirectory,
	})
	t.AppendRow(table.Row{
		"CONFIGURATION",
		s.Config.Common.ConfigurationFile,
	})
	t.AppendRow(table.Row{
		"CRON JOB(S)",
		s.Config.Common.ConfigurationJobs,
	})
	t.AppendRow(table.Row{
		"EXPORT",
		s.Config.Common.ExportDirectory,
	})
	t.AppendRow(table.Row{
		"API",
		apiStatus,
	})
	t.AppendRow(table.Row{
		"TRACKER",
		trackerStatus,
	})
	t.AppendRow(table.Row{
		"EVENT(S)",
		s.Information.Event,
	})
	t.AppendRow(table.Row{
		"MODULE(S)",
		len(s.Modules),
	})
	t.AppendRow(table.Row{
		"TARGET(S)",
		len(s.Targets),
	})
	s.Stream.Render(t)
}

// View Api EndPoints Information
func ViewApiInformation(s *session.Session) {
	a := api.PushARestFul(s)
	r := a.LoadRouter()
	ta := s.Stream.GenerateTable()
	ta.SetOutputMirror(os.Stdout)
	ta.AppendHeader(table.Row{
		"Endpoint",
	})
	_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		ta.AppendRow(table.Row{
			t,
		})
		return nil
	})
	s.Stream.Render(ta)
}
