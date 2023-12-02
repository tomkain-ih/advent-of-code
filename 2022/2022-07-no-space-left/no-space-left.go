package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Part 1
	//commands := exampleTerminalOutput()
	commands := terminalOutputFromFile()
	m := computeDiskUtilByDirectory(commands)
	fmt.Println(sumDirectoriesLteBytes(m, 100000))

	//Part 2
	filesystemMax := 70000000
	unusedMin := 30000000
	used := m["/"]
	unused := filesystemMax - used
	toClear := unusedMin - unused
	fmt.Println(getSmallestSatisfactoryDirectorySize(m, toClear))
}

func getSmallestSatisfactoryDirectorySize(m map[string]int, toClear int) (result int) {
	for _, v := range m {
		if v > toClear {
			if result <= 0 || v < result {
				result = v
			}
		}
	}
	return
}

func sumDirectoriesLteBytes(m map[string]int, ceiling int) (total int) {
	for _, e := range m {
		if e <= ceiling {
			total = total + e
		}
	}
	return
}

func computeDiskUtilByDirectory(commands []string) map[string]int {
	dirKey := ""
	d := ""
	m := make(map[string]int)
	for _, line := range commands {
		if strings.HasPrefix(line, "$ cd") {
			d = line[5:]
			if d == ".." {
				split := strings.Split(dirKey, "-")
				dirKey = strings.Join(split[:len(split)-1], "-")
			} else {
				if len(dirKey) <= 0 {
					dirKey = d
				} else {
					dirKey = dirKey + "-" + d
				}
			}
		} else if strings.HasPrefix(line, "$ ls") || strings.HasPrefix(line, "dir ") {
			continue
		} else {
			size, _ := strconv.Atoi(strings.Split(line, " ")[0])
			m[dirKey] = m[dirKey] + size
		}
	}
	return sumChildDirectorySizes(m)
}

func sumChildDirectorySizes(m map[string]int) (results map[string]int) {
	results = make(map[string]int)
	for k, e := range m {
		for _, d := range getDirectories(k) {
			results[d] = results[d] + e
		}
	}
	return results
}

func getDirectories(k string) (results []string) {
	keys := strings.Split(k, "-")
	for i := len(keys); i > 0; i-- {
		elems := keys[:i]
		join := strings.Join(elems, "/")
		replace := strings.Replace(join, "//", "/", 2)
		results = append(results, replace)
	}
	return
}

func terminalOutputFromFile() (result []string) {
	file, err := os.Open("2022-07-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return
}

func exampleTerminalOutput() []string {
	return strings.Split(
		`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`, "\n")
}
