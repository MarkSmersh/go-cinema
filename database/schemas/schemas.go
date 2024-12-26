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

	fmt.Printf("%+v\n", rows)

	columns, err = rows.Columns()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", columns)
}

func (s Schema) FindById(id int16) {
	query := fmt.Sprintf("SELECT 1 FROM public.%s WHERE name = %d;", s.name, id)

	rows, err := database.DB.Query(query)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", rows)
}

// I have no idea how to implement
// generic type for values
// (depends on the values of s)

func (s Schema) Create(col []string, val []string) {
	var quoteValues []string

	for i := 0; i < len(val); i++ {
		quoteValues = append(quoteValues, fmt.Sprintf("'%s'", val[i]))
	}

	query := fmt.Sprintf(
		"INSERT INTO public.%s (%s) VALUES (%s);",
		s.name,
		strings.Join(col, ","),
		strings.Join(quoteValues, ","),
	)

	fmt.Println(query)

	rows, err := database.DB.Query(query)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", rows)
}

func (s Schema) Find(col string, val string) {
	// I'm curious how does it work with int

	query := fmt.Sprintf("SELECT * FROM public.%s WHERE %s = '%s'", s.name, col, val)

	database.DB.Query(query)
}

func (s Schema) Query(statement string, data string, args ...any) {
	query := fmt.Sprintf(
		"%s FROM public.%s %s",
		statement,
		s.name,
		fmt.Sprintf(data, args...),
	)

	fmt.Println(query)

	rows, err := database.DB.Query(query)

	if err != nil {
		log.Fatalln()
	}

	fmt.Printf("%+v\n", rows)
}
