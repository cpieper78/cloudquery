package recipes

import (
	"github.com/xanzy/go-gitlab"
)

func Settings() []*Resource {
	resources := []*Resource{
		{
			Service:    "settings",
			SubService: "settings",
			PKColumns:  []string{"id"},
			Struct:     &gitlab.Settings{},
		},
	}

	return resources
}
