package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type Node struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func main() {
	root := "content"

	nodes := map[string]Node{}
	edges := make([]Edge, 0) // non-nil so it becomes [] not null

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(strings.ToLower(path), ".md") {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return nil
		}
		rel = filepath.ToSlash(rel)
		urlPath := "/" + strings.TrimSuffix(rel, ".md") + "/"
		label := strings.TrimSuffix(filepath.Base(path), ".md")

		nodes[urlPath] = Node{ID: urlPath, Label: label, URL: urlPath}

		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		content := string(data)

		linkPattern := regexp.MustCompile(`\[[^\]]*\]\((/[^)]+)\)`)
		matches := linkPattern.FindAllStringSubmatch(content, -1)
		for _, m := range matches {
			link := strings.TrimSuffix(m[1], ".md")
			if !strings.HasSuffix(link, "/") {
				link += "/"
			}
			edges = append(edges, Edge{From: urlPath, To: link})
		}

		return nil
	})

	if err != nil {
		fmt.Println("WalkDir error:", err)
		return
	}

	nodeKeys := make([]string, 0, len(nodes))
	for k := range nodes {
		nodeKeys = append(nodeKeys, k)
	}
	sort.Strings(nodeKeys)
	nodeList := make([]Node, 0, len(nodeKeys))
	for _, k := range nodeKeys {
		nodeList = append(nodeList, nodes[k])
	}

	os.MkdirAll("public", os.ModePerm)
	output := map[string]interface{}{
		"nodes": nodeList,
		"edges": edges,
	}

	outfile := "public/link-graph.json"
	f, err := os.Create(outfile)
	if err != nil {
		fmt.Println("create file error:", err)
		return
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(output); err != nil {
		fmt.Println("json encode error:", err)
	}

	fmt.Printf("✅ Link graph written to %s (%d nodes, %d edges)\n", outfile, len(nodeList), len(edges))
}
