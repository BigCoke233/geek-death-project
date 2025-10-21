package main

import (
	"bufio"
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

type Frontmatter struct {
	Slug           string
	Aliases        []string
	Draft          bool
	Title          string
	RelatedCards   []string
	MentionedBooks []string
}

type Node struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
}

var (
	contentDir           = "content"
	internalLinkPattern   = regexp.MustCompile(`$begin:math:display$[^$end:math:display$]*\]$begin:math:text$(/[^)]+)$end:math:text$`)
	markdownLinkPattern   = regexp.MustCompile(`$begin:math:display$.*?$end:math:display$$begin:math:text$(http.*?)$end:math:text$`)
	urlMap                = make(map[string]string)       // file path -> canonical url
	aliasMap              = make(map[string]string)       // alias url -> canonical url
	linkGraph             = make(map[string][]string)     // canonical url -> links
	titleMap              = make(map[string]string)       // canonical url -> title
	externalLinks         = make(map[string]int)          // domain -> count
	excludedExternalHosts = map[string]bool{"image.guhub.cn": true, "www.geedea.pro": true}
)

func parseFrontmatter(filepath string) (Frontmatter, int) {
	file, err := os.Open(filepath)
	if err != nil {
		return Frontmatter{}, 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() || strings.TrimSpace(scanner.Text()) != "---" {
		return Frontmatter{}, 0
	}

	fm := Frontmatter{}
	lineNum := 1

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "---" {
			break
		}

		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			key := strings.TrimSpace(parts[0])
			value := strings.Trim(strings.TrimSpace(parts[1]), `"'`)

			switch key {
			case "aliases":
				for scanner.Scan() {
					next := strings.TrimSpace(scanner.Text())
					lineNum++
					if !strings.HasPrefix(next, "- ") {
						break
					}
					fm.Aliases = append(fm.Aliases, strings.TrimSpace(next[2:]))
				}
			case "slug":
				fm.Slug = value
			case "draft":
				fm.Draft = strings.ToLower(value) == "true"
			case "title":
				fm.Title = value
			case "related_cards":
				for scanner.Scan() {
					next := strings.TrimSpace(scanner.Text())
					lineNum++
					if !strings.HasPrefix(next, "- ") {
						break
					}
					fm.RelatedCards = append(fm.RelatedCards, strings.TrimSpace(next[2:]))
				}
			case "mentioned_books":
				for scanner.Scan() {
					next := strings.TrimSpace(scanner.Text())
					lineNum++
					if !strings.HasPrefix(next, "- ") {
						break
					}
					fm.MentionedBooks = append(fm.MentionedBooks, strings.TrimSpace(next[2:]))
				}
			}
		}
	}
	return fm, lineNum
}

func getCanonicalURL(filepath string, fm Frontmatter) string {
	rel, _ := filepathRel(contentDir, filepath)
	dir := filepath.Dir(rel)
	base := strings.TrimSuffix(filepath.Base(filepath), ".md")
	slug := fm.Slug
	if slug == "" {
		slug = strings.ToLower(base)
	}

	if dir != "." {
		return fmt.Sprintf("/%s/%s/", filepathToSlash(dir), slug)
	}
	return fmt.Sprintf("/%s/", slug)
}

func resolveURL(u string) string {
	if resolved, ok := aliasMap[u]; ok {
		return resolved
	}
	return u
}

func resolveURLFromDir(link, subdir string) string {
	targetDir := filepath.Join(contentDir, subdir)
	for path, canonical := range urlMap {
		if strings.HasPrefix(path, targetDir) {
			name := strings.TrimSuffix(filepath.Base(path), ".md")
			if name == link {
				return canonical
			}
		}
	}
	return ""
}

func getIcon(u string) string {
	switch {
	case strings.HasPrefix(u, "/cards/"):
		return "📄"
	case strings.HasPrefix(u, "/posts/"):
		return "📜"
	case strings.HasPrefix(u, "/library/"):
		return "📖"
	case strings.HasPrefix(u, "/fictions/"):
		return "🧙‍♀️"
	case strings.HasPrefix(u, "/en/"):
		return "🇬🇧"
	default:
		return "📁"
	}
}

func collectURLs() {
	filepath.WalkDir(contentDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".md") || d.Name() == "_index.md" {
			return nil
		}

		fm, _ := parseFrontmatter(path)
		if fm.Draft {
			return nil
		}

		canonical := getCanonicalURL(path, fm)
		urlMap[path] = canonical
		linkGraph[canonical] = []string{}

		title := fm.Title
		if title == "" {
			title = strings.Trim(canonical, "/")
			if parts := strings.Split(title, "/"); len(parts) > 0 {
				title = parts[len(parts)-1]
			}
		}
		titleMap[canonical] = fmt.Sprintf("%s %s", getIcon(canonical), title)

		for _, alias := range fm.Aliases {
			aliasMap[alias] = canonical
		}

		return nil
	})
}

func buildGraphAndExternalLinks() {
	for path, canonical := range urlMap {
		content, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		text := string(content)

		// Internal links
		for _, m := range internalLinkPattern.FindAllStringSubmatch(text, -1) {
			link := resolveURL(m[1])
			if _, ok := linkGraph[link]; ok {
				linkGraph[canonical] = append(linkGraph[canonical], link)
			}
		}

		// External links
		for _, m := range markdownLinkPattern.FindAllStringSubmatch(text, -1) {
			u, err := url.Parse(m[1])
			if err == nil && u.Host != "" && !excludedExternalHosts[u.Host] {
				externalLinks[u.Host]++
			}
		}

		// Related cards and books
		fm, _ := parseFrontmatter(path)
		for _, card := range fm.RelatedCards {
			if target := resolveURLFromDir(card, "cards"); target != "" {
				linkGraph[canonical] = append(linkGraph[canonical], target)
			}
		}
		for _, book := range fm.MentionedBooks {
			if target := resolveURLFromDir(book, "library"); target != "" {
				linkGraph[canonical] = append(linkGraph[canonical], target)
			}
		}
	}
}

func main() {
	collectURLs()
	buildGraphAndExternalLinks()

	os.MkdirAll("public", os.ModePerm)
	os.MkdirAll("data", os.ModePerm)

	// Write link-graph.json
	var nodes []Node
	var edges []Edge
	for url, targets := range linkGraph {
		nodes = append(nodes, Node{ID: url, Label: titleMap[url], URL: url})
		for _, tgt := range targets {
			edges = append(edges, Edge{From: url, To: tgt})
		}
	}
	writeJSON("public/link-graph.json", map[string]interface{}{"nodes": nodes, "edges": edges})
	fmt.Println("✅ Link graph generated: public/link-graph.json")

	// Write externallinks.json
	type kv struct {
		Key   string
		Count int
	}
	var sorted []kv
	for domain, count := range externalLinks {
		sorted = append(sorted, kv{domain, count})
	}
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].Count > sorted[j].Count })

	writeJSON("data/externallinks.json", sorted)
	fmt.Println("✅ External link data written to data/externallinks.json")
}

func writeJSON(path string, data interface{}) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("❌ Error creating file:", err)
		return
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		fmt.Println("❌ Error writing JSON:", err)
	}
}

func filepathRel(base, target string) (string, error) {
	rel, err := filepath.Rel(base, target)
	if err != nil {
		return "", err
	}
	return rel, nil
}

func filepathToSlash(p string) string {
	return strings.ReplaceAll(p, string(os.PathSeparator), "/")
}
