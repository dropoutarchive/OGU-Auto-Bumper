package regexs

import "regexp"

var (
	myPostKeyRegex, _ = regexp.Compile(`"my_post_key" value="(.*)"`)
	subjectRegex, _   = regexp.Compile(`"subject" value="(.*)"`)
	postHashRegex, _  = regexp.Compile(`"posthash" value="(.*)" id="posthash"`)
	lastPIDRegex, _   = regexp.Compile(`"lastpid" value="(.*)"`)
	tidRegex, _       = regexp.Compile(`"tid" value="(.*)"`)
)

type Values struct {
	PostKey  string
	Subject  string
	PostHash string
	LastPID  string
	TID      string
}
