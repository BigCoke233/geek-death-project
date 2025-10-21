package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func main() {
	externalLinks := countExternalLinks("content")
	writeJSON("data/externallinks.json", externalLinks)
	fmt.Println("✅ External link rank data generated.")
}

// countExternalLinks scans markdown files and counts occurrences of each external domain.
func countExternalLinks(root string) [][2]interface{} {
	linkRegex := regexp.MustCompile(`https?://[^\s)>"']+`)
	counts := make(map[string]int)

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		matches := linkRegex.FindAllString(string(data), -1)
		for _, match := range matches {
			u, err := url.Parse(match)
			if err == nil && u.Host != "" && u.Host != "image.guhub.cn" {
				counts[u.Host]++
			}
		}
		return nil
	})

	// convert map to sorted array
	type kv struct {
		Key   string
		Value int
	}
	var list []kv
	for k, v := range counts {
		list = append(list, kv{k, v})
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Value > list[j].Value })

	result := make([][2]interface{}, len(list))
	for i, kv := range list {
		result[i] = [2]interface{}{kv.Key, kv.Value}
	}
	return result
}

// writeJSON writes the result to file with indentation.
func writeJSON(path string, data interface{}) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	enc.Encode(data)
}
