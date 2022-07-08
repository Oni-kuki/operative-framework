package modules

import (
	"github.com/Oni-kuki/operative-framework/modules/abusech"
	"github.com/Oni-kuki/operative-framework/modules/account_checker"
	"github.com/Oni-kuki/operative-framework/modules/bing_vhost"
	"github.com/Oni-kuki/operative-framework/modules/cookies"
	"github.com/Oni-kuki/operative-framework/modules/darksearch"
	"github.com/Oni-kuki/operative-framework/modules/directory_search"
	"github.com/Oni-kuki/operative-framework/modules/find"
	"github.com/Oni-kuki/operative-framework/modules/get_ipaddress"
	"github.com/Oni-kuki/operative-framework/modules/google"
	"github.com/Oni-kuki/operative-framework/modules/header_retrieval"
	"github.com/Oni-kuki/operative-framework/modules/image_reverse_search"
	"github.com/Oni-kuki/operative-framework/modules/info_greffe"
	"github.com/Oni-kuki/operative-framework/modules/instagram"
	"github.com/Oni-kuki/operative-framework/modules/ip_information"
	"github.com/Oni-kuki/operative-framework/modules/linkedin_search"
	"github.com/Oni-kuki/operative-framework/modules/mac_vendor"
	"github.com/Oni-kuki/operative-framework/modules/metatag_spider"
	"github.com/Oni-kuki/operative-framework/modules/module_base/session_help"
	"github.com/Oni-kuki/operative-framework/modules/module_base/session_import"
	"github.com/Oni-kuki/operative-framework/modules/module_base/session_stream"
	"github.com/Oni-kuki/operative-framework/modules/pastebin"
	"github.com/Oni-kuki/operative-framework/modules/pastebin_email"
	"github.com/Oni-kuki/operative-framework/modules/phone_buster"
	"github.com/Oni-kuki/operative-framework/modules/phone_generator"
	"github.com/Oni-kuki/operative-framework/modules/phone_generator_fr"
	"github.com/Oni-kuki/operative-framework/modules/pictures"
	"github.com/Oni-kuki/operative-framework/modules/regex"
	"github.com/Oni-kuki/operative-framework/modules/report"
	"github.com/Oni-kuki/operative-framework/modules/sample"
	"github.com/Oni-kuki/operative-framework/modules/searchsploit"
	"github.com/Oni-kuki/operative-framework/modules/societe_com"
	"github.com/Oni-kuki/operative-framework/modules/system"
	"github.com/Oni-kuki/operative-framework/modules/tools_suggester"
	"github.com/Oni-kuki/operative-framework/modules/twitter"
	"github.com/Oni-kuki/operative-framework/modules/viewdns_search"
	"github.com/Oni-kuki/operative-framework/modules/whatsapp"
	"github.com/Oni-kuki/operative-framework/session"
)

func LoadModules(s *session.Session) {
	s.Modules = append(s.Modules, account_checker.PushAccountCheckerModule(s))
	s.Modules = append(s.Modules, bing_vhost.PushBingVirtualHostModule(s))
	s.Modules = append(s.Modules, abusech.PushAbuseChModule(s))
	s.Modules = append(s.Modules, find.PushFindModule(s))
	s.Modules = append(s.Modules, darksearch.PushDarkSearchModule(s))
	s.Modules = append(s.Modules, directory_search.PushModuleDirectorySearch(s))
	s.Modules = append(s.Modules, get_ipaddress.PushGetIpAddressModule(s))
	s.Modules = append(s.Modules, header_retrieval.PushModuleHeaderRetrieval(s))
	s.Modules = append(s.Modules, google.PushGoogleSearchModule(s))
	s.Modules = append(s.Modules, google.PushGoogleTwitterModule(s))
	s.Modules = append(s.Modules, google.PushGoogleDorkModule(s))
	s.Modules = append(s.Modules, cookies.PushGetCookiesModule(s))
	s.Modules = append(s.Modules, instagram.PushInstagramFollowersModule(s))
	s.Modules = append(s.Modules, instagram.PushInstagramFeedModule(s))
	s.Modules = append(s.Modules, instagram.PushInstagramFollowingModule(s))
	s.Modules = append(s.Modules, instagram.PushInstagramFriendsModule(s))
	s.Modules = append(s.Modules, instagram.PushInstagramInfoModule(s))
	s.Modules = append(s.Modules, instagram.PushInstagramCommentsModule(s))
	s.Modules = append(s.Modules, image_reverse_search.PushImageReverseModule(s))
	s.Modules = append(s.Modules, session_import.PushModuleImport(s))
	s.Modules = append(s.Modules, info_greffe.PushInfoGreffeRegistrationModule(s))
	s.Modules = append(s.Modules, ip_information.PushIpInformationModule(s))
	s.Modules = append(s.Modules, linkedin_search.PushLinkedinSearchModule(s))
	s.Modules = append(s.Modules, mac_vendor.PushMacVendorModule(s))
	s.Modules = append(s.Modules, metatag_spider.PushMetaTagModule(s))
	s.Modules = append(s.Modules, pastebin_email.PushPasteBinEmailModule(s))
	s.Modules = append(s.Modules, pastebin.PushPasteBinModule(s))
	s.Modules = append(s.Modules, phone_buster.PushPhoneBusterModule(s))
	s.Modules = append(s.Modules, phone_generator.PushPhoneGeneratorModule(s))
	s.Modules = append(s.Modules, phone_generator_fr.PushPhoneGeneratorFrModule(s))
	s.Modules = append(s.Modules, pictures.PushPictureExifModule(s))
	s.Modules = append(s.Modules, regex.PushFindWithRegexModule(s))
	s.Modules = append(s.Modules, report.PushReportPDFModule(s))
	s.Modules = append(s.Modules, report.PushReportJSONModule(s))
	s.Modules = append(s.Modules, sample.PushSampleModuleModule(s))
	s.Modules = append(s.Modules, system.PushSystemModuleModule(s))
	s.Modules = append(s.Modules, session_help.PushModuleHelp(s))
	s.Modules = append(s.Modules, session_stream.PushSessionStreamModule(s))
	s.Modules = append(s.Modules, societe_com.PushSocieteComModuleModule(s))
	s.Modules = append(s.Modules, searchsploit.PushSearchSploitModule(s))
	s.Modules = append(s.Modules, twitter.PushTwitterFollowerModule(s))
	s.Modules = append(s.Modules, twitter.PushTwitterRetweetModule(s))
	s.Modules = append(s.Modules, twitter.PushTwitterFollowingModule(s))
	s.Modules = append(s.Modules, twitter.PushTwitterInfoModule(s))
	s.Modules = append(s.Modules, twitter.PushTwitterGeoTweetModule(s))
	s.Modules = append(s.Modules, twitter.PushTwitterSearchModule(s))
	s.Modules = append(s.Modules, tools_suggester.PushModuleToolsSuggester(s))
	s.Modules = append(s.Modules, viewdns_search.PushWSearchModule(s))
	s.Modules = append(s.Modules, whatsapp.PushWhatsappExtractorModule(s))

	for _, mod := range s.Modules {
		for _, tp := range mod.GetType() {
			s.PushType(tp)
		}
		mod.CreateNewParam("FILTER", "Use module filter after execution", "", false, session.STRING)
		mod.CreateNewParam("BACKGROUND", "Run this task in background", "false", false, session.BOOL)
		mod.CreateNewParam("DISABLE_OUTPUT", "Display module result in stdout", "false", false, session.BOOL)
	}
}
