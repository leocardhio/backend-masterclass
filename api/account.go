package api

import (
	"database/sql"
	"net/http"

	db "github.com/leocardhio/masterclass/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=IDR JPY"`
}

func (s *Server) createAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createAccountRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		arg := db.CreateAccountParams{
			Owner:    req.Owner,
			Currency: req.Currency,
			Balance:  0,
		}

		account, err := s.store.CreateAccount(c, arg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		c.JSON(http.StatusOK, account)
	}
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getAccountRequest

		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		account, err := s.store.GetAccount(c, req.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, errorResponse(err))
				return
			}

			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		c.JSON(http.StatusOK, account)
	}
}

type getAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *Server) getAccounts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getAccountsRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}

		arg := db.GetAccountsParams{
			Offset: (req.PageID - 1) * req.PageSize,
			Limit:  req.PageSize,
		}

		accounts, err := s.store.GetAccounts(c, arg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		c.JSON(http.StatusOK, accounts)
	}
}
