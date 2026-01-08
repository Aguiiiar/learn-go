package currentuser

// CurrentUserScope = 'identity' | 'clinic'
type CurrentUserScope string

const (
	ScopeIdentity CurrentUserScope = "identity"
	ScopeClinic   CurrentUserScope = "clinic"
)

// CurrentUserType (você usou no TS, então deixei aqui também)
type CurrentUserType string

const (
	TypeUser   CurrentUserType = "user"
	TypeSystem CurrentUserType = "system"
)

// CurrentUserProps equivale ao objeto do constructor no TS
type CurrentUserProps struct {
	UserID      string
	OrgID       *string
	ClinicID    *string
	DBSchema    *string
	Permissions []string
	Type        *CurrentUserType
	Scope       *CurrentUserScope
}

// CurrentUser equivale à class
type CurrentUser struct {
	userID      string
	scope       CurrentUserScope
	orgID       *string
	clinicID    *string
	dbSchema    *string
	permissions []string
	typ         CurrentUserType
}

// NewCurrentUser equivale ao constructor + defaults
func NewCurrentUser(p CurrentUserProps) *CurrentUser {
	// defaults do TS:
	// permissions = []
	// scope = 'identity'
	// type = 'user'
	perms := p.Permissions
	if perms == nil {
		perms = []string{}
	}

	scope := ScopeIdentity
	if p.Scope != nil {
		scope = *p.Scope
	}

	typ := TypeUser
	if p.Type != nil {
		typ = *p.Type
	}

	return &CurrentUser{
		userID:      p.UserID,
		orgID:       p.OrgID,
		clinicID:    p.ClinicID,
		dbSchema:    p.DBSchema,
		permissions: perms,
		typ:         typ,
		scope:       scope,
	}
}

/* ===== Getters (porque os campos são privados) ===== */

func (c *CurrentUser) UserID() string          { return c.userID }
func (c *CurrentUser) Scope() CurrentUserScope { return c.scope }
func (c *CurrentUser) OrgID() *string          { return c.orgID }
func (c *CurrentUser) ClinicID() *string       { return c.clinicID }
func (c *CurrentUser) DBSchema() *string       { return c.dbSchema }
func (c *CurrentUser) Permissions() []string   { return append([]string(nil), c.permissions...) }
func (c *CurrentUser) Type() CurrentUserType   { return c.typ }

/* ===== Métodos equivalentes ===== */

func (c *CurrentUser) IsSystem() bool {
	return c.typ == TypeSystem
}

func (c *CurrentUser) IsUser() bool {
	return c.typ == TypeUser
}

func (c *CurrentUser) HasRole(role string) bool {
	for _, r := range c.permissions {
		if r == role {
			return true
		}
	}
	return false
}
