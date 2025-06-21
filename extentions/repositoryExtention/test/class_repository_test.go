package test

import (
	"cbt/extentions/models"
	repositoryextention "cbt/extentions/repositoryExtention"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func setupTestDB(t *testing.T) (repositoryextention.ClassRepositoryInterface, *gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err, "Gagal koneksi ke database in-memory")
	err = db.AutoMigrate(&models.Class{})
	require.NoError(t, err, "Gagal migrasi skema")
	repo := repositoryextention.NewClassRepository(db)
	return repo, db, err
}
func TestCreateClass(t *testing.T) {
	repo, _, _ := setupTestDB(t)
	newClass := &models.Class{
		ClassName:   "IPA",
		Description: "Class Paling Mantaps",
		GradeLevel:  "12",
	}

	newClass, err := repo.Create(newClass)

	require.NoError(t, err)
	require.NotZero(t, newClass.ID)
	require.Equal(t, "IPA", newClass.ClassName)
}

func TestGetProductByID(t *testing.T) {
	repo, _, _ := setupTestDB(t)
	existingClass := &models.Class{
		ClassName:   "IPA",
		Description: "Class Paling Mantaps",
		GradeLevel:  "12",
	}
	repo.Create(existingClass)
	require.NotZero(t, existingClass.ID)

	foundClass, err := repo.GetByID(existingClass.ID.String())

	// Assert:
	require.NoError(t, err)
	require.NotNil(t, foundClass)
	require.Equal(t, existingClass.ClassName, foundClass.ClassName)
	require.Equal(t, existingClass.GradeLevel, foundClass.GradeLevel)

	notFoundProduct, err := repo.GetByID("asdasd23123sadsa")

	// Assert:
	require.Error(t, err)
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
	require.Nil(t, notFoundProduct)
}

// // TestUpdateProduct memverifikasi fungsionalitas pembaruan produk.
// func TestUpdateProduct(t *testing.T) {
// 	repo, db := setupTestDB(t)

// 	// Arrange: Buat produk awal
// 	productToUpdate := &Product{Name: "Keyboard Mekanik", Price: 800000, Code: "K-MECH-03"}
// 	db.Create(productToUpdate)
// 	require.NotZero(t, productToUpdate.ID)

// 	// Modifikasi data produk
// 	updatedName := "Keyboard RGB"
// 	updatedPrice := uint(950000)
// 	productToUpdate.Name = updatedName
// 	productToUpdate.Price = updatedPrice

// 	// Act: Jalankan fungsi update
// 	err := repo.Update(productToUpdate)

// 	// Assert:
// 	require.NoError(t, err)

// 	// Verifikasi dengan mengambil data terbaru dari DB
// 	updatedProduct, _ := repo.GetByID(productToUpdate.ID)
// 	require.Equal(t, updatedName, updatedProduct.Name)
// 	require.Equal(t, updatedPrice, updatedProduct.Price)
// }

// // TestDeleteProduct memverifikasi fungsionalitas penghapusan produk.
// func TestDeleteProduct(t *testing.T) {
// 	repo, db := setupTestDB(t)

// 	// Arrange: Buat produk yang akan dihapus
// 	productToDelete := &Product{Name: "Webcam HD", Price: 500000, Code: "W-HD-04"}
// 	db.Create(productToDelete)
// 	require.NotZero(t, productToDelete.ID)

// 	// Act: Hapus produk
// 	err := repo.Delete(productToDelete.ID)
// 	require.NoError(t, err)

// 	// Assert: Coba cari lagi produk yang sudah dihapus
// 	_, err = repo.GetByID(productToDelete.ID)
// 	require.Error(t, err)                           // Harusnya error
// 	require.ErrorIs(t, err, gorm.ErrRecordNotFound) // Errornya harus not found
// }
