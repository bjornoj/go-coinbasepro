package coinbasepro

import (
	"fmt"
)

type WithdrawalCrypto struct {
	Currency      string `json:"currency"`
	Amount        string `json:"amount"`
	CryptoAddress string `json:"crypto_address"`
}

type WithdrawalCoinbase struct {
	Currency          string `json:"currency"`
	Amount            string `json:"amount"`
	CoinbaseAccountID string `json:"coinbase_account_id"`
}

type WithdrawalFiatRequest struct {
	Currency        string `json:"currency"`
	Amount          string `json:"amount"`
	PaymentMethodID string `json:"payment_method_id"`
}

type WithdrawalFiatResponse struct {
	ID string `json:"id"`
	WithdrawalFiatRequest
}

func (c *Client) CreateWithdrawalCrypto(newWithdrawalCrypto *WithdrawalCrypto) (WithdrawalCrypto, error) {
	var savedWithdrawal WithdrawalCrypto
	url := fmt.Sprintf("/withdrawals/crypto")
	_, err := c.Request("POST", url, newWithdrawalCrypto, &savedWithdrawal)
	return savedWithdrawal, err
}

func (c *Client) CreateWithdrawalCoinbase(newWithdrawalCoinbase *WithdrawalCoinbase) (WithdrawalCoinbase, error) {
	var savedWithdrawal WithdrawalCoinbase
	url := fmt.Sprintf("/withdrawals/coinbase-account")
	_, err := c.Request("POST", url, newWithdrawalCoinbase, &savedWithdrawal)
	return savedWithdrawal, err
}

func (c *Client) CreateWithdrawalFiat(newWithdrawalFiat *WithdrawalFiatRequest) (WithdrawalFiatResponse, error) {
	var savedWithdrawal WithdrawalFiatResponse
	url := fmt.Sprintf("/withdrawals/payment-method")
	_, err := c.Request("POST", url, newWithdrawalFiat, &savedWithdrawal)
	return savedWithdrawal, err
}
