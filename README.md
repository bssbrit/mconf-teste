# mconf-teste

Book search application with API and CLI runner for the Mconf interview challenge.

## Overview

This project consists of two components:
- **API**: Go-based REST API that fetches book data from OpenLibrary
- **Runner**: Python CLI tool that queries the API and displays results

## Requirements

- Docker
- Git

## Usage

### 1. Clone the repository
```bash
git clone <repository-url>
cd mconf-teste
```

### 2. Build and run the API
```bash
# Build the API container
docker build -t mconf/api:candidato-1 ./api

# Run the API (will be accessible on port 3000)
docker run --ti --rm -p 3000:3000 mconf/api:candidato-1
```

### 3. Build and run the Runner
```bash
# Build the runner container
docker build -t mconf/runner:candidato-1 ./runner #(if in root)

# Run a book search (replace "Lord of the Rings" with your search term)
docker run --ti --rm --network host mconf/runner:candidato-1 "Lord of the Rings"
```

## API Endpoints

- `GET /?book_name={search_term}` - Search for books by name

## Example

```bash
# Search for Harry Potter books
docker run --ti --rm --network host mconf/runner:candidato-1 "Harry Potter"
```

The runner will display the book search results in JSON format with title, author, publication year, ISBN, and publisher information.

---

**Challenge**: https://gist.github.com/daronco/dd9698f7654b686aff9f31dae8ec7992