package main

import "prototype/prototype"

func main() {
	file1 := prototype.NewFile("file1")
	file2 := prototype.NewFile("file2")
	file3 := prototype.NewFile("file3")

	folder1 := prototype.NewFolder("folder1", []prototype.Inode{file2, file3})
	root := prototype.NewFolder("root", []prototype.Inode{file1, folder1})

	root.Print("")

	root_clone := root.Clone()
	root_clone.Print("")

}
