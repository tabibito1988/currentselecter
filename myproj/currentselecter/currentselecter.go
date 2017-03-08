package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type Projectdata []struct {
	Name     string `json:"name"`
	Filepath string `json:"filepath"`
}

type Readfile interface {
	Readfile(filepath string) ([]byte, error)
}

type Readjson struct{}
type Readtxt struct{}

func (x Readjson) Readfile(filepath string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filepath)
	return bytes, err
}

func (x Readtxt) Readfile(filepath string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filepath)
	return bytes, err
}

func walk_func(path string, folders os.FileInfo, err error) error {
	// ディレクトリか判断
	//for folder := range folders {
	fmt.Println(folders)
	//if folder.IsDir() {
	// hoge/foo以下は無視する
	//	if folder.Name() == "currentselecterconfig" {
	// 無視するときはfilepath.SkipDirを戻す
	//		return filepath.SkipDir
	//	}
	//}
	//}

	// 普通はnilを戻す
	return nil
}

func start() {
	searchfoldername := "C:\\currentselecterconfig"
	filepath.Walk(searchfoldername, walk_func)
	//fmt.Println(path)
}

func main() {
	start()
	projectcount := 0
	var selectnumber string

	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("D:\\Users\\morimoto\\godev\\src\\github.com\\workspace\\myproj\\currentselecter\\currentpathdata.json")
	if err != nil {
		log.Fatal(err)
	}

	// JSONデコード
	var projectdatas Projectdata
	if err := json.Unmarshal(bytes, &projectdatas); err != nil {
		log.Fatal(err)
	}

	sprojectdatas := make(Projectdata, len(projectdatas)+1, len(projectdatas)+1)
	copy(sprojectdatas, projectdatas[:])
	sprojectdatas[cap(sprojectdatas)-1].Name = "update the config"
	sprojectdatas[cap(sprojectdatas)-1].Filepath = "D:\\Users\\morimoto\\godev\\src\\github.com\\workspace\\myproj\\currentselecter\\"
	Filename := "currentpathdata.json"

	//デコードしたデータを表示
	for _, sp := range sprojectdatas {
		projectcount++
		fmt.Printf("%d:%s\n", projectcount, sp.Name)
	}

	for {
		fmt.Scan(&selectnumber)
		inputnumber, _ := strconv.Atoi(selectnumber)
		if 0 < inputnumber && inputnumber <= projectcount {
			inputnumber--
			if inputnumber == projectcount-1 {
				os.Chdir(sprojectdatas[inputnumber].Filepath)
				exec.Command("notepad", Filename).Run()
			} else {
				exec.Command("code", sprojectdatas[inputnumber].Filepath).Run()
			}
			break
		} else {
			fmt.Printf("1-%dの中から選んでください。\n", projectcount)
		}
	}

}
