package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
)

type UploadInput struct {
	File        []byte
	Name        string
	Size        int64
	ContentType string
}

type FileStorage struct {
	db *sqlx.DB
}

func NewFileStorage(db *sqlx.DB) *FileStorage {
	return &FileStorage{
		db: db,
	}
}

func (r FileStorage) Upload(id int, input UploadInput) (string, error) {
	var fileName string
	pathDir := "assets/avatars"

	switch input.ContentType {
	case "image/jpeg":
		fileName = input.Name + ".jpeg"
	case "image/png":
		fileName = input.Name + ".png"
	}

	tempFile, err := ioutil.TempFile(pathDir, fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	tempFile.Write(input.File)

	query := fmt.Sprintf("UPDATE %s SET avatar_path=$1 WHERE id=$2;", userTable)

	_, err = r.db.Exec(query, pathDir+"/"+fileName, id)
	if err != nil {
		return "", err
	}

	return pathDir + "/" + fileName, nil
}
