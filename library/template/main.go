package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	packagePath, packageName, serviceName string
	TempDir                               string = "library/template/temp"
)

func genTemlate(tmppath, path string) {
	temp, err := template.ParseFiles(TempDir + "/" + tmppath)
	die(err)

	var f *os.File
	dir := packagePath
	f, err = os.Create(fmt.Sprintf("%s/%s", dir, path))
	die(err)
	defer f.Close()

	temp.Execute(f, struct {
		PackagePath string
		PackageName string
		ServiceName string
	}{
		packagePath,
		packageName,
		serviceName,
	})
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Generate pacakge.path")
		return
	}

	packageName = os.Args[1]
	packagePath = strings.Replace(packageName, ".", "/", -1)

	arrs := strings.Split(packageName, ".")
	serviceName = strings.Title(arrs[len(arrs)-1])

	fmt.Println("Generate dir")
	os.Mkdir(filepath.Join(packagePath), os.ModePerm)

	//generate api
	os.Mkdir(filepath.Join(packagePath, "api"), os.ModePerm)
	genTemlate("api/temp.proto", fmt.Sprintf("api/%s.proto", serviceName))

	//generate cmd
	os.Mkdir(filepath.Join(packagePath, "cmd"), os.ModePerm)
	genTemlate("cmd/main.go", "cmd/main.go")

	//generate dao
	os.Mkdir(filepath.Join(packagePath, "dao"), os.ModePerm)
	genTemlate("dao/dao.go", "dao/dao.go")

	//generate hander
	os.Mkdir(filepath.Join(packagePath, "handler"), os.ModePerm)
	genTemlate("handler/init.go", "handler/init.go")
	genTemlate("handler/temp.go", fmt.Sprintf("handler/%s.go", serviceName))

	//generate model
	os.Mkdir(filepath.Join(packagePath, "model"), os.ModePerm)
	//generate utils
	os.Mkdir(filepath.Join(packagePath, "utils"), os.ModePerm)

	genTemlate("Makefile", "Makefile")

	//	fmt.Println(packageName, packagePath, serviceName)

}
