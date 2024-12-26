package init

import (
	"fmt"

	"github.com/MarkSmersh/go-cinema/database/schemas"
)

func init() {
	schemas.User.Init()

	schemas.User.Create(
		[]string{"name"},
		[]string{"test"},
	)

	fmt.Println("Nice!")
}
