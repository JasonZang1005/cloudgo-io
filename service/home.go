package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	uuid "github.com/satori/go.uuid"
	"github.com/unrolled/render"
)

type User struct {
	UUID     uuid.UUID
	Username string
	Password string
}

func loginHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Method)
		if req.Method == "GET" {
			formatter.HTML(w, http.StatusOK, "form", "")
		} else {
			req.ParseForm()

			curuser := new(User)
			decoder := schema.NewDecoder()
			err := decoder.Decode(curuser, req.PostForm)
			if err != nil {
				panic(err)
			}

			fmt.Println(curuser)

			u1 := uuid.NewV4()
			curuser.UUID = u1
			formatter.HTML(w, http.StatusOK, "list", struct {
				Password string    `json:"password"`
				Username string    `json:"username"`
				UUID     uuid.UUID `json:"uuid"`
			}{Password: curuser.Password, Username: curuser.Username, UUID: curuser.UUID})
		}

	}
}
