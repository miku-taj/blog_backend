package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"project/internal/model"
	"project/internal/utils"
	"time"
)

func (s *Service) ListUsers(page int) ([]*model.User, error) {
	return s.Repository.GetUsers(page)
}

func (s *Service) FindUserByID(userID int) (*model.User, error) {
	user, err := s.Repository.GetUserByID(userID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("Failed to find author by ID - Author with ID %d does not exist\n", userID)
	}
	user.Password = ""
	return user, err
}

func (s *Service) UsersByUsername(username string) ([]*model.User, error) {
	users, err := s.Repository.GetUsersByUsernameLike(username)
	if err == nil && len(users) == 0 {
		return nil, fmt.Errorf("Users with username containing %s not found\n", username)
	}
	return users, err
}

func (s *Service) UpdateUser(user *model.User, userID int) (*model.User, error) {
	user.UserID = userID
	oldUser, err := s.Repository.GetUserByID(user.UserID)
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("User with ID %d not found\n", user.UserID)
	}
	if err != nil {
		return nil, fmt.Errorf("Failed to Update User - Repository.GetUserByID error %v\n", err)
	}
	if user.Username == "" {
		user.Username = oldUser.Username
	} else {
		err = user.ValidateUsername()
		if err != nil {
			return nil, err
		}
		// Проверка существования пользователя
		userByUsername, err := s.Repository.GetUserByExactUsername(user.Username)

		// Если ошибка не nil и это не ошибка "record not found", возвращаем её
		if err != nil && !errors.As(err, &gorm.ErrRecordNotFound) {
			log.Printf("failed to get user: %v", err)
			return nil, err
		}

		// Если пользователь существует (userByID не nil), возвращаем ошибку
		if userByUsername != nil {
			log.Printf("user with name %s is already exists", user.Username)
			return nil, fmt.Errorf("user with name %s is already exists", user.Username)
		}
	}
	if user.Password == "" {
		user.Password = oldUser.Password
	} else {
		err = user.ValidatePassword()
		if err != nil {
			return nil, err
		}
		// Хешируем пароль пользователя
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to hash password: %w", err)
		}

		// Присваиваем хэшированный пароль пользователю
		user.Password = hashedPassword
	}
	currentTime := time.Now()
	user.UpdatedAt = &currentTime
	user.CreatedAt = oldUser.CreatedAt
	return s.Repository.EditUser(user)
}

func (s *Service) DeleteUser(id int) error {
	user, err := s.Repository.GetUserByID(id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if user == nil {
		return fmt.Errorf("Post with id %d not found %v", id, err)
	}

	return nil
}

func (s *Service) CreateUser(u *model.User) (*model.User, error) {
	user := model.User{Username: u.Username, Password: u.Password}
	err := user.ValidateUser()
	if err != nil {
		return nil, err
	}
	// Проверка существования пользователя
	userByUsername, err := s.Repository.GetUserByExactUsername(user.Username)

	// Если ошибка не nil и это не ошибка "record not found", возвращаем её
	if err != nil && !errors.As(err, &gorm.ErrRecordNotFound) {
		log.Printf("failed to get user: %v", err)
		return nil, err
	}

	// Если пользователь существует (userByID не nil), возвращаем ошибку
	if userByUsername != nil {
		log.Printf("user with name %s is already exists", user.Username)
		return nil, fmt.Errorf("user with name %s is already exists", user.Username)
	}

	// Хешируем пароль пользователя
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Присваиваем хэшированный пароль пользователю
	user.Password = hashedPassword

	// Добавляем пользователя в репозиторий
	return s.Repository.AddUser(&user)
}

func (s *Service) SignIn(u *model.User) (string, error) {
	user, err := s.Repository.GetUserByExactUsername(u.Username)

	if err != nil {
		if errors.As(err, &gorm.ErrRecordNotFound) {
			return "", fmt.Errorf("пользователь с именем %s не найден", u.Username)
		}
		return "", err
	}

	if !utils.CheckPasswordHash(u.Password, user.Password) {
		return "", fmt.Errorf("введен неправильный пароль")
	}

	token, err := utils.GenerateJWT(*user)
	if err != nil {
		return "", fmt.Errorf("не удалось сгенерировать токен: %w", err)
	}

	return token, nil
}
