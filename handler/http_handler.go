package handler

import (
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
	wrapper := request.ContextWrapper(c)

	req := new(dto.SignUpReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [SignUp] bad request: %v", err), "")
	}

	res, err := h.d.Service.SignUp(*req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [SignUp] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
