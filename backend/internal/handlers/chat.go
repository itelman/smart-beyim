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

type ResponeseChatGet struct {
	resp.Response
	Messages []openai.ChatCompletionMessage
}

func (h *Handler) chatGET(w http.ResponseWriter, r *http.Request) {
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

	responseOKChatGet(w, r, messages[1:])
}

func responseOKChatGet(w http.ResponseWriter, r *http.Request, messages []openai.ChatCompletionMessage) {
	render.JSON(w, r, ResponeseChatGet{
		Response: resp.OK(),
		Messages: messages,
	})
}

type ResponeseChatPost struct {
	resp.Response
	Answer string `json:"answer"`
}

type RequestPost struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}

func (h *Handler) chatPost(w http.ResponseWriter, r *http.Request) {
	var req RequestPost

	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		h.log.Error("failed to decode request", sl.Err(err))
		render.JSON(w, r, resp.Error("failed render request"))
		return
	}

	h.log.Info("request body decoded", slog.Any("request", req))

	userID, err := strconv.Atoi(req.UserID)
	if err != nil {
		h.log.Error("failed to convert UserID", sl.Err(err))
		render.JSON(w, r, resp.Error("failed convert user id"))
		return
	}

	answer, err := h.service.SendMessage(userID, req.Content)
	if err != nil {
		h.log.Error("failed to send message", sl.Err(err))
		render.JSON(w, r, resp.Error("failed send message ai"))
		return
	}

	responseOKChatPost(w, r, answer)
}

func responseOKChatPost(w http.ResponseWriter, r *http.Request, answer string) {
	render.JSON(w, r, ResponeseChatPost{
		Response: resp.OK(),
		Answer:   answer,
	})
}
