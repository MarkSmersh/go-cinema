package schemas

import (
	"fmt"
	"log"
	"strings"

	"github.com/MarkSmersh/go-cinema/database"
)

type Column struct {
	name       string
	columnType string
}

type Schema struct {
	name    string
	columns []Column
}

func (s Schema) Init() {
	var columns []string

	for i := 0; i < len(s.columns); i++ {
		columns = append(columns, fmt.Sprintf("%s %s", s.columns[i].name, s.columns[i].columnType))
	}

	query := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS public.%s (%s);",
		s.name,
		strings.Join(columns, ","),
	)

	rows, err := database.DB.Query(query)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v", rows)
}
