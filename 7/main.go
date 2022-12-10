package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getLines(s string) []string {
	return strings.Split(strings.Trim(s, "\n"), "\n")
}

func main() {
	data, e := os.ReadFile("input7")
	check(e)

	answer1 := task1(string(data))
	fmt.Println("Answer1:", answer1)
	answer2 := task2(string(data))
	fmt.Println("Answer2:", answer2)
}

func task1(data string) int {
	result := 0
	dir := parseInput(data)
	for _, d := range dir.getDirs() {
		size := d.getSize()
		if size < 100000 {
			result += size
		}
	}
	return result
}

func task2(data string) int {
	dir := parseInput(data)
	totalSpace := 70000000
	neededSpace := 30000000
	usedSpace := dir.getSize()
	freeSpace := totalSpace - usedSpace
	spaceToClear := neededSpace - freeSpace
	result := math.MaxInt
	for _, d := range dir.getDirs() {
		size := d.getSize()
		if size > spaceToClear && size < result {
			result = size
		}
	}
	return result
}

type File int
type Dir struct {
	files map[string]File
	dirs  map[string]Dir
}

func makeDir() Dir {
	d := Dir{}
	d.files = map[string]File{}
	d.dirs = map[string]Dir{}
	return d
}

func (d *Dir) addFile(name string, size File) {
	d.files[name] = size
}

func (d *Dir) addDir(name string) {
	d.dirs[name] = makeDir()
}

func (d *Dir) getSize() int {
	size := 0
	for _, f := range d.files {
		size += int(f)
	}
	for _, dd := range d.dirs {
		size += dd.getSize()
	}
	return size
}

func (d *Dir) getDirs() []Dir {
	dirs := []Dir{*d}
	for _, dd := range d.dirs {
		dirs = append(dirs, dd.getDirs()...)
	}
	return dirs
}

func parseInput(data string) Dir {
	dirs := []Dir{makeDir()}
	for _, line := range getLines(data) {
		parts := strings.Split(line, " ")
		if parts[0] == "$" && parts[1] == "cd" {
			subdirName := parts[2]
			if subdirName == "/" {
				dirs = dirs[:1]
			} else if subdirName == ".." {
				dirs = dirs[:len(dirs)-1]
			} else {
				subdir := dirs[len(dirs)-1].dirs[subdirName]
				dirs = append(dirs, subdir)
			}
		}
		if parts[0] == "dir" {
			subdirName := parts[1]
			dirs[len(dirs)-1].addDir(subdirName)
		}
		size, ok := strconv.Atoi(parts[0])
		if ok == nil {
			fileName := parts[1]
			dirs[len(dirs)-1].addFile(fileName, File(size))
		}
	}
	return dirs[0]
}
