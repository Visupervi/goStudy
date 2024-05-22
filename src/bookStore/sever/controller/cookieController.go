package controller

import (
	"fmt"
	"net/http"
	"net/url"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "username",
		Value:    url.QueryEscape("学院君"),
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "website",
		Value:    "https://xueyuanjun.com",
		HttpOnly: true,
	}
	w.Header().Add("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
	fmt.Fprintln(w, "通过 HTTP 响应头发送 Cookie 信息")
}
