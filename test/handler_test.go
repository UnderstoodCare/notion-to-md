package test

import (
	"context"
	"os"
	"testing"

	notiontomd "github.com/Kible/notion-to-md"
	"github.com/joho/godotenv"
)

func TestPageToMarkdown(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Fatalf("Failed to load .env file: %v", err)
	}

	md, err := notiontomd.New(notiontomd.Params{
		Config: &notiontomd.Config{
			Notion: &notiontomd.NotionConfig{
				Token:           os.Getenv("NOTION_API_KEY"),
				ScrapeURLTitles: true,
			},
		},
	})

	if err != nil {
		t.Fatalf("Failed to create handler: %v", err)
	}

	blocks, err := md.PageToMarkdownFull(context.Background(), "1c757d75-c00f-80a2-960b-d63f54a9f2e4")
	if err != nil {
		t.Fatalf("Failed to get blocks: %v", err)
	}

	// Create a markdown file to write the output
	file, err := os.Create("./data/output.md")
	if err != nil {
		t.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()

	mdString, err := md.ToMarkdownString(blocks)
	if err != nil {
		t.Fatalf("Failed to convert blocks to markdown string: %v", err)
	}

	file.WriteString(mdString)

	outputContent, err := os.ReadFile("./data/output.md")
	if err != nil {
		t.Fatalf("Failed to read output.md: %v", err)
	}

	sourceOfTruthContent, err := os.ReadFile("./data/north-dakota-state-resources-export.md")
	if err != nil {
		t.Fatalf("Failed to read north-dakota-state-resources-export.md: %v", err)
	}

	// Compare the content
	if string(outputContent) != string(sourceOfTruthContent) {
		t.Errorf("Generated output does not match the actual value")
	} else {
		t.Logf("Generated output matches the actual value")
	}
}
