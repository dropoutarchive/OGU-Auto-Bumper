package ogu

import (
	"bytes"
	"errors"
	"net/url"
	"ogu.gg/autobumper/internal/http"
	"ogu.gg/autobumper/internal/regexs"
	"strings"
)

func (c *OGU) GetParameters(postURL string) (*regexs.Values, error) {
	client, err := http.New()
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"authority":                 "ogu.gg",
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
		"accept-language":           "en-GB,en;q=0.9",
		"cache-control":             "max-age=0",
		"cookie":                    "ogumybbuser=" + c.Session,
		"referer":                   "https://ogu.gg/dropout",
		"sec-ch-ua":                 "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Brave\";v=\"108\"",
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        "\"Windows\"",
		"sec-fetch-dest":            "document",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-site":            "same-origin",
		"sec-fetch-user":            "?1",
		"sec-gpc":                   "1",
		"upgrade-insecure-requests": "1",
		"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
	}

	res, err := client.Request("GET", postURL, nil, headers)
	if err != nil {
		return nil, err
	}

	if res.Status != 200 {
		return nil, errors.New("invalid post url")
	} else {
		params := regexs.ParsePage(res.Body)
		return &params, nil
	}
}

func (c *OGU) PostReply(content string, postParameters regexs.Values) error {
	client, err := http.New()
	if err != nil {
		return err
	}

	query := url.Values{
		"my_post_key":            {postParameters.PostKey},
		"subject":                {postParameters.Subject},
		"action":                 {"do_newreply"},
		"posthash":               {postParameters.PostHash},
		"quoted_ids":             {""},
		"lastpid":                {postParameters.LastPID},
		"from_page":              {"1"},
		"tid":                    {postParameters.TID},
		"method":                 {"quickreply"},
		"message":                {content},
		"postoptions[signature]": {"1"},
	}

	headers := map[string]string{
		"authority":                 "ogu.gg",
		"accept":                    "application/json, text/html, */*; q=0.01",
		"accept-language":           "en-GB,en;q=0.9",
		"cache-control":             "max-age=0",
		"content-type":              "application/x-www-form-urlencoded; charset=UTF-8",
		"cookie":                    "ogumybbuser=" + c.Session,
		"sec-ch-ua":                 "\"Not?A_Brand\";v=\"8\", \"Chromium\";v=\"108\", \"Brave\";v=\"108\"",
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        "\"Windows\"",
		"sec-fetch-dest":            "document",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-site":            "same-origin",
		"sec-fetch-user":            "?1",
		"sec-gpc":                   "1",
		"upgrade-insecure-requests": "1",
		"user-agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36",
	}

	res, err := client.Request("POST", "https://ogu.gg/newreply.php?ajax=1", bytes.NewBuffer([]byte(query.Encode())), headers)
	if err != nil {
		return err
	}

	if strings.Contains(res.Body, "error") {
		return errors.New(res.Body[:35])
	} else {
		return nil
	}
}
