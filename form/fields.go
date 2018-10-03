package form

import (
	"reflect"
)

// func HTML(strct interface{}, tpl *template.Template) template.HTML {
// 	return template.HTML("")
// }

func fields(strct interface{}) []field {
	rv := reflect.ValueOf(strct)
	t := rv.Type()
	var ret []field
	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)
		rvf := rv.Field(i)
		f := field{
			Label:       tf.Name,
			Name:        tf.Name,
			Type:        "text",
			Placeholder: tf.Name,
			Value:       rvf.Interface(),
		}
		ret = append(ret, f)
	}
	return ret
}

type field struct {
	Label       string
	Name        string
	Type        string
	Placeholder string
	Value       interface{}
}