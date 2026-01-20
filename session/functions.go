package session

import "github.com/gofiber/fiber/v2"

const userIDKey = "user_id"

func CreateSession(ctx *fiber.Ctx, userID uint) error {
	return Set(ctx, userIDKey, userID)
}

func DestroySession(ctx *fiber.Ctx) error {
	return Delete(ctx, userIDKey)
}

func GetSessionUserID(ctx *fiber.Ctx) (uint, error) {
	value, err := Get(ctx, userIDKey)
	if err != nil || value == nil {
		return 0, err
	}

	userID, ok := value.(uint)
	if !ok {
		return 0, nil
	}

	return userID, nil
}
