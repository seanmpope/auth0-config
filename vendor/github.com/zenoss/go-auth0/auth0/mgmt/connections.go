package mgmt

// Connection is a connection in Auth0
type Connection struct {
	Name               string                 `json:"name,omitempty"`
	Options            map[string]interface{} `json:"options,omitempty"`
	ID                 string                 `json:"id,omitempty"`
	Strategy           string                 `json:"strategy,omitempty"`
	Realms             []string               `json:"realms,omitempty"`
	EnabledClients     []string               `json:"enabled_clients,omitempty"`
	IsDomainConnection bool                   `json:"is_domain_connection,omitempty"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
}

// ConnectionSearch defines what can be used to search connections
type ConnectionSearch struct {
	PerPage       int      `json:"per_page,omitempty"`
	Page          int      `json:"page,omitempty"`
	Strategy      []string `json:"strategy,omitempty"`
	Name          string   `json:"name,omitempty"`
	Fields        string   `json:"fields,omitempty"`
	IncludeFields bool     `json:"include_fields,omitempty"`
}
