package core

import (
    "time"
    "os"
    "path/filepath"
    "log"
    "os/exec"
    "archive/zip"
    "io"
    "io/ioutil"
    "runtime"
    //"fmt"
)

func Check_system() string {
    //检查系统的函数
//case "darwin": case "windows": case "linux":

    sysname := runtime.GOOS
    return sysname
}



func Test_time() bool {
    //检查时间的函数
	nowtime := time.Now().Format("2006-01-02")
	timeover := "2019-05-31"

	//先把时间字符串格式化成相同的时间类型
	t1, _ := time.Parse("2006-01-02", nowtime)
	t2, _ := time.Parse("2006-01-02", timeover)

	a := t1.Before(t2)

	return a
}

func Get_lj() (string,string,string) {
    //获取相对路径的函数
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	dir2 := dir + "/v2ray/v2ray/config.json"
    dir3 := dir + "/V.zip"
    dir4 := dir + "/v2ray"
	return dir2,dir3,dir4
}

func Start_v2(sysname string) {

    if sysname == "darwin"{
        //启动v2的函数
        dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
        if err != nil {
            log.Fatal(err)
        }

        //v2lj := dir + "/v2ray"
        v2lj := dir + "/v2ray/v2ray/v2ray"

        cmd := exec.Command(v2lj)
        cmd.Stdout = os.Stdout
        cmd.Run()
    }

    if sysname == "windows"{
        //启动v2的函数
        dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
        if err != nil {
            log.Fatal(err)
        }

        //v2lj := dir + "/v2ray"
        v2lj := dir + "/v2ray/v2ray/v2ray.exe"

        cmd := exec.Command(v2lj)
        cmd.Stdout = os.Stdout
        cmd.Run()

    }

}


func Unzip(zipFile string, destDir string) error {
    //解压压缩包的函数
    zipReader, err := zip.OpenReader(zipFile)
    if err != nil {
        return err
    }
    defer zipReader.Close()
 
    for _, f := range zipReader.File {
        fpath := filepath.Join(destDir, f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(fpath, os.ModePerm)
        } else {
            if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
                return err
            }
 
            inFile, err := f.Open()
            if err != nil {
                return err
            }
            defer inFile.Close()
 
            outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer outFile.Close()
 
            _, err = io.Copy(outFile, inFile)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func Get_file(x string) []string {
    //获取文件夹下文件名称的函数
    a := []string {}

    files, _ := ioutil.ReadDir(x)

        for _, f := range files {
            a = append(a,f.Name())
        }

    return a
}

func PathExists(path string) (bool) {
    //验证文件夹是否存在的函数
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    if os.IsNotExist(err) {
        return false
    }
    return false
}


func Yz_v2_wz() bool {
    //判断文件夹是否存在的函数
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }

    lj := dir + "/v2ray"

    if PathExists(lj) {
        return true
    } else {
        return false
    }


}



func Rmconfig() {
    //删除配置文件的函数
    //判断文件夹是否存在的函数
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }

    lj := dir + "/v2ray/v2ray/config.json"
    time.Sleep(time.Duration(1)*time.Second)
    _ = os.Remove(lj) 




}


