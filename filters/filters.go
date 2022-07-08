package filters

import (
	"github.com/Oni-kuki/operative-framework/filters/follower_to_screenName"
	"github.com/Oni-kuki/operative-framework/filters/say_hello"
	"github.com/Oni-kuki/operative-framework/session"
)

func LoadFilters(s *session.Session) {
	s.Filters = append(s.Filters, say_hello.PushSayHelloFilter(s))
	s.Filters = append(s.Filters, phone_to_instagram.PushPhoneToInstagramFilter(s))
	s.Filters = append(s.Filters, follower_to_screenName.PushFollowerScreenNameFilter(s))
}
