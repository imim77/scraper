package main

import "testing"

func TestGetH1FromHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic h1 tag",
			input:    "<html><body><h1>Test Title</h1></body></html>",
			expected: "Test Title",
		},
		{
			name:     "no h1 tag",
			input:    "<html><body><h2>Not H1</h2></body></html>",
			expected: "",
		},
		{
			name:     "multiple h1 tags",
			input:    "<html><body><h1>First</h1><h1>Second</h1></body></html>",
			expected: "First",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := getH1FromHTML(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}
		})
	}

}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "outside and main paragraph",
			input: `<html><body>
				<p>Outside paragraph.</p>
				<main>
					<p>Main paragraph.</p>
				</main>
			</body></html>`,
			expected: "Main paragraph.",
		},
		{
			name: "no main tag - gets first paragraph",
			input: `<html><body>
				<p>First paragraph.</p>
				<p>Second paragraph.</p>
			</body></html>`,
			expected: "First paragraph.",
		},
		{
			name: "main with multiple paragraphs",
			input: `<html><body>
				<main>
					<p>First in main.</p>
					<p>Second in main.</p>
				</main>
			</body></html>`,
			expected: "First in main.",
		},
		{
			name: "no paragraphs",
			input: `<html><body>
				<div>No paragraphs here</div>
			</body></html>`,
			expected: "",
		},
		{
			name: "empty paragraph",
			input: `<html><body>
				<main>
					<p></p>
					<p>Second paragraph.</p>
				</main>
			</body></html>`,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tt.input)
			if actual != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, actual)
			}
		})
	}
}
