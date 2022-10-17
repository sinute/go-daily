package main

import (
	"fmt"
	"html/template"
	"strings"
)

var done = make(chan struct{}, 1)

func main()  {
	//go func() {
	//	done <- struct{}{}
	//}()
	//
	//select {
	//case c := <-done: // <- struct{}{}:
	//	fmt.Println(c)
	//default:
	//	fmt.Println("default")
	//}
	//fmt.Println("aaa")
	//return

	//a:= "你好"
	//fmt.Println(len([]byte(a)))
	//b := []rune(a)
	//fmt.Println(string(b[0:1]))
	//fmt.Println(time.Date(2021, 2, 31, 0, 0, 0, 0, time.Local).Format("200601"))
	tmpl, err := template.New("sql_executor").Funcs(template.FuncMap{
		"join": strings.Join,
	}).Parse("1{{.}}1")
	if err != nil {
		fmt.Println(err)
		return
	}
	//tmpl, err := template.New("sql_executor").Funcs(template.FuncMap{
	//	"join": strings.Join,
	//}).Parse("{{if and .pipeline .pip}} T1 {{end}}{{join .a \",\" }}")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	content := strings.Builder{}
	//if err = tmpl.Execute(&content, map[string]interface{}{
	//	"a": []string{"1", "2"},
	//	"pipeline": "a",
	//	"pip": false,
	//}); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	if err = tmpl.Execute(&content, fmt.Errorf("asdf")); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(content.String())
}
