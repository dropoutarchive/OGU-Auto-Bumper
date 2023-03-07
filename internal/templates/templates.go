package templates

import (
	"github.com/google/uuid"
	"github.com/syrinsec/termfx"
	"ogu.gg/autobumper/internal/config"
	"ogu.gg/autobumper/internal/utils"
)

func Render(template string) string {
	registry := termfx.New()
	registry.RegisterVariable("uuid", uuid.New().String())

	content, _ := registry.ExecuteString(string(utils.ReadFile(config.Conf.Post.Content)))
	return content
}
