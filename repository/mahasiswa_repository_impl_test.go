package repository

import (
	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestMahasiswaInsert(t *testing.T) {
	mahasiswaRepository := NewMahasiswaRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	mahasiswa := entity.Mahasiswa{
		Name:    "Budi Kurniawan",
		Npm:     "01.07.0093",
		Jurusan: "Teknik Informatika",
	}
	result, err := mahasiswaRepository.Insert(ctx, mahasiswa)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestMahasiswaFindById(t *testing.T) {
	mahasiswaRepository := NewMahasiswaRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	mahasiswa, err := mahasiswaRepository.FindById(ctx, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(mahasiswa)
}

func TestMahasiswaFindAll(t *testing.T) {
	mahasiswaRepository := NewMahasiswaRepository(belajar_golang_database.GetConnection())
	ctx := context.Background()
	mahasiswas, err := mahasiswaRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, mahasiswa := range mahasiswas {
		fmt.Println(mahasiswa)
	}
}
