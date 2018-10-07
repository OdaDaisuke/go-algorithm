package net_http

import (
	"net/http"
	"fmt"
)

func Start() {
	pattern1()
	pattern2()
}

/*--------------------
 * Pattern 1
 ----------------------*/

/*
 * Define middleware
 */
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 何らかの前処理(認証などここに入れると良さそう)

		next.ServeHTTP(w, r)

		// 何らかの後処理(defer func()で実現できない処理をここに書くと良さそう?)
	})
}

/*
 * Sample AccountHandler
 */
type AccountHandler struct {}

func NewAccountHandler() *AccountHandler {
	return &AccountHandler{}
}

func (ah *AccountHandler) SampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK\n"))
}

/*
 * Main func
 */
func pattern1() {
	accountHandler := NewAccountHandler()

	http.Handle("/", middleware(http.HandlerFunc(accountHandler.SampleHandler)))
	http.ListenAndServe(":8080", nil)

}

/*--------------------
 * Pattern 2
 ----------------------*/
type HandlerProvider struct {}

/*--------------------------------
 * ダックタイピングで勝手にcallされる
 ---------------------------------*/
func (hp HandlerProvider) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* 何らかの処理(認証などここに入れると良さそう) */
	fmt.Println("Hello world!")

	/* Define endpoints */
	switch r.URL.Path {
	case "/account":
		fmt.Println("create account")
		break
	}

	/* 何らかの後処理(defer func()で実現できない処理をここに書くと良さそう?) */
	fmt.Println("ByeBye!")

}

func pattern2() {
	handlerProvider := &HandlerProvider{}

	http.Handle("/", handlerProvider)
	http.ListenAndServe(":8081", nil)

}
