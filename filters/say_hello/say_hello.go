package say_hello

import (
	"github.com/Oni-kuki/operative-framework/session"
)

type SayHelloFilter struct {
	session.SessionFilter
	Sess *session.Session `json:"-"`
}

func PushSayHelloFilter(s *session.Session) *SayHelloFilter {
	mod := SayHelloFilter{
		Sess: s,
	}
	mod.AddModule("instagram.feed")
	return &mod
}

func (filter *SayHelloFilter) Name() string {
	return "say_hello"
}

func (filter *SayHelloFilter) Description() string {
	return "Exemple filter"
}

func (filter *SayHelloFilter) Author() string {
	return "Tristan Granier"
}

func (filter *SayHelloFilter) Start(mod session.Module) {
	filter.Sess.Stream.Success("Filter as running successfully after module " + mod.Name() + " !")
}
