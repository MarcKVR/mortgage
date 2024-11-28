package handler

import (
	"github.com/MarcKVR/mortgage/service"
	"github.com/gofiber/fiber/v2"
)

type (
	PaymentHandler struct {
		service service.PaymentService
	}

	CreateReq struct {
		MonthlyPayment  float64 `json:"monthly_payment"`
		DamageInsurance float64 `json:"damage_insurance"`
		LifeInsurance   float64 `json:"life_insurance"`
		Interests       float64 `json:"interests"`
		Capital         float64 `json:"capital"`
		Rate            float32 `json:"rate"`
		PaymentNumber   int32   `json:"payment_number"`
		DatePayment     string  `json:"date_payment"`
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
		// Meta   *meta.Meta  `json:"meta,omitempty"`
	}
)

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}

func (h *PaymentHandler) Create(c *fiber.Ctx) error {
	payment := new(CreateReq)

	if err := c.BodyParser(&payment); err != nil {
		return err
	}

	paymentCreated, err := h.service.Create(
		payment.MonthlyPayment,
		payment.DamageInsurance,
		payment.LifeInsurance,
		payment.Interests,
		payment.Capital,
		payment.Rate,
		payment.PaymentNumber,
		payment.DatePayment,
	)

	if err != nil {
		c.Status(422)
		return c.JSON(Response{Status: 422, Err: err.Error()})
	}

	return c.JSON(Response{Status: 200, Data: paymentCreated})
}

func (h *PaymentHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	payment, err := h.service.Get(id)
	if err != nil {
		c.Status(404)
		return c.JSON(Response{Status: 404, Err: err.Error()})
	}
	return c.JSON(Response{Status: 200, Data: payment})
}
