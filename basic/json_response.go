package basic

import (
	"fmt"
	"io"
	"net/http"
)

func json_response(w io.Writer) error {
	fmt.Fprintln(w, "Hello, Go Standart Library")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/post/1")
	if err != nil {
		return fmt.Errorf("Error: %w\n", err)
	}
	defer resp.Body.Close()

	fmt.Fprintf(w, "HTTP Response Status: %s\n", resp.Status)
	return nil
}
