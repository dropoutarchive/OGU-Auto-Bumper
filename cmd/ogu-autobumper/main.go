package main

import (
	"fmt"
	"ogu.gg/autobumper/internal/config"
	"ogu.gg/autobumper/internal/logging"
	"ogu.gg/autobumper/internal/ogu"
	"ogu.gg/autobumper/internal/templates"
	"time"
)

func main() {
	client := ogu.OGU{Session: config.Conf.OGU.Session}
	duration := time.Duration(config.Conf.Post.Interval) * time.Second

	for {
		params, err := client.GetParameters(config.Conf.Post.URL)
		if err != nil {
			logging.Logger.Error().Err(err).Msg("Failed to fetch post, retrying in 25 seconds...")
			time.Sleep(25 * time.Second)
			continue
		}

		err = client.PostReply(templates.Render(config.Conf.Post.Content), *params)
		if err != nil {
			logging.Logger.Error().Err(err).Msg("Failed to post reply, retrying in 1 minute...")
			time.Sleep(1 * time.Minute)
			continue
		}

		logging.Logger.Info().Msg(fmt.Sprintf("Successfully bumped post, sleeping for %v...", duration.String()))
		time.Sleep(duration)
	}
}
