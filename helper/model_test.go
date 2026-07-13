package helper

import (
	"golang-banking-api/model/domain"
	"testing"

	"github.com/shopspring/decimal"
)

func TestToUserResponse_DoesNotExposeSensitiveFields(t *testing.T) {
	user := domain.User{
		Id:           1,
		Email:        "test@example.com",
		Password:     "hashed_secret",
		Name:         "Test User",
		Role:         "user",
		RefreshToken: "refresh_token_value",
	}

	resp := ToUserResponse(user)

	if resp.Id != 1 {
		t.Errorf("expected Id 1, got %d", resp.Id)
	}
	if resp.Email != "test@example.com" {
		t.Errorf("expected Email test@example.com, got %s", resp.Email)
	}
	if resp.Name != "Test User" {
		t.Errorf("expected Name Test User, got %s", resp.Name)
	}
	if resp.Role != "user" {
		t.Errorf("expected Role user, got %s", resp.Role)
	}
}

func TestToAccountResponse(t *testing.T) {
	account := domain.Account{
		Id:          1,
		UserId:      2,
		AccountBank: "BCA",
		Balance:     decimal.RequireFromString("1000.50"),
		AccountType: "savings",
	}

	resp := ToAccountResponse(account)

	if resp.Id != 1 {
		t.Errorf("expected Id 1, got %d", resp.Id)
	}
	if resp.UserId != 2 {
		t.Errorf("expected UserId 2, got %d", resp.UserId)
	}
	if resp.AccountBank != "BCA" {
		t.Errorf("expected AccountBank BCA, got %s", resp.AccountBank)
	}
	if resp.Balance.String() != "1000.5" {
		t.Errorf("expected Balance 1000.5, got %s", resp.Balance)
	}
	if resp.AccountType != "savings" {
		t.Errorf("expected AccountType savings, got %s", resp.AccountType)
	}
}

func TestToTransactionResponse(t *testing.T) {
	transaction := domain.Transaction{
		Id:            1,
		FromAccountId: 2,
		ToAccountId:   3,
		Amount:        decimal.RequireFromString("500.00"),
		Type:          "transfer",
		Description:   "Payment",
	}

	resp := ToTransactionResponse(transaction)

	if resp.Id != 1 {
		t.Errorf("expected Id 1, got %d", resp.Id)
	}
	if resp.FromAccountId != 2 {
		t.Errorf("expected FromAccountId 2, got %d", resp.FromAccountId)
	}
	if resp.ToAccountId != 3 {
		t.Errorf("expected ToAccountId 3, got %d", resp.ToAccountId)
	}
	if resp.Amount.String() != "500" {
		t.Errorf("expected Amount 500, got %s", resp.Amount)
	}
	if resp.Type != "transfer" {
		t.Errorf("expected Type transfer, got %s", resp.Type)
	}
	if resp.Description != "Payment" {
		t.Errorf("expected Description Payment, got %s", resp.Description)
	}
}

func TestToUserResponses(t *testing.T) {
	users := []domain.User{
		{Id: 1, Email: "a@test.com", Name: "A", Role: "user"},
		{Id: 2, Email: "b@test.com", Name: "B", Role: "admin"},
	}

	resps := ToUserResponses(users)

	if len(resps) != 2 {
		t.Fatalf("expected 2 responses, got %d", len(resps))
	}
	if resps[0].Email != "a@test.com" {
		t.Errorf("expected first email a@test.com, got %s", resps[0].Email)
	}
	if resps[1].Role != "admin" {
		t.Errorf("expected second role admin, got %s", resps[1].Role)
	}
}

func TestToAccountResponses(t *testing.T) {
	accounts := []domain.Account{
		{Id: 1, AccountBank: "A", Balance: decimal.RequireFromString("100"), AccountType: "savings"},
		{Id: 2, AccountBank: "B", Balance: decimal.RequireFromString("200"), AccountType: "checking"},
	}

	resps := ToAccountResponses(accounts)

	if len(resps) != 2 {
		t.Fatalf("expected 2 responses, got %d", len(resps))
	}
	if resps[0].Balance.String() != "100" {
		t.Errorf("expected first balance 100, got %s", resps[0].Balance)
	}
}

func TestToTransactionResponses(t *testing.T) {
	transactions := []domain.Transaction{
		{Id: 1, Amount: decimal.RequireFromString("100"), Type: "deposit"},
		{Id: 2, Amount: decimal.RequireFromString("200"), Type: "withdrawal"},
	}

	resps := ToTransactionResponses(transactions)

	if len(resps) != 2 {
		t.Fatalf("expected 2 responses, got %d", len(resps))
	}
	if resps[0].Type != "deposit" {
		t.Errorf("expected first type deposit, got %s", resps[0].Type)
	}
}
