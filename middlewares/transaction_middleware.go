package middlewares

import (
	"github.com/DiansSopandi/goride_be/pkg/db"
	"github.com/gofiber/fiber/v2"
)

const UserServiceCtxKey = "userServiceWithTx"
const TxContextKey = "tx"

func WithTransaction(handler fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		dbConn := db.InitDatabase()
		tx, err := dbConn.Begin() // start transaction commit rollback
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to start transaction")
		}

		// repo, err := repository.NewUserRepository(true) // true = start TX
		// if err != nil {
		// 	return fiber.NewError(fiber.StatusInternalServerError, "Failed to start transaction")
		// }

		// Inject service dengan repo yg punya TX ke context
		// userService := service.NewUserService(repo)
		// c.Locals(UserServiceCtxKey, userService)
		// Simpan TX ke dalam context
		c.Locals(TxContextKey, tx)

		// Jalankan handler utama
		err = handler(c)

		if err != nil {
			// Rollback jika ada error
			db.RollbackOnError(tx, err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "failed to commit transaction", // fmt.Sprintf("Failed to commit transaction: %v", err),
				"status":  fiber.StatusInternalServerError,
				"success": false,
				"error":   err.Error(),
				"data":    nil,
			})
		}

		// Commit transaksi jika sukses
		if tx != nil {
			if err := tx.Commit(); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "failed to commit transaction", // fmt.Sprintf("Failed to commit transaction: %v", err),
					"status":  fiber.StatusInternalServerError,
					"success": false,
					"error":   err.Error(),
					"data":    nil,
				})
			}
		}

		return nil
	}
}
