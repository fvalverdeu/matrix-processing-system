package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"qr-service-go/internal/models"
	"qr-service-go/internal/services"
)

type QRHandler struct {
	qrService *services.QRService
}

func NewQRHandler(qrService *services.QRService) *QRHandler {
	return &QRHandler{qrService: qrService}
}

func (h *QRHandler) Decompose(c *fiber.Ctx) error {
	var req models.QRRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
			Error: models.APIError{
				Code:    "INVALID_REQUEST",
				Message: "Request body must be valid JSON with a matrix field",
			},
		})
	}

	q, r, err := h.qrService.Decompose(req.Matrix)
	if err != nil {
		var validationErr *services.ValidationError
		if errors.As(err, &validationErr) {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
				Error: models.APIError{
					Code:    "INVALID_MATRIX",
					Message: validationErr.Message,
				},
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
			Error: models.APIError{
				Code:    "INTERNAL_ERROR",
				Message: "An unexpected error occurred while processing the matrix",
			},
		})
	}

	return c.JSON(models.QRResponse{
		Q: q,
		R: r,
	})
}
