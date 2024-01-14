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

func login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public_navbar", "login")
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		// sessionをブラウザに保存
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		// ログインに成功したらリダイレクト
		http.Redirect(w, r, "/", 302)
	} else {
		// passwordが一致しない場合
		http.Redirect(w, r, "/login", 302)
	}
}
