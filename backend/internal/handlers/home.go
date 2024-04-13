package handlers

import (
	"net/http"
	resp "smart-beyim/lib/api/response"
	"smart-beyim/lib/logger/sl"
	"smart-beyim/models"
	"strconv"

	"github.com/go-chi/render"
)

type ResponeseHomeGet struct {
	resp.Response
	IeltsList []models.Ielts
}

func (h *Handler) homeGet(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("user_id")

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		h.log.Error("failed to convert User id", sl.Err(err))
		render.JSON(w, r, resp.Error("bad request"))
		return
	}

	ieltsList, err := h.service.GetIelts(userIDInt)

	if err != nil {
		h.log.Error("failed to get Messages", sl.Err(err))
		render.JSON(w, r, resp.Error("internal server error"))
		return
	}
	responseOKHomeGet(w, r, ieltsList)
}

func responseOKHomeGet(w http.ResponseWriter, r *http.Request, ieltsList []models.Ielts) {
	render.JSON(w, r, ResponeseHomeGet{
		Response:  resp.OK(),
		IeltsList: ieltsList,
	})
}
