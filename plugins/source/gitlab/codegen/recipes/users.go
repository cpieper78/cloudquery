package recipes

import (
	"github.com/xanzy/go-gitlab"
)

func Users() []*Resource {
	resources := []*Resource{
		{
			Service:    "users",
			SubService: "users",
			PKColumns:  []string{"id"},
			Struct:     &gitlab.User{},
		},
	}

	return resources
}