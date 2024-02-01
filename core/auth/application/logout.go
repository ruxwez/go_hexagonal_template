package authApplication

// Logout implements AuthService.
func (a *authService) Logout(token string) error {
	return a.auth.DeleteSession(token)
}
