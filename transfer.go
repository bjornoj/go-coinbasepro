package coinbasepro

import (
	"errors"
	"fmt"
)

type Transfer struct {
	Type              string `json:"type"`
	Amount            string `json:"amount"`
	CoinbaseAccountID string `json:"coinbase_account_id,string"`
}

type ListTransferRequest struct {
	// Type can be deposit or withdraw. Required
	Type string `json:"type"`
	// (Optional) Limit list of transfers to this profile_id. By default, it retrieves transfers across all of the user's profiles
	ProfileID string `json:"profile_id"`

	Pagination PaginationParams
}

func (c *Client) CreateTransfer(newTransfer *Transfer) (Transfer, error) {
	var savedTransfer Transfer

	url := fmt.Sprintf("/transfers")
	_, err := c.Request("POST", url, newTransfer, &savedTransfer)
	return savedTransfer, err
}

func (c *Client) ListTransfers(request ListTransferRequest) ([]Transfer, error) {
	var savedTransfers []Transfer

	if request.Type == "" {
		return nil, errors.New("type is required")
	}

	url := fmt.Sprintf("/transfers")
	_, err := c.Request("GET", url, request, &savedTransfers)
	return savedTransfers, err
}

func (c *Client) ListWithdrawals(profileID string, pagination *PaginationParams) ([]Transfer, error) {

	if pagination == nil {
		pagination = &PaginationParams{}
	}

	var savedTransfers []Transfer

	request := ListTransferRequest{
		Type:             "withdraw",
		ProfileID:        profileID,
		Pagination: *pagination,
	}

	url := fmt.Sprintf("/transfers")
	_, err := c.Request("GET", url, request, &savedTransfers)
	return savedTransfers, err
}

func (c *Client) ListDeposits(profileID string, pagination *PaginationParams) ([]Transfer, error) {

	if pagination == nil {
		pagination = &PaginationParams{}
	}

	var savedTransfers []Transfer
	request := ListTransferRequest{
		Type:             "deposit",
		ProfileID:        profileID,
		Pagination: *pagination,
	}

	url := fmt.Sprintf("/transfers")
	_, err := c.Request("GET", url, request, &savedTransfers)
	return savedTransfers, err
}
