package schemas

var User = Schema{
	name: "user",
	columns: []Column{
		{
			name:       "name",
			columnType: "VARCHAR(255)",
		},
	},
}
