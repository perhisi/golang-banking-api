package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"golang-banking-api/model/web"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/shopspring/decimal"
)

type mockTransactionService struct {
	createResp     web.TransactionResponse
	updateResp     web.TransactionResponse
	deleteCalled   bool
	findByIdResp   web.TransactionResponse
	findAllResp    []web.TransactionResponse
	findByAcctResp []web.TransactionResponse
}

func (m *mockTransactionService) Create(ctx context.Context, request web.TransactionCreateRequest) web.TransactionResponse {
	return m.createResp
}
func (m *mockTransactionService) Update(ctx context.Context, transactionId int, request web.TransactionUpdateRequest) web.TransactionResponse {
	m.updateResp.Id = transactionId
	return m.updateResp
}
func (m *mockTransactionService) Delete(ctx context.Context, transactionId int) {
	m.deleteCalled = true
}
func (m *mockTransactionService) FindById(ctx context.Context, transactionId int) web.TransactionResponse {
	m.findByIdResp.Id = transactionId
	return m.findByIdResp
}
func (m *mockTransactionService) FindAll(ctx context.Context) []web.TransactionResponse {
	return m.findAllResp
}
func (m *mockTransactionService) FindByAccountId(ctx context.Context, accountId int) []web.TransactionResponse {
	return m.findByAcctResp
}
func (m *mockTransactionService) Deposit(ctx context.Context, request web.TransactionDepositRequest) web.TransactionResponse {
	return m.createResp
}
func (m *mockTransactionService) Withdraw(ctx context.Context, request web.TransactionWithdrawRequest) web.TransactionResponse {
	return m.createResp
}
func (m *mockTransactionService) Transfer(ctx context.Context, request web.TransactionTransferRequest) web.TransactionResponse {
	return m.createResp
}
func (m *mockTransactionService) FindMyTransactions(ctx context.Context, userId int) []web.TransactionResponse {
	return m.findAllResp
}
func (m *mockTransactionService) FindMyTransactionById(ctx context.Context, userId int, transactionId int) web.TransactionResponse {
	m.findByIdResp.Id = transactionId
	return m.findByIdResp
}

func TestTransactionController_Create(t *testing.T) {
	mockSvc := &mockTransactionService{createResp: web.TransactionResponse{Id: 1, FromAccountId: 1, ToAccountId: 2, Amount: decimal.RequireFromString("100"), Type: "transfer", Description: "Test"}}
	ctrl := NewTransactionController(mockSvc)

	body, _ := json.Marshal(web.TransactionCreateRequest{FromAccountId: 1, ToAccountId: 2, Amount: decimal.RequireFromString("100"), Type: "transfer", Description: "Test"})
	req := httptest.NewRequest("POST", "/api/admin/transactions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	params := httprouter.Params{}

	ctrl.Create(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp web.WebResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	if resp.Code != 200 {
		t.Errorf("expected response code 200, got %d", resp.Code)
	}
}

func TestTransactionController_FindById(t *testing.T) {
	mockSvc := &mockTransactionService{findByIdResp: web.TransactionResponse{Id: 1, FromAccountId: 1, ToAccountId: 2, Amount: decimal.RequireFromString("100"), Type: "transfer"}}
	ctrl := NewTransactionController(mockSvc)

	req := httptest.NewRequest("GET", "/api/admin/transactions/1", nil)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{Key: "transactionId", Value: "1"}}

	ctrl.FindById(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp web.WebResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	data := resp.Data.(map[string]interface{})
	if int(data["id"].(float64)) != 1 {
		t.Errorf("expected transaction id 1, got %v", data["id"])
	}
}

func TestTransactionController_FindAll(t *testing.T) {
	mockSvc := &mockTransactionService{findAllResp: []web.TransactionResponse{
		{Id: 1, FromAccountId: 1, ToAccountId: 2, Amount: decimal.RequireFromString("100"), Type: "transfer"},
		{Id: 2, FromAccountId: 2, ToAccountId: 1, Amount: decimal.RequireFromString("50"), Type: "transfer"},
	}}
	ctrl := NewTransactionController(mockSvc)

	req := httptest.NewRequest("GET", "/api/admin/transactions", nil)
	w := httptest.NewRecorder()
	params := httprouter.Params{}

	ctrl.FindAll(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp web.WebResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	arr := resp.Data.([]interface{})
	if len(arr) != 2 {
		t.Fatalf("expected 2 transactions, got %d", len(arr))
	}
}

func TestTransactionController_GetMyTransactionsByAccountId(t *testing.T) {
	mockSvc := &mockTransactionService{findByAcctResp: []web.TransactionResponse{
		{Id: 1, FromAccountId: 1, ToAccountId: 2, Amount: decimal.RequireFromString("100"), Type: "deposit"},
	}}
	ctrl := NewTransactionController(mockSvc)

	req := httptest.NewRequest("GET", "/api/user/transactions/1", nil)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{Key: "accountId", Value: "1"}}

	ctrl.GetMyTransactionsByAccountId(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp web.WebResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	arr := resp.Data.([]interface{})
	if len(arr) != 1 {
		t.Fatalf("expected 1 transaction, got %d", len(arr))
	}
}

func TestTransactionController_Update(t *testing.T) {
	mockSvc := &mockTransactionService{updateResp: web.TransactionResponse{Id: 1, FromAccountId: 1, ToAccountId: 2, Amount: decimal.RequireFromString("200"), Type: "transfer", Description: "Updated"}}
	ctrl := NewTransactionController(mockSvc)

	body, _ := json.Marshal(web.TransactionUpdateRequest{FromAccountId: 1, ToAccountId: 2, Amount: decimal.RequireFromString("200"), Type: "transfer", Description: "Updated"})
	req := httptest.NewRequest("PUT", "/api/admin/transactions/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{Key: "transactionId", Value: "1"}}

	ctrl.Update(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp web.WebResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	data := resp.Data.(map[string]interface{})
	if data["amount"].(string) != "200" {
		t.Errorf("expected amount 200, got %s", data["amount"])
	}
}

func TestTransactionController_Delete(t *testing.T) {
	mockSvc := &mockTransactionService{}
	ctrl := NewTransactionController(mockSvc)

	req := httptest.NewRequest("DELETE", "/api/admin/transactions/1", nil)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{Key: "transactionId", Value: "1"}}

	ctrl.Delete(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !mockSvc.deleteCalled {
		t.Error("expected Delete to be called")
	}
}
