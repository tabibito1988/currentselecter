package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
)

type Projectdata []struct {
	Name     string `json:"name"`
	Filepath string `json:"filepath"`
}

func main() {

	projectcount := 0
	var i int
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

	//デコードしたデータを表示
	for _, p := range projectdatas {
		projectcount++
		fmt.Printf("%d  %s : %s\n", projectcount, p.Name, p.Filepath)
	}

	for {
		fmt.Scan(&selectnumber)
		i, _ = strconv.Atoi(selectnumber)
		if 0 < i && i <= projectcount {
			i--
			break
		} else {
			fmt.Printf("1-%dの中から選んでください。\n", projectcount)
		}
	}

	exec.Command("code", projectdatas[i].Filepath).Run()

}
