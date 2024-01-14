package controllers

import (
	"golang_udemy/to-do-app/app/models"
	"log"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		generateHTML(w, nil, "layout", "public_navbar", "signup")
	} else if r.Method == "POST" {
		// form を解析
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		// 値をuserの各フィールドに受け取る
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		// user登録
		if err := user.CreateUser(); err != nil {
			log.Panicln(err)
		}

		// リダイレクト
		http.Redirect(w, r, "/", 302)
	}

}
