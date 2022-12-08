package models

import "github.com/spf13/viper"

type Catalog struct {
	Id    int64 `gorm:"primaryKey"`
	Title string
	File  string
}

func (c *Catalog) CatalogFilePath() string {
	if c.File == "" {
		return c.File
	}
	return viper.GetString("files_paths.media_url") + c.File
}
