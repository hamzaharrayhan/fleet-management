package controller

import (
	"fleet_management/internal/dto"
	"fleet_management/internal/service"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type LocationController struct {
	service service.LocationService
}

func NewLocationController(service service.LocationService) *LocationController {
	return &LocationController{service: service}
}

func (c *LocationController) GetLatestLocation(ctx *fiber.Ctx) error {
	vehicleID := ctx.Params("vehicle_id")

	loc, err := c.service.GetLatestLocation(ctx.Context(), vehicleID)
	if err != nil {
		logrus.WithError(err).Error("GetLatestLocation failed")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get latest location"})
	}

	resp := dto.LocationResponse{
		VehicleID: loc.VehicleID,
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
		Timestamp: loc.Timestamp,
	}

	return ctx.JSON(resp)
}

func (c *LocationController) GetHistory(ctx *fiber.Ctx) error {
	vehicleID := ctx.Params("vehicle_id")

	startStr := ctx.Query("start")
	endStr := ctx.Query("end")

	start, err := strconv.ParseInt(startStr, 10, 64)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start timestamp"})
	}

	end, err := strconv.ParseInt(endStr, 10, 64)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end timestamp"})
	}

	history, err := c.service.GetHistory(ctx.Context(), vehicleID, start, end)
	if err != nil {
		logrus.WithError(err).Error("GetHistory failed")
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get location history"})
	}

	resp := make([]dto.LocationResponse, 0, len(history))
	for _, loc := range history {
		resp = append(resp, dto.LocationResponse{
			VehicleID: loc.VehicleID,
			Latitude:  loc.Latitude,
			Longitude: loc.Longitude,
			Timestamp: loc.Timestamp,
		})
	}

	return ctx.JSON(resp)
}
