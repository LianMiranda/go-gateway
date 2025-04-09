package domain

import "errors"

var (
	ErrAccountNotFound = errors.New("account not found") //Erro de conta não encontrada
	ErrDuplicatedAPIKey = errors.New("api key already exists") //Erro de chave API duplicada
	ErrInvoiceNotFound = errors.New("invoice not found") //Erro de fatura não encontrada
	ErrUnauthorizedAccess = errors.New("unauthorized access") //Erro de acesso não autorizado
	ErrInvalidAmount = errors.New("invalid amount") //Erro de valor inválido
	ErrInvalidStatus = errors.New("invalid status") //Erro de status inválido
)