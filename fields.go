package toforms

type field struct {
	Label       string
	Name        string
	Type        string
	Placeholder string
}

func fields(strct interface{}) field {
	return field{}
}

// func HTML(strct interface{}, tpl *template.Template) template.HTML {
// }
