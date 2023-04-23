package post

import (
    "log"
    "time"
    
    "github.com/jmoiron/sqlx"
    
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
