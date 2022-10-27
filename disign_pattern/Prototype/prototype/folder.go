package prototype

import "fmt"

type Folder struct {
	name     string
	children []Inode
}

func NewFolder(name string, children []Inode) *Folder {
	return &Folder{name: name, children: children}
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.Print(indentation + "    ")
	}
}
func (f *Folder) Clone() Inode {
	clonedFolder := &Folder{name: f.name + "_clone"}
	for _, child := range f.children {
		clonedFolder.children = append(clonedFolder.children, child.Clone())
	}
	return clonedFolder
}
