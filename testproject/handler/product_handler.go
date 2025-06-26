package handler

import (
	"github.com/Oakhouse-Technology/go-to-oakhouse/dto/product"
	"github.com/Oakhouse-Technology/go-to-oakhouse/service"
	"github.com/Oakhouse-Technology/go-to-oakhouse/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) FindAll(c *fiber.Ctx) error {
	var dto product.GetProductDto
	if err := c.QueryParser(&dto); err != nil {
		return util.SendError(c, fiber.StatusBadRequest, "Invalid query parameters", err.Error())
	}
	
	dto.SetDefaults()
	
	products, total, err := h.service.FindWithPagination(c.Context(), &dto)
	if err != nil {
		return util.SendError(c, fiber.StatusInternalServerError, "Failed to fetch products", err.Error())
	}
	
	pagination := util.CalculatePagination(dto.Page, dto.PageSize, total)
	return util.SendPaginatedSuccess(c, "Products retrieved successfully", products, pagination)
}

func (h *ProductHandler) FindByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return util.SendError(c, fiber.StatusBadRequest, "Invalid ID format", err.Error())
	}
	
	product, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		return util.SendError(c, fiber.StatusNotFound, "Product not found", err.Error())
	}
	
	return util.SendSuccess(c, "Product retrieved successfully", product)
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var dto product.CreateProductDto
	if err := c.BodyParser(&dto); err != nil {
		return util.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}
	
	product, err := h.service.Create(c.Context(), &dto)
	if err != nil {
		return util.SendError(c, fiber.StatusInternalServerError, "Failed to create product", err.Error())
	}
	
	return util.SendSuccess(c, "Product created successfully", product)
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return util.SendError(c, fiber.StatusBadRequest, "Invalid ID format", err.Error())
	}
	
	var dto product.UpdateProductDto
	if err := c.BodyParser(&dto); err != nil {
		return util.SendError(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}
	
	product, err := h.service.Update(c.Context(), id, &dto)
	if err != nil {
		return util.SendError(c, fiber.StatusInternalServerError, "Failed to update product", err.Error())
	}
	
	return util.SendSuccess(c, "Product updated successfully", product)
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return util.SendError(c, fiber.StatusBadRequest, "Invalid ID format", err.Error())
	}
	
	if err := h.service.Delete(c.Context(), id); err != nil {
		return util.SendError(c, fiber.StatusInternalServerError, "Failed to delete product", err.Error())
	}
	
	return util.SendSuccess(c, "Product deleted successfully", nil)
}
