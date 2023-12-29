package model

// Veritabanı işlemleri için gerekli kütüphaneler eklenecek

// Todo modelini tanımla
type Todo struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Task   string `json:"task"`
}

// Kullanıcıya ait yapılacaklar listesini getirmek için fonksiyon
func GetTodosByUserID(userID int) ([]Todo, error) {
	// Veritabanından ilgili kullanıcının yapılacaklar listesini çekme işlemleri
	// Örnek olarak, basit bir liste döndürülüyor
	return []Todo{
		{ID: 1, UserID: userID, Task: "Alışveriş yap"},
		{ID: 2, UserID: userID, Task: "Fatura öde"},
	}, nil
}

// Yeni bir todo eklemek için fonksiyon
func AddTodo(userID int, task string) (*Todo, error) {
	// Veritabanına yeni bir todo ekleme işlemleri
	// Örnek olarak, basit bir todo nesnesi döndürülüyor
	newTodo := &Todo{
		UserID: userID,
		Task:   task,
	}
	return newTodo, nil
}
