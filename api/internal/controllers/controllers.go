package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func extractBookName(r *http.Request) (string, error) {
    bookName := r.URL.Query().Get("book_name")
    if bookName == "" {
        return "", fmt.Errorf("book_name parameter is required")
    }
    // Replace spaces with "+" for URL encoding
    bookName = strings.ReplaceAll(bookName, " ", "+")
    return bookName, nil
}

func trimBookJSON(rawJSON []byte) ([]byte, error) {
    var response struct {
        Docs []struct {
            Title         string   `json:"title"`
            AuthorName    []string `json:"author_name"`
            FirstPublish  int      `json:"first_publish_year"`
            ISBN          []string `json:"isbn"`
            Publisher     []string `json:"publisher"`
        } `json:"docs"`
    }
    
    if err := json.Unmarshal(rawJSON, &response); err != nil {
        return nil, fmt.Errorf("invalid JSON response: %w", err)
    }
    
    // Return only the trimmed data
    trimmedData, err := json.Marshal(response.Docs)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal trimmed data: %w", err)
    }
    
    return trimmedData, nil
}

func fetchBookData(bookName string) ([]byte, error) {
    url := "https://openlibrary.org/search.json?q=" + bookName
    
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch book: %w", err)
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response: %w", err)
    }
    
    return body, nil
}


func FetchBook() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extract book name
        bookName, err := extractBookName(r)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        
        // Fetch book data
        rawData, err := fetchBookData(bookName)
        if err != nil {
            log.Println("Error:", err)
            http.Error(w, "Failed to fetch book", http.StatusInternalServerError)
            return
        }
        
        // Trim JSON to required fields
        trimmedData, err := trimBookJSON(rawData)
        if err != nil {
            log.Println("Error:", err)
            http.Error(w, "Failed to process response", http.StatusInternalServerError)
            return
        }
        
        // Set response headers and write the JSON response
        w.Header().Set("Content-Type", "application/json")
        w.Write(trimmedData)
    }
}