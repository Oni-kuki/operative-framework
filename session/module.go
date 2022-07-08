package session

import (
	"errors"
	"github.com/graniet/go-pretty/table"
	"github.com/joho/godotenv"
	"os"
	"os/user"
	"strconv"
	"strings"
)

const (
	INT    = 1
	STRING = 2
	BOOL   = 3
	FLOAT  = 4
)

type Module interface {
	Start()
	Name() string
	Author() string
	Description() string
	GetType() []string
	ListArguments()
	GetExport() []OpfResults
	SetExport(result OpfResults)
	GetResults() []string
	GetInformation() ModuleInformation
	CheckRequired() bool
	SetParameter(name string, value string) (bool, error)
	GetParameter(name string) (Param, error)
	GetAllParameters() []Param
	WithProgram(name string) bool
	GetExternal() []string
	CreateNewParam(name string, description string, value string, isRequired bool, paramType int)
}

type ModuleEvent struct {
	ModuleName string
	Results    interface{}
}

type Param struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Value       string `json:"value"`
	IsRequired  bool   `json:"is_required"`
	ParamType   int    `json:"param_type"`
}

type SessionModule struct {
	Module
	Export     []OpfResults
	Parameters []Param  `json:"parameters"`
	History    []string `json:"history"`
	External   []string `json:"external"`
	Results    []string
}

type ModuleInformation struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	Type        []string `json:"type"`
	Parameters  []Param  `json:"parameters"`
}

func (s *Session) SearchModule(name string) (Module, error) {
	for _, module := range s.Modules {
		if module.Name() == name {
			return module, nil
		}
	}
	return nil, errors.New("error: This module not found")
}

func (module *SessionModule) GetParameter(name string) (Param, error) {
	for _, param := range module.Parameters {
		if strings.ToUpper(param.Name) == strings.ToUpper(name) {
			return param, nil
		}
	}
	return Param{}, errors.New("parameter not found")
}

func (module *SessionModule) SetParameter(name string, value string) (bool, error) {
	for k, param := range module.Parameters {
		if strings.ToUpper(param.Name) == strings.ToUpper(name) {
			module.Parameters[k].Value = value
			return true, nil
		}
	}
	return false, errors.New("argument not found")
}

func (module *SessionModule) CheckRequired() bool {
	for _, param := range module.Parameters {
		if param.IsRequired == true {
			switch param.ParamType {
			case STRING:
				if param.Value == "" {
					return false
				}
			case INT:
				value, _ := strconv.Atoi(param.Value)
				if value == 0 {
					return false
				}
			case BOOL:
				value := strings.TrimSpace(param.Value)
				if value == "" {
					return false
				}
			}
		}
	}
	return true
}

func (module *SessionModule) CreateNewParam(name string, description string, value string, isRequired bool, paramType int) {
	newParam := Param{
		Name:        strings.ToUpper(name),
		Value:       value,
		Description: description,
		IsRequired:  isRequired,
		ParamType:   paramType,
	}
	module.Parameters = append(module.Parameters, newParam)
}

func (module *SessionModule) WithProgram(name string) bool {
	module.External = append(module.External, name)
	return true
}

func (module *SessionModule) ListArguments() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetAllowedColumnLengths([]int{40, 40})
	t.AppendHeader(table.Row{"argument", "description", "value", "required", "type"})
	if len(module.Parameters) > 0 {
		for _, argument := range module.Parameters {
			argumentType := ""
			argumentRequired := ""

			if argument.ParamType == STRING {
				argumentType = "STRING"
			} else if argument.ParamType == INT {
				argumentType = "INTEGER"
			} else if argument.ParamType == BOOL {
				argumentType = "BOOLEAN"
			}

			if argument.IsRequired == true {
				argumentRequired = "YES"
			} else {
				argumentRequired = "NO"
			}

			if argument.Value == "" {
				argument.Value = "NO DEFAULT"
			}
			t.AppendRow([]interface{}{argument.Name, argument.Description, argument.Value, argumentRequired, argumentType})
		}
	} else {
		t.AppendRow([]interface{}{"No argument."})
	}
	t.Render()
}

func (module *SessionModule) SetExport(result OpfResults) {
	module.Export = append(module.Export, result)
}

func (module *SessionModule) GetExport() []OpfResults {
	return module.Export
}

func (module *SessionModule) GetAllParameters() []Param {
	return module.Parameters
}

func (module *SessionModule) GetResults() []string {
	return module.Results
}

func (module *SessionModule) GetExternal() []string {
	return module.External
}

func (s *Session) ListModules() {
	t := s.Stream.GenerateTable()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"NAME",
		"DESCRIPTION",
		"TYPE",
	})
	for _, module := range s.Modules {
		t.AppendRow(table.Row{
			module.Name(),
			module.Description(),
			strings.Join(module.GetType(), ","),
		})
	}
	s.Stream.Render(t)
}

func (s *Session) ParseModuleConfig() {
	u, _ := user.Current()
	for _, module := range s.Modules {
		fileName := u.HomeDir + "/.opf/external/modules/" + module.Name() + ".conf"
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			continue
		}
		configuration, err := godotenv.Read(fileName)
		if err == nil {
			s.Config.Modules[module.Name()] = configuration
		}
	}
}

func (s *Session) LoadModuleConfiguration(module string) (map[string]string, error) {

	if _, ok := s.Config.Modules[module]; !ok {
		return nil, errors.New("Configuration not found")
	}

	return s.Config.Modules[module], nil
}

func (s *Session) WithFilter(module Module) bool {
	useFilter := false
	filter, err := module.GetParameter("FILTER")
	if err == nil {
		_, err := s.SearchFilter(filter.Value)
		if err == nil {
			useFilter = true
		}
	}
	return useFilter
}
