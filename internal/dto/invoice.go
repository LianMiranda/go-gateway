package dto

import (
	"time"

	"github.com/lianmiranda/imersaofullcycle/go-gateway/internal/domain"
)

const (
	StatusPending  = string(domain.StatusPending)
	StatusApproved = string(domain.StatusApproved)
	StatusRejected = string(domain.StatusRejected)
)

type CreateInvoiceInput struct {
	APIKey         string
	Amount         float64 `json:"amount"`
	Description    string  `json:"description"`
	PaymentType    string  `json:"payment_type"`
	CardNumber     string  `json:"card_number"`
	CVV            string  `json:"cvv"`
	ExpiryMonth    int     `json:"expiry_month"`
	ExpiryYear     int     `json:"expiry_year"`
	CardHolderName string  `json:"card_holder_name"`
}

type InvoiceOutput struct {
	ID             string
	AccountID      string
	Amount         float64
	Status         string
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func ToInvoice(input CreateInvoiceInput, accountID string) (*domain.Invoice, error) {
	card := domain.CreditCard{
		Number:         input.CardNumber,
		CVV:            input.CVV,
		ExpiryMonth:    input.ExpiryMonth,
		ExpiryYear:     input.ExpiryYear,
		CardHolderName: input.CardHolderName,
	}

	return domain.NewInvoice(accountID, input.Amount, input.Description, input.PaymentType, card)

}

func FromInvoice(invoice *domain.Invoice) *InvoiceOutput {
	return &InvoiceOutput{
		ID: invoice.ID,
		AccountID: invoice.AccountID,
		Amount: invoice.Amount,
		Status: string(invoice.Status),
		Description: invoice.Description,
		PaymentType: invoice.PaymentType,
		CardLastDigits: invoice.CardLastDigits,
		CreatedAt: invoice.CreatedAt,
		UpdatedAt: invoice.UpdatedAt,
	}
}
