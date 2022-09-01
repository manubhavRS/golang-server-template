package handler

import (
	"github.com/tejashwikalptaru/tutorial/database/helper"
	"github.com/tejashwikalptaru/tutorial/utils"
	"net/http"
)

func isErr(err error, typeErr string) bool {
	return err.Error() == "pq: duplicate key value violates unique constraint "+typeErr
}

func Greet(writer http.ResponseWriter, request *http.Request) {

	userID, err := helper.CreateUser("test", "test@test.com")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	user, userErr := helper.GetUser(userID)
	if userErr != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	jsonData, err := utils.DecodeToJson(user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = writer.Write(jsonData)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
