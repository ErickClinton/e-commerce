package dto

type CreateWalletRequest struct {
	UserId  uint    `json:"user_id" binding:"required"`
	Balance float64 `json:"balance"`
}
