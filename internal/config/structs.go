package config

type Config struct {
	OGU struct {
		Session string `toml:"session"`
	} `toml:"ogu"`

	Post struct {
		URL      string `toml:"url"`
		Content  string `toml:"content"`
		Interval int    `toml:"interval"`
	}
}
