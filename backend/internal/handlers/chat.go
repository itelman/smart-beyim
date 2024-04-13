package handlers

import (
	"log/slog"
	"net/http"
	resp "smart-beyim/lib/api/response"
	"smart-beyim/lib/logger/sl"
	"strconv"

	"github.com/go-chi/render"
	"github.com/sashabaranov/go-openai"
)

type Response struct {
	resp.Response
	Messages []openai.ChatCompletionMessage
}

type RequestGet struct {
	UserID string `json:"user_id"`
}

func (h *Handler) chatGET(w http.ResponseWriter, r *http.Request) {
	h.log.Debug("chat get is here")
    userID := r.URL.Query().Get("user_id")
    
    userIDInt, err := strconv.Atoi(userID)
    if err != nil {
        h.log.Error("failed to convert User id", sl.Err(err))
        render.JSON(w, r, resp.Error("bad request"))
        return
    }

    messages, err := h.service.GetMessages(userIDInt)
    if err != nil {
        h.log.Error("failed to get Messages", sl.Err(err))
        render.JSON(w, r, resp.Error("internal server error"))
        return
    }

    h.log.Debug("just messages", slog.Any("messages", messages))

    responseOK(w, r, messages)
}


func responseOK(w http.ResponseWriter, r *http.Request, messages []openai.ChatCompletionMessage) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		Messages: messages,
	})
}
