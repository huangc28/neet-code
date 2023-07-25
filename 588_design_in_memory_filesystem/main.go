package main

/*
trie
TrieNode
 children: map[byte]*TrieNode
 type: file | directory
 content: string

ls: list all TrieNode in the current directory
mkdir: traverse through every node. If node does not exists, create it
addContentToFile: traverse to that node, add content
readContentFromFile: traverse to that node, add content
*/

import (
	"strings"
)

type TrieNode struct {
	children map[string]*TrieNode
	isFile   bool
	content  string
}

type FileSystem struct {
	root *TrieNode
}

func Constructor() FileSystem {
	fs := FileSystem{
		root: &TrieNode{
			children: make(map[string]*TrieNode),
		},
	}

	fs.root.children[""] = &TrieNode{
		children: make(map[string]*TrieNode),
	}

	return fs
}

func (this *FileSystem) Ls(path string) []string {
	if path == "/" {
		return list(this.root, "")
	}
	return list(this.root, path)
}

// display
func list(node *TrieNode, path string) []string {
	curr := node

	segs := strings.Split(path, "/")
	for i := 0; i < len(segs); i++ {
		seg := segs[i]
		curr = curr.children[seg]
	}

	files := []string{}

	if curr.isFile {
		files = append(files, segs[len(segs)-1])
	} else {
		for path := range curr.children {
			files = append(files, path)
		}
	}

	return files
}

// insert
// mkdir("/") ---> [ ]
// mkdir("/a")---> [ , a]
func (this *FileSystem) Mkdir(path string) {
	insert(this.root, path)
}

func insert(node *TrieNode, path string) {
	curr := node
	segs := strings.Split(path, "/")
	for i := 0; i < len(segs); i++ {
		seg := segs[i]
		if _, exists := curr.children[seg]; !exists {
			curr.children[seg] = &TrieNode{
				children: make(map[string]*TrieNode),
			}
		}
		curr = curr.children[seg]
	}
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	curr := this.root

	segs := strings.Split(filePath, "/")
	pathsegs := segs[:len(segs)-1]
	filename := segs[len(segs)-1]

	for i := 0; i < len(pathsegs); i++ {
		pathseg := pathsegs[i]
		curr = curr.children[pathseg]
	}

	if _, fileExists := curr.children[filename]; !fileExists {
		curr.children[filename] = &TrieNode{
			children: make(map[string]*TrieNode),
			isFile:   true,
			content:  content,
		}

		return
	}

	curr.children[filePath].content += content

}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	curr := this.root

	segs := strings.Split(filePath, "/")
	for i := 0; i < len(segs); i++ {
		seg := segs[i]
		curr = curr.children[seg]
	}

	return curr.content
}
