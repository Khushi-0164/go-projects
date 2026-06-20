# 🎓 Student Management System
A modular CLI application for managing student records, demonstrating clean code architecture, dependency injection, and pointer manipulations in Go.

## 🚀 Features
* **Create/Read/Update/Delete** student records.
* **Architecture**: Fully modular package design with distinct database layers and service interfaces.
* **Database**: Local storage using `students.json`.

## 🛠️ Go Concepts Demonstrated
* **Modularization**: Codebase segmented into four packages: `models`, `storage`, `service`, and `main`.
* **Pointer vs. Value Receivers**: Knowing when to use struct pointer methods (`func (s *Service) ...`) to modify fields versus value receiver methods.
* **Mutating Slice Elements via Pointers**: Accessing and changing struct properties in-place inside slices by getting their memory address (`student := &students[i]`).
* **Dependency Injection**: Injecting the database storage dependency into the service manager, reducing structural coupling.

## 📖 How to Run
* **Add a Student:**
  ```bash
  go run . add "Khushi" 21
  ```
* **List Students:**
  ```bash
  go run . list
  ```
* **Update Student:**
  ```bash
  go run . update 1 "Khushi Patel" 22
  ```
* **Delete Student:**
  ```bash
  go run . delete 1
  ```
