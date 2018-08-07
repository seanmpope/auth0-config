package mgmt

import (
	"github.com/zenoss/go-auth0/auth0/http"
)

// UsersService provides a service for user related functions
type UsersService struct {
	c *http.Client
}

// User is a user in Auth0
type User struct {
	Email         string                 `json:"email,omitempty"`
	EmailVerified bool                   `json:"email_verified,omitempty"`
	Username      string                 `json:"username,omitempty"`
	PhoneNumber   string                 `json:"phone_number,omitempty"`
	PhoneVerified bool                   `json:"phone_verified,omitempty"`
	ID            string                 `json:"user_id,omitempty"`
	CreatedAt     string                 `json:"created_at,omitempty"`
	UpdatedAt     string                 `json:"updated_at,omitempty"`
	Identities    []Identity             `json:"identities,omitempty"`
	AppMetadata   map[string]interface{} `json:"app_metadata,omitempty"`
	UserMetadata  map[string]interface{} `json:"user_metadata,omitempty"`
	Picture       string                 `json:"picture,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Nickname      string                 `json:"nickname,omitempty"`
	Multifactor   []string               `json:"multifactor,omitempty"`
	LastIP        string                 `json:"last_ip,omitempty"`
	LastLogin     string                 `json:"last_login,omitempty"`
	LoginCount    uint                   `json:"logins_count,omitempty"`
	Blocked       bool                   `json:"blocked,omitempty"`
	FirstName     string                 `json:"given_name,omitempty"`
	LastName      string                 `json:"family_name,omitempty"`
}

// UserOpts are options which can be used to create a User
type UserOpts struct {
	ID            string                 `json:"user_id,omitempty"`
	Connection    string                 `json:"connection,omitempty"`
	Email         string                 `json:"email,omitempty"`
	Username      string                 `json:"username,omitempty"`
	Password      string                 `json:"password,omitempty"`
	PhoneNumber   string                 `json:"phone_number,omitempty"`
	UserMetadata  map[string]interface{} `json:"user_metadata,omitempty"`
	EmailVerified bool                   `json:"email_verified,omitempty"`
	VerifyEmail   bool                   `json:"verify_email,omitempty"`
	PhoneVerified bool                   `json:"phone_verified,omitempty"`
	AppMetadata   map[string]interface{} `json:"app_metadata,omitempty"`
}

// UserUpdateOpts are options which can be used to update a user
type UserUpdateOpts struct {
	Blocked           bool                   `json:"blocked,omitempty"`
	EmailVerified     bool                   `json:"email_verified,omitempty"`
	Email             string                 `json:"email,omitempty"`
	VerifyEmail       bool                   `json:"verify_email,omitempty"`
	PhoneNumber       string                 `json:"phone_number,omitempty"`
	PhoneVerified     bool                   `json:"phone_verified,omitempty"`
	VerifyPhoneNumber bool                   `json:"verify_phone_number,omitempty"`
	Password          string                 `json:"password,omitempty"`
	VerifyPassword    bool                   `json:"verify_password,omitempty"`
	UserMetadata      map[string]interface{} `json:"user_metadata,omitempty"`
	AppMetadata       map[string]interface{} `json:"app_metadata,omitempty"`
	Connection        string                 `json:"connection,omitempty"`
	Username          string                 `json:"username,omitempty"`
	ClientID          string                 `json:"client_id,omitempty"`
}

// Identity is the identity of a user in Auth0
type Identity struct {
	Connection string `json:"connection,omitempty"`
	ID         string `json:"user_id,omitempty"`
	Provider   string `json:"provider,omitempty"`
	IsSocial   bool   `json:"isSocial,omitempty"`
}

// GetAll returns all users
func (svc *UsersService) GetAll() ([]User, error) {
	var users []User
	err := svc.c.Get("/users", &users)
	return users, err
}

// Get returns a users
func (svc *UsersService) Get(userID string) (User, error) {
	var user User
	err := svc.c.Get("/users/"+userID, &user)
	return user, err
}

// Create creates a user
func (svc *UsersService) Create(opts UserOpts) (User, error) {
	var user User
	err := svc.c.Post("/users", opts, &user)
	return user, err
}

// Delete deletes a users
func (svc *UsersService) Delete(userID string) error {
	return svc.c.Delete("/users/"+userID, nil, nil)
}

// Update updates a user
func (svc *UsersService) Update(userID string, opts UserUpdateOpts) (User, error) {
	var user User
	err := svc.c.Patch("/users/"+userID, &opts, &user)
	return user, err
}
