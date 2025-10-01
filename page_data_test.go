package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  PageData
	}{
		{
			name:     "complete page with all elements",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
				<h1>Test Title</h1>
				<p>This is the first paragraph.</p>
				<a href="/link1">Link 1</a>
				<img src="/image1.jpg" alt="Image 1">
			</body></html>`,
			expected: PageData{
				H1:             "Test Title",
				URL:            "https://blog.boot.dev",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks:  []string{"https://blog.boot.dev/link1"},
				ImageURLs:      []string{"https://blog.boot.dev/image1.jpg"},
			},
		},
		{
			name:     "multiple links and images",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
				<h1>Multiple Elements</h1>
				<p>First paragraph here.</p>
				<a href="/link1">Link 1</a>
				<a href="https://external.com">External</a>
				<img src="/image1.jpg">
				<img src="https://cdn.example.com/image2.png">
			</body></html>`,
			expected: PageData{
				H1:             "Multiple Elements",
				URL:            "https://blog.boot.dev",
				FirstParagraph: "First paragraph here.",
				OutgoingLinks: []string{
					"https://blog.boot.dev/link1",
					"https://external.com",
				},
				ImageURLs: []string{
					"https://blog.boot.dev/image1.jpg",
					"https://cdn.example.com/image2.png",
				},
			},
		},
		{
			name:     "missing h1",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
				<p>First paragraph.</p>
				<a href="/link1">Link 1</a>
			</body></html>`,
			expected: PageData{
				H1:             "",
				URL:            "https://blog.boot.dev",
				FirstParagraph: "First paragraph.",
				OutgoingLinks:  []string{"https://blog.boot.dev/link1"},
				ImageURLs:      []string{},
			},
		},
		{
			name:     "missing paragraph",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
				<h1>Title Only</h1>
				<a href="/link1">Link 1</a>
			</body></html>`,
			expected: PageData{
				H1:             "Title Only",
				URL:            "https://blog.boot.dev",
				FirstParagraph: "",
				OutgoingLinks:  []string{"https://blog.boot.dev/link1"},
				ImageURLs:      []string{},
			},
		},
		{
			name:     "no links",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
				<h1>No Links Page</h1>
				<p>Just text content.</p>
				<img src="/image.jpg">
			</body></html>`,
			expected: PageData{
				H1:             "No Links Page",
				URL:            "https://blog.boot.dev",
				FirstParagraph: "Just text content.",
				OutgoingLinks:  []string{},
				ImageURLs:      []string{"https://blog.boot.dev/image.jpg"},
			},
		},
		{
			name:     "no images",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
				<h1>No Images</h1>
				<p>Text only page.</p>
				<a href="/link1">Link</a>
			</body></html>`,
			expected: PageData{
				H1:             "No Images",
				URL:            "https://blog.boot.dev",
				FirstParagraph: "Text only page.",
				OutgoingLinks:  []string{"https://blog.boot.dev/link1"},
				ImageURLs:      []string{},
			},
		},
		{
			name:      "empty page",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body></body></html>`,
			expected: PageData{
				H1:             "",
				URL:            "https://blog.boot.dev",
				FirstParagraph: "",
				OutgoingLinks:  []string{},
				ImageURLs:      []string{},
			},
		},
		{
			name:     "main tag with paragraph",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
				<h1>Title</h1>
				<p>Outside paragraph.</p>
				<main>
					<p>Main paragraph.</p>
				</main>
			</body></html>`,
			expected: PageData{
				H1:             "Title",
				URL:            "https://blog.boot.dev",
				FirstParagraph: "Main paragraph.",
				OutgoingLinks:  []string{},
				ImageURLs:      []string{},
			},
		},
		{
			name:     "invalid links and images should be skipped",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
				<h1>With Invalid URLs</h1>
				<p>Content here.</p>
				<a href=":\\invalid">Bad Link</a>
				<a href="/good">Good Link</a>
				<img src=":\\invalid">
				<img src="/good.jpg">
			</body></html>`,
			expected: PageData{
				H1:             "With Invalid URLs",
				URL:            "https://blog.boot.dev",
				FirstParagraph: "Content here.",
				OutgoingLinks:  []string{"https://blog.boot.dev/good"},
				ImageURLs:      []string{"https://blog.boot.dev/good.jpg"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := extractPageData(tt.inputBody, tt.inputURL)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected %+v, got %+v", tt.expected, actual)
			}
		})
	}
}
