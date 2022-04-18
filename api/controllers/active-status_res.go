package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// swagger:route GET/games/{id}/active-status Games getGame
// Get game using game id
// Responses:
//  200: GetActiveStatusRes
//  400: StatusCode
//  404: StatusCode
//  500: StatusCode
// GetActiveStatusRes are get next round,player and dart data for verify with frontend URL
func GetActiveStatusRes(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	gameId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)

		return ctx.Status(400).JSON(types.StatusCode{
			StatusCode: 400,
			Message:    "Bad Request",
		})
	}
	activeRes := types.ActiveStatus{}
	players := types.NextTurn{}

	activejson, err := models.GetActiveStatusRes(db, gameId, activeRes, players)
	if err != nil {

		fmt.Println(err)
		return ctx.Status(404).JSON(types.StatusCode{
			StatusCode: 404,
			Message:    "No Content",
		})
	}
	return ctx.Status(200).JSON(activejson)
}
