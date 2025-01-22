package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/project-inari/core-user-server/dto"
	"github.com/project-inari/core-user-server/pkg/request"
	"github.com/project-inari/core-user-server/pkg/response"
)

type httpHandler struct {
	d Dependencies
}

func newHTTPHandler(d Dependencies) *httpHandler {
	return &httpHandler{
		d: d,
	}
}

func (h *httpHandler) SignUp(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(dto.SignUpReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [SignUp] bad request: %v", err), "")
	}

	res, err := h.d.Service.SignUp(ctx, *req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [SignUp] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *httpHandler) Inquiry(c echo.Context) error {
	username := c.Param("username")
	if username == "" {
		return response.ErrorResponse(c, http.StatusBadRequest, "error - [Inquiry] bad request: username is required", "")
	}

	res, err := h.d.Service.Inquiry(username)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [Inquiry] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
