package cli

import (
	"fmt"
	"gca/service/hoge/domain/entity"

	"github.com/spf13/cobra"
)

// ListResources ...
func ListResources(cmd *cobra.Command, args []string) {
	resources := []entity.Resource{
		entity.Resource{
			ID:          "1",
			Description: "aaa",
			Data:        "aaaaaa",
		},
		entity.Resource{
			ID:          "2",
			Description: "bbb",
			Data:        "bbbbbb",
		},
		entity.Resource{
			ID:          "3",
			Description: "ccc",
			Data:        "cccccc",
		},
	}

	fmt.Println(resources)
}
