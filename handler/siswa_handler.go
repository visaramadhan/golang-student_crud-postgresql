package handler

import (
	"fmt"

	"github.com/go-embed-go-web/models"
	"github.com/go-embed-go-web/service"
)

type SiswaHandler struct {
	service service.SiswaService
}

func NewSiswaHandler(service service.SiswaService) *SiswaHandler {
	return &SiswaHandler{service: service}
}

func (h *SiswaHandler) AddStudent() {
	var name string
	var class string
	fmt.Print("Enter student name: ")
	fmt.Scan(&name)
	fmt.Print("Enter student class: ")
	fmt.Scan(&class)

	siswa := models.Siswa{Nama: name, Class: class}
	err := h.service.RegisterNewSiswa(siswa)
	if err != nil {
		fmt.Println("Error adding student: ", err)
	} else {
		fmt.Println("Student added successfully")
	}
}

func (h *SiswaHandler) ViewStudents() {
	siswas, err := h.service.GetAllSiswa()
	if err != nil {
		fmt.Println("Error fetching students: ", err)
	} else {
		fmt.Println("Students:")
		for _, siswa := range siswas {
			fmt.Printf("%s - %s\n", siswa.Nama, siswa.Class)
		}
	}
}

func (h *SiswaHandler) DeleteStudent() {
	var id int
	fmt.Print("Enter student ID to delete: ")
	fmt.Scan(&id)

	err := h.service.DeleteSiswa(id)
	if err != nil {
		fmt.Println("Error deleting student: ", err)
	} else {
		fmt.Println("Student deleted successfully")
	}

}

func (h *SiswaHandler) UpdateStudent() {
	var id int
	var name string
	var class string
	fmt.Print("Enter student ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter new student name: ")
	fmt.Scan(&name)
	fmt.Print("Enter new student class: ")
	fmt.Scan(&class)

	siswa := models.Siswa{ID: id, Nama: name, Class: class}
	err := h.service.UpdateDataSiswa(siswa)
	if err != nil {
		fmt.Println("Error updating student: ", err)
	} else {
		fmt.Println("Student updated successfully")
	}
}
