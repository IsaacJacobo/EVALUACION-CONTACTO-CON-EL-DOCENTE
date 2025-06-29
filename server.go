package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Estructura del libro con JSON tags
type Book struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

// Base de datos en memoria
var books []Book

// Servicio 1: Listar libros
func listarLibros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Servicio 2: Agregar libro
func agregarLibro(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	var nuevo Book
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	books = append(books, nuevo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevo)
}

// Servicio 3: Buscar por autor
func buscarPorAutor(w http.ResponseWriter, r *http.Request) {
	autor := r.URL.Query().Get("autor")
	var encontrados []Book
	for _, b := range books {
		if strings.EqualFold(b.Author, autor) {
			encontrados = append(encontrados, b)
		}
	}
	if len(encontrados) == 0 {
		http.Error(w, "No se encontraron libros", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(encontrados)
}

// Servicio 4: Buscar por categoría
func buscarPorCategoria(w http.ResponseWriter, r *http.Request) {
	categoria := r.URL.Query().Get("categoria")
	var encontrados []Book
	for _, b := range books {
		if strings.EqualFold(b.Category, categoria) {
			encontrados = append(encontrados, b)
		}
	}
	if len(encontrados) == 0 {
		http.Error(w, "No se encontraron libros", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(encontrados)
}

// Servicio 5: Eliminar libro por índice
func eliminarLibro(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 || id >= len(books) {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	books = append(books[:id], books[id+1:]...)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Libro eliminado correctamente")
}

// Servicio 6: Editar libro
func editarLibro(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 0 || id >= len(books) {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var editado Book
	if err := json.NewDecoder(r.Body).Decode(&editado); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	books[id] = editado
	json.NewEncoder(w).Encode(editado)
}

// Servicio 7: Contar libros
func contarLibros(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Total de libros: %d", len(books))
}

// Servicio 8: Listar solo títulos
func listarTitulos(w http.ResponseWriter, r *http.Request) {
	var titulos []string
	for _, b := range books {
		titulos = append(titulos, b.Title)
	}
	json.NewEncoder(w).Encode(titulos)
}

// Configuración del servidor
func main() {
	http.HandleFunc("/libros", listarLibros)                     // GET
	http.HandleFunc("/libros/agregar", agregarLibro)             // POST
	http.HandleFunc("/libros/autor", buscarPorAutor)             // GET
	http.HandleFunc("/libros/categoria", buscarPorCategoria)     // GET
	http.HandleFunc("/libros/eliminar", eliminarLibro)           // DELETE con ?id=
	http.HandleFunc("/libros/editar", editarLibro)               // PUT con ?id=
	http.HandleFunc("/libros/count", contarLibros)               // GET
	http.HandleFunc("/libros/titulos", listarTitulos)            // GET

	fmt.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
