package helper

import (
	"html/template"
	"net/http"
	"bytes"
	"log"
)

const tplPath = "templates/"

func Templating(w http.ResponseWriter, tmplName string, layout string, args ...any) {
	logger, logFile, err := CreateLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer CloseLogger(logFile)

	buf := &bytes.Buffer{}

	tmpl, err := template.New(layout).Funcs(template.FuncMap{"N": N}).ParseFiles(tplPath+tmplName+".html", tplPath+layout+".html")
	if err != nil {
		logger.Printf("helper.Templating; template.New()| %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(args) == 0 {
		err = tmpl.Execute(buf, nil)
	} else {
		err = tmpl.Execute(buf, args[0])
	}
	if err != nil {
		logger.Printf("helper.Templating; tmpl.Execute(len(args)=%v)| %v\n", len(args), err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}