package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"golang-banking-api/model/domain"
	"golang-banking-api/model/web"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/shopspring/decimal"
)

type mockAccountService struct {
	createResp     web.AccountResponse
	updateResp     web.AccountResponse
	deleteCalled   bool
	findByIdResp   web.AccountResponse
	findAllResp    []web.AccountResponse
	findByUserResp []web.AccountResponse
}

func (m *mockAccountService) Create(ctx context.Context, request web.AccountCreateRequest) web.AccountResponse {
	return m.createResp
}
func (m *mockAccountService) Update(ctx context.Context, accountId int, request web.AccountUpdateRequest) web.AccountResponse {
	m.updateResp.Id = accountId
	return m.updateResp
}
func (m *mockAccountService) Delete(ctx context.Context, accountId int) {
	m.deleteCalled = true
}
func (m *mockAccountService) FindById(ctx context.Context, accountId int) web.AccountResponse {
	m.findByIdResp.Id = accountId
	return m.findByIdResp
}
func (m *mockAccountService) FindAll(ctx context.Context) []web.AccountResponse {
	return m.findAllResp
}
func (m *mockAccountService) FindByUserId(ctx context.Context, userId int) []web.AccountResponse {
	return m.findByUserResp
}

func TestAccountController_Create(t *testing.T) {
	mockSvc := &mockAccountService{createResp: web.AccountResponse{Id: 1, UserId: 2, AccountBank: "BCA", Balance: decimal.RequireFromString("1000"), AccountType: "savings"}}
	ctrl := NewAccountController(mockSvc)

	body, _ := json.Marshal(web.AccountCreateRequest{UserId: 2, AccountBank: "BCA", Balance: decimal.RequireFromString("1000"), AccountType: "savings"})
	req := httptest.NewRequest("POST", "/api/admin/accounts", bytes.NewReader(body))
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

func TestAccountController_FindById(t *testing.T) {
	mockSvc := &mockAccountService{findByIdResp: web.AccountResponse{Id: 1, UserId: 2, AccountBank: "BCA", Balance: decimal.RequireFromString("500"), AccountType: "checking"}}
	ctrl := NewAccountController(mockSvc)

	req := httptest.NewRequest("GET", "/api/admin/accounts/1", nil)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{Key: "accountId", Value: "1"}}

	ctrl.FindById(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp web.WebResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	data := resp.Data.(map[string]interface{})
	if int(data["id"].(float64)) != 1 {
		t.Errorf("expected account id 1, got %v", data["id"])
	}
}

func TestAccountController_GetMyAccounts(t *testing.T) {
	mockSvc := &mockAccountService{findByUserResp: []web.AccountResponse{
		{Id: 1, UserId: 1, AccountBank: "A", Balance: decimal.RequireFromString("100"), AccountType: "savings"},
	}}
	ctrl := NewAccountController(mockSvc)

	req := httptest.NewRequest("GET", "/api/user/accounts", nil)
	w := httptest.NewRecorder()
	params := httprouter.Params{}

	ctrl.GetMyAccounts(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp web.WebResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	arr := resp.Data.([]interface{})
	if len(arr) != 1 {
		t.Fatalf("expected 1 account, got %d", len(arr))
	}
}

func TestAccountController_GetMyAccountById(t *testing.T) {
	mockSvc := &mockAccountService{findByIdResp: web.AccountResponse{Id: 1, UserId: 1, AccountBank: "A", Balance: decimal.RequireFromString("100"), AccountType: "savings"}}
	ctrl := NewAccountController(mockSvc)

	req := httptest.NewRequest("GET", "/api/user/accounts/1", nil)
	req = req.WithContext(domain.ContextWithUser(req.Context(), 1, domain.RoleUser))
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{Key: "accountId", Value: "1"}}

	ctrl.GetMyAccountById(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp web.WebResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	data := resp.Data.(map[string]interface{})
	if int(data["id"].(float64)) != 1 {
		t.Errorf("expected account id 1, got %v", data["id"])
	}
}

func TestAccountController_FindAll(t *testing.T) {
	mockSvc := &mockAccountService{findAllResp: []web.AccountResponse{
		{Id: 1, UserId: 1, AccountBank: "A", Balance: decimal.RequireFromString("100"), AccountType: "savings"},
		{Id: 2, UserId: 2, AccountBank: "B", Balance: decimal.RequireFromString("200"), AccountType: "checking"},
	}}
	ctrl := NewAccountController(mockSvc)

	req := httptest.NewRequest("GET", "/api/admin/accounts", nil)
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
		t.Fatalf("expected 2 accounts, got %d", len(arr))
	}
}

func TestAccountController_Update(t *testing.T) {
	mockSvc := &mockAccountService{updateResp: web.AccountResponse{Id: 1, UserId: 1, AccountBank: "Updated", Balance: decimal.RequireFromString("999"), AccountType: "checking"}}
	ctrl := NewAccountController(mockSvc)

	body, _ := json.Marshal(web.AccountUpdateRequest{AccountBank: "Updated", Balance: decimal.RequireFromString("999"), AccountType: "checking"})
	req := httptest.NewRequest("PUT", "/api/admin/accounts/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{Key: "accountId", Value: "1"}}

	ctrl.Update(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp web.WebResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	data := resp.Data.(map[string]interface{})
	if data["account_bank"].(string) != "Updated" {
		t.Errorf("expected account_bank Updated, got %s", data["account_bank"])
	}
}

func TestAccountController_Delete(t *testing.T) {
	mockSvc := &mockAccountService{}
	ctrl := NewAccountController(mockSvc)

	req := httptest.NewRequest("DELETE", "/api/admin/accounts/1", nil)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{Key: "accountId", Value: "1"}}

	ctrl.Delete(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !mockSvc.deleteCalled {
		t.Error("expected Delete to be called")
	}
}
