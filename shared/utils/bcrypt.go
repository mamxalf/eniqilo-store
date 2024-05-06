package utils

import (
	"github.com/mamxalf/eniqilo-store/config"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword will generate hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), config.Get().BcryptSalt)
	return string(bytes), err
}

// CheckPasswordHash will
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

// HashOTP will generate hashed password
func HashOTP(otp string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(otp), config.Get().BcryptSalt)
	return string(bytes), err
}

// CheckOTPHash will
func CheckOTPHash(otp, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(otp))

	return err
}

// HashPIN will generate hashed password
func HashPIN(pin string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pin), config.Get().BcryptSalt)
	return string(bytes), err
}

// CheckPINHash will
func CheckPINHash(pin, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pin))

	return err
}
