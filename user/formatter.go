package user

type UserFormatter struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	AvatarURL string `json:"avatar_url"`
}

func FormatUser(user User, token string) UserFormatter {
	formatUser := UserFormatter{
		ID:        int(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Token:     user.Token,
		AvatarURL: user.AvatarFile,
	}

	return formatUser
}
