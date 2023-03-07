package regexs

func ParsePage(page string) Values {
	return Values{
		PostKey:  myPostKeyRegex.FindStringSubmatch(page)[1],
		Subject:  subjectRegex.FindStringSubmatch(page)[1],
		PostHash: postHashRegex.FindStringSubmatch(page)[1],
		LastPID:  lastPIDRegex.FindStringSubmatch(page)[1],
		TID:      tidRegex.FindStringSubmatch(page)[1],
	}
}
