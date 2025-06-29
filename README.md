# EVALUACION-CONTACTO-CON-EL-DOCENTE
Este es mi programa de gestión de libros electrónicos en base a las directrices solicitadas 
# Sistema de Gestión de Libros Electrónicos

##  Grupo
- Isaac Chamorro

## Objetivo del programa
Desarrollar una aplicación CLI (interfaz de línea de comandos) en Go para gestionar libros electrónicos: agregar, buscar y listar libros usando principios de POO como encapsulación e interfaces. Se añadieron además 8 servicios web como parte del aprendizaje sobre APIs REST.
### Temas tratados en el semestre:

- ¿Qué es Go? Lenguaje compilado, rápido y concurrente.
- Funciones y estructuras de control (if, for, switch).
- Slices, arrays y mapas.
- Objetos y structs en Go.
- POO: Encapsulación, interfaces.
- Concurrencia con goroutines.
- Desarrollo de servicios web con JSON.

## Fecha
27 de Junio de 2025

## Funcionalidades principales
- Agregar libros con título, autor y categoría
- Listar todos los libros almacenados
- Buscar libros por nombre del autor
- Aplicación con interfaz CLI
- Servicios web con salida en JSON (GET, POST, etc.)

## Tecnologías usadas
- Go (Golang)
- JSON (para serialización de datos)
- API REST con net/http
- Git y GitHub para control de versiones

## Estructura del Proyecto

gestion-libros-electronicos/
├── main.go ← Aplicación de consola
├── server.go ← Servidor con 8 servicios web
├── README.md ← Explicación del proyecto
├── presentacion/ ← Diapositivas .pptx o exportadas
├── video/ ← Demostración grabada


---

## Servicios Web Implementados

| Servicio | Ruta                   | Método | Función                  |
|----------|------------------------|--------|---------------------------|
| 1        | `/libros`              | GET    | Listar libros             |
| 2        | `/libros/agregar`      | POST   | Agregar libro             |
| 3        | `/libros/autor`        | GET    | Buscar por autor          |
| 4        | `/libros/categoria`    | GET    | Buscar por categoría      |
| 5        | `/libros/eliminar`     | DELETE | Eliminar por ID (índice)  |
| 6        | `/libros/editar`       | PUT    | Editar libro por ID       |
| 7        | `/libros/count`        | GET    | Contar libros             |
| 8        | `/libros/titulos`      | GET    | Listar sólo los títulos   |

---

## Temas Integrados de la Asignatura

- ✅ ¿Qué es Go? (Lenguaje compilado, concurrente)
- ✅ Funciones y estructuras de control (`for`, `if`, `switch`)
- ✅ Slices y arrays
- ✅ Objetos y POO (struct `Book` con encapsulación)
- ✅ Interfaces (`Printable` en consola)
- ✅ Servicios Web (`net/http`, JSON)
- ✅ Concurrencia (Go puede usar `go func()` fácilmente)

---

## Visualización del Futuro

En una futura versión, este sistema se puede transformar en una plataforma educativa multiplataforma:

- Biblioteca digital para escuelas y colegios
- Sincronización con la nube
- App móvil con préstamos, puntuaciones, login de usuarios
- Recomendaciones por IA
- Multilenguaje y estadísticas de lectura

---

## Conclusión Reflexiva

Este proyecto me permitió aplicar todos los conocimientos aprendidos durante el semestre. Lo más desafiante fue la integración del sistema de consola con la lógica web, pero logré implementarlo con éxito. Aprendí a trabajar con servicios web, manejar estructuras de datos en Go, y representar la información en JSON para consumirla desde el navegador o herramientas como Postman.

