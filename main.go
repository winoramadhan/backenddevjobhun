package post

import (
    "net/http"
    "github.com/jmoiron/sqlx"
    "errors"
    
)

type mahasiswa struct{
    ID int
    Nama string
    Usia int
    Gender int
    Registrasi int
}

type jurusan struct{
    ID int
    Jurusan string
}

type hobi struct{
    ID mahasiswa
    ID hobi
}
