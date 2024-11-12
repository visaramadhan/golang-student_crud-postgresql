package models

type Siswa struct{
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Class	string `json:"kelas"`
}