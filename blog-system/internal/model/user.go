package model

import (
	"errors"
	"fmt"
	"regexp"
	"time"
	"unicode"
)

type User struct {
	UserID    int `gorm:"column:user_id;primaryKey"`
	Username  string
	Password  string
	CreatedAt *time.Time
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime:false"`
}

// ValidateUser функция, которая проверяет данные пользователя на соответствие заданным критериям.
func (u *User) ValidateUser() error {
	if err := u.ValidateUsername(); err != nil {
		return err
	}
	if err := u.ValidatePassword(); err != nil {
		return err
	}

	return nil
}

func (u *User) ValidateUsername() error {
	if len(u.Username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	return nil
}

// ValidatePassword проверяет сложность пароля.
func (u *User) ValidatePassword() error {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	if u.Password == "" {
		return fmt.Errorf("field password is empty")
	}

	if len(u.Password) >= 8 {
		hasMinLen = true
	}

	for _, char := range u.Password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasMinLen {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	if !hasUpper {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if !hasNumber {
		return fmt.Errorf("password must contain at least one number")
	}
	if !hasSpecial {
		return fmt.Errorf("password must contain at least one special character")
	}

	return nil
}

type Profile struct {
	ProfileID   int
	UserID      int
	DateOfBirth time.Time
	Location    string
	About       string
	Email       string
}

// ValidateProfile функция, которая проверяет данные профиля пользователя на соответствие заданным критериям.
func (p *Profile) ValidateProfile() error {
	if !p.isValidEmail() {
		return errors.New("invalid email format")
	}
	return nil
}

// isValidEmail проверяет корректность email с использованием регулярного выражения.
func (p *Profile) isValidEmail() bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(p.Email)
}
