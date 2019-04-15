package core

import (
    "time"
    "os"
    "path/filepath"
    "log"
    "os/exec"
)


func Test_time() bool {
	nowtime := time.Now().Format("2006-01-02")
	timeover := "2019-05-31"

	//先把时间字符串格式化成相同的时间类型
	t1, _ := time.Parse("2006-01-02", nowtime)
	t2, _ := time.Parse("2006-01-02", timeover)

	a := t1.Before(t2)

	return a
}

func Get_lj() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	dir2 := dir + "/config.json"
	return dir2

}

func Start_v2() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	//v2lj := dir + "/v2ray"
	v2lj := dir + "/v2ray.exe"

	cmd := exec.Command(v2lj)
	cmd.Stdout = os.Stdout
	cmd.Run()

}





