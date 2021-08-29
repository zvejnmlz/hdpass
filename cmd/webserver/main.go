package main

import (
	"flag"
	"fmt"
	"hdpass/src"
	"html/template"
	"net/http"
	"os"
	"path"
	"strconv"
)

type Form struct {
	PasswordIndex string
	LoginIndex    string
	Login         string
	Password      string
	Domain        string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	seed := flag.String("seed", "", "HD Seed")
	port := flag.String("port", ":8008", "Server port")
	flag.Parse()

	defaultLoginRule, err := hdpass.NewRule(8, map[string]uint8{
		hdpass.RuleLowerAlpha: 8,
	})
	checkErr(err)

	defaultPasswordRule, err := hdpass.NewRule(10, map[string]uint8{
		hdpass.RuleLowerAlpha: 0,
		hdpass.RuleUpperAlpha: 1,
		hdpass.RuleDigit:      1,
		hdpass.RuleSpecial1:   1,
	})
	checkErr(err)

	defaultSeedRule, err := hdpass.NewRule(20, map[string]uint8{
		hdpass.RuleUpperAlpha: 1,
		hdpass.RuleLowerAlpha: 10,
		hdpass.RuleDigit:      8,
		hdpass.RuleSpecial1:   1,
	})
	checkErr(err)

	s, err := hdpass.RestoreFromSeed(*seed, defaultSeedRule)
	checkErr(err)

	p, err := os.Executable()
	filepath := path.Join(path.Dir(p), "/forms.html")
	tmpl := template.Must(template.ParseFiles(filepath))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			err := tmpl.Execute(w, Form{
				Login:         "Empty",
				Password:      "Empty",
				PasswordIndex: "0",
				LoginIndex:    "0",
			})

			if err != nil {
				fmt.Printf("Error:%s", err)
			}

			return
		}

		domain := r.FormValue("domain")
		a := s.NewDomain(domain, defaultLoginRule, defaultPasswordRule)

		loginIndex, err := strconv.Atoi(r.FormValue("loginIndex"))
		if err != nil {
			loginIndex = 0
		}

		passIndex, err := strconv.Atoi(r.FormValue("passwordIndex"))
		if err != nil {
			passIndex = 0
		}

		login := a.SelectLogin(loginIndex)
		password := a.SelectPassword(passIndex)

		err = tmpl.Execute(w, Form{
			Login:         login,
			Password:      password,
			Domain:        domain,
			PasswordIndex: strconv.Itoa(passIndex),
			LoginIndex:    strconv.Itoa(loginIndex),
		})

		if err != nil {
			fmt.Printf("Error:%s", err)
		}
	})

	err = http.ListenAndServe(*port, nil)

	if err != nil {
		fmt.Printf("Error:%s", err)
	}
}
