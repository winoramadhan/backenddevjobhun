package post

import (
    "net/http"
    "github.com/jmoiron/sqlx"
    "errors"
    
)

type mahasiswa struct{
    ID          int     `json:"id"`
    Nama        string  `json:"nama"`
    Usia        int     `json:"usia"`
    Gender      int     'json:"gender"`
    Registrasi  int     'json:"registrasi"`
}

type jurusan struct{
    ID      int         `json:"id"`
    Jurusan string      `json:"jurusan"`
}

type hobi struct{
    ID mahasiswa
    ID hobi
}


func (c *Client) CreatePost(content string) (id int64, err error) {
	// run the insert script to db
	result, err := c.DB.Exec(`
		INSERT INTO post (content, create_time)
		VALUES (?, now())
	`,
		content)
	if err != nil {
		log.Println(err)
		return
	}

	// get the last insert id
	id, err = result.LastInsertId()
	if err != nil {
		log.Println(err)
		return
	}

	return
}
