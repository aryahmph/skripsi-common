package user

func (m AuthUser) IsAdmin() bool {
	return m.Role == string(UserRoleAdmin)
}

func (m AuthUser) IsSame(ID string) bool {
	return m.ID == ID
}
