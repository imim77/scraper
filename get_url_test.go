package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{

		{
			name:     "relative URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a>
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{},
		},
		{
			name:     "bad HTML",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html body>
	<a href="path/one">
		<span>Boot.dev</span>
	</a>
</html body>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "invalid href URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href=":\\invalidURL">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{},
		},
		{
			name:      "single absolute link",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><a href="https://blog.boot.dev/path">Link</a></body></html>`,
			expected:  []string{"https://blog.boot.dev/path"},
		},
		{
			name:      "multiple links",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><a href="/path1">Link1</a><a href="/path2">Link2</a></body></html>`,
			expected:  []string{"https://blog.boot.dev/path1", "https://blog.boot.dev/path2"},
		},
		{
			name:      "no links",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><p>No links here</p></body></html>`,
			expected:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, err := url.Parse(tt.inputURL)
			if err != nil {
				t.Errorf("couldn't parse input URL: %v", err)
				return
			}
			actual, err := getURLsFromHTML(tt.inputBody, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "absolute URL",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><img src="https://blog.boot.dev/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://blog.boot.dev/logo.png"},
		},
		{
			name:      "relative URL",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://blog.boot.dev/logo.png"},
		},
		{
			name:     "multiple images",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
				<img src="/logo.png" alt="Logo">
				<img src="https://cdn.boot.dev/banner.jpg">
			</body></html>`,
			expected: []string{
				"https://blog.boot.dev/logo.png",
				"https://cdn.boot.dev/banner.jpg",
			},
		},
		{
			name:      "no images",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><p>No images here</p></body></html>`,
			expected:  []string{},
		},
		{
			name:      "img without src",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><img alt="Logo"></body></html>`,
			expected:  []string{},
		},
		{
			name:      "invalid src URL",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><img src=":\\invalidURL" alt="Logo"></body></html>`,
			expected:  []string{},
		},
		{
			name:      "relative path without leading slash",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><img src="images/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://blog.boot.dev/images/logo.png"},
		},
		{
			name:     "mixed absolute and relative",
			inputURL: "https://blog.boot.dev/article",
			inputBody: `<html><body>
				<img src="/logo.png">
				<img src="banner.jpg">
				<img src="https://external.com/image.png">
			</body></html>`,
			expected: []string{
				"https://blog.boot.dev/logo.png",
				"https://blog.boot.dev/banner.jpg",
				"https://external.com/image.png",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsedURL, err := url.Parse(tt.inputURL)
			if err != nil {
				t.Fatalf("couldn't parse input URL: %v", err)
			}
			actual, err := getImagesFromHTML(tt.inputBody, parsedURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
