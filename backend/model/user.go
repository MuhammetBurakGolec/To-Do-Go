package model

import "errors"

// Veritabanı bağlantı ve sorgulama için gerekli kütüphaneler eklenecek

// User modelini tanımla
type User struct {
	ID       int
	Username string
	Password string // Gerçek bir uygulamada şifre hash'lenmelidir
}

// Yeni bir kullanıcı oluşturmak için fonksiyon
func CreateUser(username, password string) (*User, error) {
	// Veritabanına kullanıcı ekleme işlemleri
	// Örnek olarak, basit bir user nesnesi döndürülüyor
	newUser := &User{
		Username: username,
		Password: password, // Gerçek bir uygulamada şifre hash'lenmelidir
	}
	return newUser, nil
}

// Kullanıcı doğrulama için fonksiyon
func AuthenticateUser(username, password string) (*User, error) {
	// Veritabanında kullanıcı sorgulama ve şifre doğrulaması
	// Örnek olarak, basit bir kontrol yapılıyor
	if username == "admin" && password == "password" {
		return &User{Username: username, Password: password}, nil
	}
	return nil, errors.New("Invalid username or password")
}
