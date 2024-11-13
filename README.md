# Book Library Management System

A simple Book Library Management System implemented in Go. This system allows users to manage books through both a Command Line Interface (CLI) and a RESTful API. Users can add, view, update, and delete books either via terminal commands or HTTP requests.

## Features

- **CLI Interface**:
  - View all books
  - View book by ID
  - Add a new book
  - Update book by ID
  - Delete book by ID
  - Interactive terminal-based menu

- **RESTful API**:
  - `GET /books`: View all books
  - `GET /books/{id}`: View a specific book by ID
  - `POST /books`: Create a new book
  - `PUT /books/{id}`: Update a book by ID
  - `DELETE /books/{id}`: Delete a book by ID

## Installation

To run the Book Library Management System, you need to have Go installed on your machine. Follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/book-library-management.git
   cd book-library-management
