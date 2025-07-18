package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchBook() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Define the external API URL
        url := "https://openlibrary.org/search.json?q=" // Change to your desired URL

		fmt.Println(r.Body)
        bookName := r.URL.Query().Get("name")
		fmt.Print("hi" + bookName)
        if bookName == "" {
			http.Error(w, "Book name is required", http.StatusBadRequest)
			fmt.Println(bookName)
		}
		url = url + bookName
		
        resp, err := http.Get(url)
        if err != nil {
            http.Error(w, "Failed to fetch book", http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        // Read the response body
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            http.Error(w, "Failed to read response", http.StatusInternalServerError)
            return
        }

        // Optionally, validate that the response is valid JSON
        var js json.RawMessage
        if err := json.Unmarshal(body, &js); err != nil {
            http.Error(w, "Invalid JSON response", http.StatusInternalServerError)
            return
        }

        // Set response headers and write the JSON response
        w.Header().Set("Content-Type", "application/json")
        w.Write(body)
    }
}