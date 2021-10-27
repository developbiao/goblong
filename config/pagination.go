package config

import "goblong/pkg/config"

func init() {
	config.Add("pagination", config.StrMap{
		// Default per page
		"perpage": 10,

		// URL page reprent "p"
		"url_query": "p",
	})

}
