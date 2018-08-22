package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/transform", func(w http.ResponseWriter, r *http.Request) {
		if err := transform(w, r); err != nil {
			log.Printf("%s: %s", r.RemoteAddr, err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		}
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func transform(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("/transform request must be POST method")
	}

	if err := r.ParseMultipartForm(1); err != nil {
		return err
	}

	file, fr, err := r.FormFile("file")
	if err != nil {
		return err
	}

	sep := sepToRune(r.FormValue("sep"))
	if sep == 0 {
		return fmt.Errorf("Invalid separator")
	}

	cw, err := convert(file, sep)
	if err != nil {
		return fmt.Errorf("Can't convert %s", fr.Filename)
	}

	w.Header().Set("Content-Disposition", "attachment; filename=\"transform.csv\"")
	w.Header().Set("Content-type", "text/csv")
	_, err = io.Copy(w, cw)

	return err
}

func sepToRune(s string) rune {
	options := map[string]rune{
		";":   ';',
		"|":   '|',
		"tab": '\t',
	}
	return options[s]
}

func convert(file io.Reader, sep rune) (io.Reader, error) {
	var w bytes.Buffer
	csvR := csv.NewReader(file)
	csvW := csv.NewWriter(&w)

	csvW.Comma = sep

	for {
		record, err := csvR.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return &w, err
		}

		csvW.Write(record)

	}
	csvW.Flush()
	if err := csvW.Error(); err != nil {
		return &w, err
	}

	return &w, nil
}
