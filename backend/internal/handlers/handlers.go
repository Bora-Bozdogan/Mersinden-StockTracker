package handlers

import (
	"mersinden-stockapp/internal/models"
	"mersinden-stockapp/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service services.ServicesInterface
}

func NewHandler(service services.ServicesInterface) Handler {
	return Handler{service: service}
}

func (h Handler) GetItems(c *fiber.Ctx) error {
	//parse uid
	uid, err := h.parseUID(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Error trying to get valid UID",
		})
	}
	//call getitems of merchant
	items, err := h.service.GetItems(uid)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Server error trying to get products",
			"details": err.Error(),
		})
	}
	return c.JSON(items)
}

func (h Handler) CreateItem(c *fiber.Ctx) error {
	//parse uid
	uid, err := h.parseUID(c)

	//get merchant id from uid
	merchant, err := h.service.GetMerchantUID(uid)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Server error trying to get merchant",
			"details": err.Error(),
		})
	}

	merchantId := merchant.ID

	var productReq models.ProductRequest
	err = c.BodyParser(&productReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Couldn't parse product for creation",
			"details": err.Error(),
		})
	}
	err = h.service.CreateItem(merchantId, productReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Server error trying to create product",
			"details": err.Error(),
		})
	}
	return c.Status(201).SendString("Created Item")
}

func (h Handler) UpdateItem(c *fiber.Ctx) error {
	//get id from header
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Couldn't get id for update",
			"details": err.Error(),
		})
	}

	//parse product request
	var productReq models.ProductRequest
	err = c.BodyParser(&productReq)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Couldn't parse product for update",
			"details": err.Error(),
		})
	}
	err = h.service.UpdateItem(id, productReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Server error trying to update product",
			"details": err.Error(),
		})
	}
	return c.Status(200).SendString("Updated Item")
}

func (h Handler) DeleteItem(c *fiber.Ctx) error {
	//get id from header
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Couldn't get id for delete",
			"details": err.Error(),
		})
	}
	//parse int id
	err = h.service.DeleteItem(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Server error trying to delete product",
			"details": err.Error(),
		})
	}
	return c.Status(200).SendString("Deleted Item")
}

func (h Handler) GetMerchantInfo(c *fiber.Ctx) error {
	//get id from header
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Couldn't get id for update",
			"details": err.Error(),
		})
	}

	//get returning merchant
	merchant, err := h.service.GetMerchantID(id)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Couldn't parse product for creation",
			"details": err.Error(),
		})
	}
	//parse into merchant info object and send
	merchantInfo := &models.MerchantInfo{
		MerchantName: merchant.MerchantName,
		PhoneNumber:  merchant.PhoneNumber,
	}
	return c.JSON(merchantInfo)
}

func (h Handler) GetMerchantSelf(c *fiber.Ctx) error {
	//parse uid from header
	uid, err := h.parseUID(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Error trying to get valid UID",
		})
	}
	//get returning merchant
	merchant, err := h.service.GetMerchantUID(uid)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Couldn't parse product for creation",
			"details": err.Error(),
		})
	}
	//parse into merchant info object and send
	merchantInfo := &models.MerchantInfo{
		MerchantName: merchant.MerchantName,
		PhoneNumber:  merchant.PhoneNumber,
	}
	return c.JSON(merchantInfo)
}

func (h Handler) UpdateMerchantInfo(c *fiber.Ctx) error {
	//parse uid
	uid, err := h.parseUID(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Error trying to get valid UID",
		})
	}

	//parse merchantinfo
	var merchantInfo models.MerchantInfo
	err = c.BodyParser(&merchantInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Couldn't parse merchant for update",
			"details": err.Error(),
		})
	}

	//update merchant
	err = h.service.UpdateMerchant(uid, merchantInfo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Server error trying to update merchant",
			"details": err.Error(),
		})
	}
	return c.Status(200).SendString("Updated merchant")
}

func (h Handler) parseUID(c *fiber.Ctx) (string, error) {
	uidTemp := c.Locals("uid")
	uid, ok := uidTemp.(string)
	if !ok || uid == "" {
		//unauthorized, no uid, return error
		return "", c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Error trying to get valid UID",
		})
	}

	//uid parsed, query merchant db log new user in if needed
	merchant, err := h.service.GetMerchantUID(uid)

	if merchant == nil { //db healthy but merchant nonexistent
		merchantReq := models.MerchantRequest{
			UID:     uid,
			IsAdmin: false,
		}
		//merchant doesn't exist, add to db
		h.service.CreateMerchant(merchantReq)
	}

	if err != nil {
		return "", c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Server error trying to get merchant",
			"details": err.Error(),
		})
	}

	return uid, nil
}
