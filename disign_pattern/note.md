module{
    //应该是foldername == packagename
    folder1 ==> package1 => [a.go, b.go...]
    folder2 ==> package2 => [c.go, d.go...]
}
在模块内引用如main.go引用另一个package中的interface，struct等时需要import("modulename/packagename")

**导入一定是以模块为单位的**，导package没用