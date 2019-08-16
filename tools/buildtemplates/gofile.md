package adacore

{{range $index, $val := .Templates}}
const template{{$val.Name}} = "{{$val.Str}}"
{{end}}