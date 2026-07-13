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
)

type mockUserService struct {
	createResp   web.UserResponse
	updateResp   web.UserResponse
	deleteCalled bool
	findByIdResp web.UserResponse
	findAllResp  []web.UserResponse
	getMeResp    web.UserResponse
	updateMeResp web.UserResponse
}

func (m *mockUserService) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	return m.createResp
}
func (m *mockUserService) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	m.updateResp.Id = request.Id
	return m.updateResp
}
func (m *mockUserService) Delete(ctx context.Context, userId int) {
	m.deleteCalled = true
}
func (m *mockUserService) FindById(ctx context.Context, userId int) web.UserResponse {
	m.findByIdResp.Id = userId
	return m.findByIdResp
}
func (m *mockUserService) FindAll(ctx context.Context) []web.UserResponse {
	return m.findAllResp
}
func (m *mockUserService) GetMe(ctx context.Context, userId int) web.UserResponse {
	m.getMeResp.Id = userId
	return m.getMeResp
}
func (m *mockUserService) UpdateMe(ctx context.Context, request web.UserUpdateMeRequest) web.UserResponse {
	return m.updateMeResp
}

func TestUserController_Create(t *testing.T) {
	mockSvc := &mockUserService{createResp: web.UserResponse{Id: 1, Email: "a@test.com", Name: "A", Role: "user"}}
	ctrl := NewUserController(mockSvc)

	body, _ := json.Marshal(web.UserCreateRequest{Email: "a@test.com", Password: "password", Name: "A", Role: "user"})
	req := httptest.NewRequest("POST", "/api/admin/users", bytes.NewReader(body))
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

func TestUserController_FindById(t *testing.T) {
	mockSvc := &mockUserService{findByIdResp: web.UserResponse{Id: 5, Email: "b@test.com", Name: "B", Role: "admin"}}
	ctrl := NewUserController(mockSvc)

	req := httptest.NewRequest("GET", "/api/admin/users/5", nil)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{Key: "userId", Value: "5"}}

	ctrl.FindById(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp web.WebResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	data := resp.Data.(map[string]interface{})
	if int(data["id"].(float64)) != 5 {
		t.Errorf("expected user id 5, got %v", data["id"])
	}
}

func TestUserController_GetMe(t *testing.T) {
	mockSvc := &mockUserService{getMeResp: web.UserResponse{Id: 1, Email: "me@test.com", Name: "Me", Role: "user"}}
	ctrl := NewUserController(mockSvc)

	req := httptest.NewRequest("GET", "/api/user/profile", nil)
	req = req.WithContext(domain.ContextWithUser(req.Context(), 1, domain.RoleUser))
	w := httptest.NewRecorder()
	params := httprouter.Params{}

	ctrl.GetMe(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp web.WebResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	data := resp.Data.(map[string]interface{})
	if int(data["id"].(float64)) != 1 {
		t.Errorf("expected user id 1, got %v", data["id"])
	}
}

func TestUserController_UpdateMe(t *testing.T) {
	mockSvc := &mockUserService{updateMeResp: web.UserResponse{Id: 1, Email: "updated@test.com", Name: "Updated", Role: "user"}}
	ctrl := NewUserController(mockSvc)

	body, _ := json.Marshal(web.UserUpdateMeRequest{Email: "updated@test.com", Password: "newpass", Name: "Updated"})
	req := httptest.NewRequest("PUT", "/api/user/profile", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	params := httprouter.Params{}

	ctrl.UpdateMe(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp web.WebResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	data := resp.Data.(map[string]interface{})
	if data["email"].(string) != "updated@test.com" {
		t.Errorf("expected email updated@test.com, got %s", data["email"])
	}
}

func TestUserController_FindAll(t *testing.T) {
	mockSvc := &mockUserService{findAllResp: []web.UserResponse{
		{Id: 1, Email: "a@test.com", Name: "A", Role: "user"},
		{Id: 2, Email: "b@test.com", Name: "B", Role: "admin"},
	}}
	ctrl := NewUserController(mockSvc)

	req := httptest.NewRequest("GET", "/api/admin/users", nil)
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
		t.Fatalf("expected 2 users, got %d", len(arr))
	}
}

func TestUserController_Delete(t *testing.T) {
	mockSvc := &mockUserService{}
	ctrl := NewUserController(mockSvc)

	req := httptest.NewRequest("DELETE", "/api/admin/users/1", nil)
	w := httptest.NewRecorder()
	params := httprouter.Params{httprouter.Param{Key: "userId", Value: "1"}}

	ctrl.Delete(w, req, params)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !mockSvc.deleteCalled {
		t.Error("expected Delete to be called")
	}
}
