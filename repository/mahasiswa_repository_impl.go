package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type mahasiswaRepositoryImpl struct {
	DB *sql.DB
}

func NewMahasiswaRepository(db *sql.DB) MahasiswaRepository {
	return &mahasiswaRepositoryImpl{DB: db}
}
func (repository *mahasiswaRepositoryImpl) Insert(ctx context.Context, mahasiswa entity.Mahasiswa) (entity.Mahasiswa, error) {
	script := "INSERT INTO mahasiswa(name,npm,jurusan) VALUES (?,?,?)"
	result, err := repository.DB.ExecContext(ctx, script, mahasiswa.Name, mahasiswa.Npm, mahasiswa.Jurusan)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	mahasiswa.Id = int32(id)
	return mahasiswa, nil
}

func (repository *mahasiswaRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Mahasiswa, error) {
	script := "SELECT id,name,npm,jurusan FROM mahasiswa WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	mahasiswa := entity.Mahasiswa{}
	if err != nil {
		return mahasiswa, err
	}
	defer rows.Close()
	if rows.Next() {
		//	data ada
		rows.Scan(&mahasiswa.Id, &mahasiswa.Name, &mahasiswa.Name, &mahasiswa.Jurusan)
		return mahasiswa, nil
	} else {
		//	data tidak ada
		return mahasiswa, errors.New("Data Id " + strconv.Itoa(int(id)) + " Tidak ditemukan")
	}
}

func (repository *mahasiswaRepositoryImpl) FindAll(ctx context.Context) ([]entity.Mahasiswa, error) {
	script := "SELECT id, name,npm,jurusan FROM mahasiswa"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var mahasiswas []entity.Mahasiswa
	for rows.Next() {
		mahasiwa := entity.Mahasiswa{}
		rows.Scan(&mahasiwa.Id, &mahasiwa.Name, &mahasiwa.Npm, &mahasiwa.Jurusan)
		mahasiswas = append(mahasiswas, mahasiwa)
	}
	return mahasiswas, nil
}
