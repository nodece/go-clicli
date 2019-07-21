package handler

import (
	auth "github.com/nilslice/jwt"
	"net/http"
	"encoding/json"
	"io"
	"github.com/julienschmidt/httprouter"
)

func Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := r.Cookie("token")
	if err != nil {
		resStr, _ := json.Marshal(struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}{Code: 401, Msg: "鉴权失败，请重新登陆"})

		io.WriteString(w, string(resStr))
		return
	} else {
		token := t.Value
		resStr, _ := json.Marshal(struct {
			Token string `json:"token"`
		}{Token: token})
		io.WriteString(w, string(resStr))
	}
}

func AuthToken(w http.ResponseWriter, r *http.Request, level int) {
	token := r.Header.Get("token")
	if auth.Passes(token) {
		s := auth.GetClaims(token)
		if int(s["level"].(float64)) < level {
			io.WriteString(w, string("权限不足"))
			return
		}
	} else {
		io.WriteString(w, string("token无效或过期"))
		return
	}

}
