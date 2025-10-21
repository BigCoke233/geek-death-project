package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type HighlightLink struct {
	Text string `json:"text"`
	Link string `json:"link"`
}

type Result struct {
	Title      string          `json:"title"`
	URL        string          `json:"url"`
	Highlights []HighlightLink `json:"highlights"`
}

func main() {
	const contentDir = "content"
	const outputFile = "data/highlights.json"

	highlightPattern := regexp.MustCompile(`==(.+?)==`)
	titlePattern := regexp.MustCompile(`(?m)^title:\s*["']?(.+?)["']?\s*$`)
	slugPattern := regexp.MustCompile(`(?m)^slug:\s*["']?(.+?)["']?\s*$`)

	var results []Result

	err := filepath.WalkDir(contentDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".md") {
			return nil
		}

		contentBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		content := string(contentBytes)

		// Title
		title := filepath.Base(path)
		if m := titlePattern.FindStringSubmatch(content); len(m) > 1 {
			title = m[1]
		}

		// Slug and URL
		urlPath := ""
		if m := slugPattern.FindStringSubmatch(content); len(m) > 1 {
			slug := strings.Trim(m[1], "/")
			relDir, _ := filepath.Rel(contentDir, filepath.Dir(path))
			relDir = filepath.ToSlash(relDir)
			if relDir != "." {
				urlPath = fmt.Sprintf("/%s/%s/", relDir, slug)
			} else {
				urlPath = fmt.Sprintf("/%s/", slug)
			}
		} else {
			rel, _ := filepath.Rel(contentDir, path)
			rel = filepath.ToSlash(rel)
			urlPath = "/" + strings.TrimSuffix(rel, ".md") + "/"
		}

		// Highlights
		matches := highlightPattern.FindAllStringSubmatch(content, -1)
		if len(matches) == 0 {
			return nil
		}

		var highlightLinks []HighlightLink
		for _, m := range matches {
			if len(m) > 1 {
				h := m[1]
				encoded := url.QueryEscape(h)
				link := fmt.Sprintf("%s#:~:text=%s", urlPath, encoded)
				highlightLinks = append(highlightLinks, HighlightLink{Text: h, Link: link})
			}
		}

		results = append(results, Result{
			Title:      title,
			URL:        urlPath,
			Highlights: highlightLinks,
		})

		return nil
	})

	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	// Write JSON
	if err := os.MkdirAll(filepath.Dir(outputFile), os.ModePerm); err != nil {
		fmt.Println("❌ Error creating output directory:", err)
		return
	}

	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Println("❌ Error encoding JSON:", err)
		return
	}

	if err := os.WriteFile(outputFile, jsonData, 0644); err != nil {
		fmt.Println("❌ Error writing JSON file:", err)
		return
	}

	fmt.Printf("✔ 提取完成，共 %d 篇文章有高亮，已写入 %s\n", len(results), outputFile)
}
