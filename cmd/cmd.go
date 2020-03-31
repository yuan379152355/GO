package main

import (
    "fmt"
    "bufio"
    "os/exec"
    "io"
    "golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

// 工程属性
type VsProj struct {
    Projname    string
    Buildflag   string
    Env1        string
    Env2        string
}

func main() {
    var vsproj VsProj
    vsproj.Projname = "lbm_kpms_riskctrl"
    vsproj.Buildflag = "/rebuild"
    vsproj.Env1 = "oracle_debug"
    vsproj.Env2 = "x64"

    CallCmdExec(vsproj)
}

func CallCmdExec(vsproj VsProj){
    cmd := exec.Command("devenv.com", "D:\\codes\\citicsOL_patch\\src\\lbm\\" + vsproj.Projname + "\\" + vsproj.Projname + ".vcproj", vsproj.Buildflag, vsproj.Env1 + "|" + vsproj.Env2, "/Out")
    //显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return
	}
	
	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(ConvertByte2String([]byte(line),"GB18030"))
	}

	cmd.Wait()
    return
}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes,_=simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str= string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}