/*!
 *test.go version 4.0.0 
 *
 *Copyright © 2021 All Rights Reserved.
 *Author: Ryuto Otsubo
 *
 *Released under the MIT license.
 *See https://opensource.org/licenses/MIT
 *
 *Git: https://github.com/OtsuboRyuto/LMS
 *
 *------------------- References --------------------
 *
 *https://qiita.com/chan-p/items/cf3e007b82cc7fce2d81
 *https://gorm.io/docs/
 *https://golang.org/pkg/  基本的に情報はここから調べました。
 *
 *
 *---------------------------------------------------
 */
package main
/*
自作関数whetherErrについて
私のプログラムは基本例外を-1として扱っているため、ステータス-1でパニックを起こすプログラムである。

例外は-1を返し、成功は1, 失敗は0で返しています。
*/
import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"reflect"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
	//"bufio"
	"io/ioutil"
	"os"
	"strings"
	//"math"
	"io"
	"path/filepath"
	"crypto/sha256"
    "encoding/hex"
)

type Bb struct {//BBS用のデータ構造
	Id       string
	Group_id int
	Contents string
}

type How_many struct {//何グループか
	Num_of_groups int
}

type Unique_num struct {//固有のid(グループ振り分け用)
	Id        int
	Unique_id int
}

type retData struct { //Bbs起動時に返却するデータ用の関数
	Your_group string
	Hash       string
}

type Auth struct {
	Id         int
	User_print string
}

type Test struct {
	Data []string
}

type History struct {
	Id         string
	Point      int
	Test_title string
}

type Comment struct {
	Id       int
	Contents string
}

type Users struct{
	Id int
	Name string
	Passward string
	Question string
	Answer string
}
type ContentTest struct{
	Sub string
	Day string
	Number_of_ques int
	Part []string
	Ques []string
	Answer_list []string
	Answer []string
	Point []string
}
func PWD(){
	//現在のディレクトリを取得してそれを出力する
	 p, _ := os.Getwd()
 	 fmt.Println(p)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/get", fetchAriticle)
	http.HandleFunc("/add", addAriticle)
	http.HandleFunc("/dec", dec)
	http.HandleFunc("/test", receieveTest)
	http.HandleFunc("/getall", fetchAriticleAll)
	http.HandleFunc("/integration", integration)
	http.HandleFunc("/addscore", addScore)
	http.HandleFunc("/testlist", getAlldir)
	http.HandleFunc("/getHistories", getHistories)
	http.HandleFunc("/getTeacherComment", getTeacherComment)
	http.HandleFunc("/insertTeacherComment", insertTeacherComment)
	http.HandleFunc("/updateHow_many", updateHow_many)
	http.HandleFunc("/log_in", log_in)
	http.HandleFunc("/auth", auth)
	http.HandleFunc("/registration", registration)
	http.HandleFunc("/c_auth", checkAuth)
	http.HandleFunc("/checkIp", checkIp)
	http.HandleFunc("/checkPass", checkPass)
	http.HandleFunc("/addUnique", addUnique)
	http.HandleFunc("/registrationTest", registrationTest)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func whetherErr(err error) {
	//例外パターンは-1を返却
	if err != nil {
		fmt.Print("fatal error \"exception received!\" \n")
		panic(-1)
	}
}
func getSHA256(s string) string {
	//SHA256でハッシュにし、そのデータを文字列にして返却する
    r := sha256.Sum256([]byte(s))
    return hex.EncodeToString(r[:])
}

func registrationTest(w http.ResponseWriter, r *http.Request){
/*
テストを登録する関数										   yes
データを取得 -> ディレクトリ用変数を作成 -> ディレクトリがあるか確認 ------> 処理を終了
														|
														+-->  テスト実行ページのHTMLを開き、内容を書き換える 



*/
	length, _ := strconv.Atoi(r.Header.Get("Content-Length"))
 	body := make([]byte, length)
 	length, _ = r.Body.Read(body)
   	var jsonBody map[string]string
  	json.Unmarshal(body[:length], &jsonBody)
  	//fmt.Printf("%v\n", jsonBody["part"])
	var post ContentTest
	post.Sub = jsonBody["sub"]							//科目名
	post.Day = jsonBody["day"]								//テスト日時
	//a, _ := strconv.Atoi(r.FormValue("number_of_ques"))		//問題の数
	post.Number_of_ques, _ = strconv.Atoi(jsonBody["number_of_ques"])
	//fmt.Println(post.Number_of_ques)
	post.Part = Scvtoas(jsonBody["part"])							//大問
	//fmt.Println(Part)
	post.Ques = Scvtoas(jsonBody["ques"])									//小問
	post.Answer_list = Scvtoas(jsonBody["answer_list"]) 					//解答を中に入れる
	post.Answer = Scvtoas(jsonBody["answer"])								//答えを格納
	//answer = append(answer, "ans")
	post.Point = Scvtoas(jsonBody["point"])
	dir := post.Sub + "_" + post.Day
	//printEach(answer)
	//fmt.Print(post.Point[0])
	//point = append([]string{"allocation"}, point...)							//配点
	if makeAnswer("../../mnt/data/" + post.Sub + "_" + post.Day, post.Answer, post.Point) == 0{
	//if makeAnswer("../docker-laravel/infra/docker/python/data/" + dir, post.Answer, post.Point) == 0{
		fmt.Fprintf(w, "0")
	}else{
		//makeHtml("../../opt", post.Part, post.Ques, post.Answer_list, post.Answer, post.Number_of_ques, post.Day, post.Sub)
		//file_a, _ := os.OpenFile("../docker-laravel/backend/public/data/ref2.txt", os.O_WRONLY|os.O_CREATE, 0600)
		file_a, _ := os.OpenFile("../../opt/data/ref2.txt", os.O_WRONLY|os.O_CREATE, 0600)
		defer file_a.Close()
		file_a.Truncate(0)
		file_a.WriteString(dir)
		makeHtml("../../opt/", post.Part, post.Ques, post.Answer_list, post.Answer, post.Number_of_ques, post.Day, post.Sub)
		makeHtml("../../mnt/data/" + dir + "/", post.Part, post.Ques, post.Answer_list, post.Answer, post.Number_of_ques, post.Day, post.Sub)
		fmt.Fprintf(w, "1")
	}		

	//b := r.FormValue("answer")
	//fmt.Print(b)
}
func Scvtoas(tar string) []string{
	//csvをスライスにして返却する
	//fmt.Println(tar)
	//a := tar[1:len(tar) - 1]
	return strings.Split(tar, ",")
}
func printEach(tar []string){
	//標準出力にスライスの中身すべてを表示する関数
	for i:=0; i < len(tar); i++{
		fmt.Print(tar[i])
	}
}
func makeAnswer(dir string, answer []string, point []string)int{
/*
ディレクトリがあるかどうか確認して、なければ作成、あるなら0を返却する
ディレクトリには

-------dir ---- /result       		   データを入れるようのディレクトリ
		|
		+ ----- answer.csv             解答と配点を格納する

		result.csv

		1 answer     |   解答   |  ...  
		2 allocation |   配点   |  ...      というような形式にする
*/
//fmt.Println(dir)
if f, err := os.Stat(dir); os.IsNotExist(err) || !f.IsDir() {
     err := os.Mkdir(dir, 0777)
     whetherErr(err)
     file_a, err := os.OpenFile(dir + "/answers.csv", os.O_WRONLY|os.O_CREATE, 0600)
     defer file_a.Close()
     joint := jointCs(answer) 
     file_a.WriteString("ans," + joint + "\n")
     joint = jointCs(point)
     file_a.WriteString("allocation," + joint + "\n")
     os.Mkdir(dir + "/result", 0777)
     return 1
 } else {
     return 0
 }


}
func jointCs(data []string) string{
	//文字列配列を受け取り、そのデータをcsvにする
	joint := ""
	for i := 0; i < len(data) ; i++{
		joint += data[i] + ","
	}
	return joint

}

func makeHtml(dir string, part []string, ques []string, answer_list []string, answer []string, number_of_ques int, day string, sub string){
	/*online_test_2.htmlの内容を変更する


	*/
	html, err := os.OpenFile(dir + "/online_test_2.html", os.O_CREATE|os.O_WRONLY, 0666)
	html.Truncate(0)
	whetherErr(err)
	defer html.Close()
	html.WriteString("<!DOCTYPE html>\n")
	html.WriteString("<!--pythonスクリプトによる自動的に作成したwebページ/-->\n")
	html.WriteString("<html lang=\"ja\">\n")
	html.WriteString("<head>\n")
	html.WriteString("<meta charset=\"UTF-8\">\n")
	html.WriteString("<title>オンライン試験</title>")
	html.WriteString("<script src=\"https://cdnjs.cloudflare.com/ajax/libs/html2canvas/0.4.1/html2canvas.js\"></script>\n")
	html.WriteString("<script src=\"https://code.jquery.com/jquery-3.5.1.min.js\"></script>\n")
	html.WriteString("<script src=\"./js/garlic.js\"></script>\n")
	html.WriteString("<script src=\"./js/online_test.js\"></script>\n")
	html.WriteString("<style>header{width: 100%; padding: 15px 0; margin: 0 auto; text-align: center;}</style>")
	html.WriteString("</head>\n")
	html.WriteString("<header>このシステムはフルスクリーン限定です。また、スクリーンショットなどの不正行為は禁止されています<br> <button id=\"button1\">ページを表示</button></header><body style = \"background-image : url(./img/graPink.gif);\" id = \"target\"><form>")
	html.WriteString("<h3 style=\"text-align: center\">" + sub + "</h3><br>")
	html.WriteString("<h4 style = \"text-align: right\">"+ day + "</h2>")
	html.WriteString("<form>\n")
	html.WriteString("<div id = \"test_area\" style = \"visibility:visible;\">")
	html.WriteString("<ol id = \"form_list\" data-persist = \"garlic\">\n")
	var tag int
	for i := 0; i < number_of_ques ; i++{
		tag = 0
		if part[i] != ""{
			html.WriteString(part[i] + "<br>\n")
		}
		html.WriteString(ques[i] + "<br>\n")
		if answer_list[i] == ""{
			html.WriteString("解答欄に直接入力してください。<br>")
		}else{
			html.WriteString(answer_list[i] + "<br>\n")
			tag = 1
		}
		if tag == 1{
			html.WriteString("<li><input type=\"text\" maxlength = " + "1" + "pattern= \"[1-4]\" name = \"app" + strconv.Itoa(i+1) + "\" id = \"apply" + strconv.Itoa(i+1) + "\" "+ " class = \"FORM\"></li><br>\n")
		}else{
			html.WriteString("<li><input type=\"text\" maxlength = " + strconv.Itoa(len(answer[i])) + " name = \"app" + strconv.Itoa(i+1) + "\" id = \"apply" + strconv.Itoa(i+1) + "\" " + " class = \"FORM\"></li><br>\n")
			//#解答欄に空白などが入り予期せぬ間違いを減らすため、文字の最高は答えと同じ
		}
	}
	html.WriteString("</ol>")
	html.WriteString("<input type=\"button\" onclick = \"return CHECK_FORM()\" value = \"提出\">\n")
	html.WriteString("</div>")
	html.WriteString("<script src=\"./js/expend.js\"></script>")
	html.WriteString("<script>")
	html.WriteString("const number_of_ques = $('input').length - 1;\n")
	html.WriteString("const dir = " + " \"" + dir + "data" + ";\"\n") 
	html.WriteString("setInterval(INVISIBLE_IF_FULLSCREEN(), 100);");
	html.WriteString("</script>\n")
	html.WriteString("</body>\n")
	html.WriteString("</html>\n")
}


func addUnique(w http.ResponseWriter, r *http.Request){
	/*
	Unique_numからunique_idを付与するidにデータがすでに付与されているかどうか確認する 
								|
								+---------------------------+
								|							|
								V                           |
				付与されているなら、処理を終了し0を返却す津          |
				                                            |
				                                            V
				                            unique_idを降順出力し、最初の値 + 1を付与してインサートする

	*/
	err := r.ParseForm()
	whetherErr(err)
	db := GetDBConn()
	defer db.Close()
	db.SingularTable(true)
	var unique_num Unique_num
	id, _ := strconv.Atoi(r.FormValue("id"))
	db.Where("id = ?", id).First(&unique_num)
	fmt.Print(unique_num.Id)
	if unique_num.Id == 0{
		db.Order("unique_id desc").First(&unique_num)
		unique_num.Id = id
		unique_num.Unique_id = unique_num.Unique_id + 1
		db.Create(&unique_num)
		fmt.Print(unique_num.Unique_id)
	}
	fmt.Fprintf(w, "1")
}


func checkPass(w http.ResponseWriter, r *http.Request){
	/*
	postされたID、Passを取得する
			|
			V
	table名usersから, postされたIDを検索し一致するかどうか検索する
			|
			V
	一致したなら、1を返却する。そうでない場合は0を返却する

	*/
	err := r.ParseForm()
	whetherErr(err)
	id :=  r.FormValue("id")
	pass := r.FormValue("pass")
	var users []Users
	db := GetDBConn()
	defer db.Close()
	flag := 0
	//fmt.Print(pass)
	db.Where("id = ?", id).Find(&users)
	for i:=0 ; i < len(users) ; i++{
		if(users[i].Passward == pass){
			flag = 1
			break
		}
	}
	fmt.Fprintf(w, strconv.Itoa(flag))
	}
func log_in(w http.ResponseWriter, r *http.Request){
	/*
postされたidをもとにデータベースからパスワードを取り出す。その値はsha256でハッシュにしているためpostされたpasswardをハッシュにして比較
同じ値なら1をそうでないなら0を、データが存在しないなら-1を返却する
	*/
	err := r.ParseForm()
	whetherErr(err)
	id :=  r.FormValue("id")
	pass := r.FormValue("passward")
	var users Users
	db := GetDBConn()
	defer db.Close()
	type Data struct{
		Pass string
		Flag int
	}
	var data Data
	db.Where("id = ?", id).First(&users)
	if users.Passward == ""{
		data.Flag = -1
		//fmt.Fprintf(w, "-1")
	}else{
		if string(users.Passward) == getSHA256(pass){
			data.Flag = 1
			data.Pass = getSHA256(pass)
			//fmt.Fprintf(w, string(users.Passward))
		}else{
			//fmt.Fprintf(w, "0")
			data.Flag = 0
			//fmt.Println(users.Passward)
			//fmt.Print(hex.EncodeToString(getSHA256(pass)))
		}
	}
	retJson, _ := json.Marshal(data)
	fmt.Fprintf(w, string(retJson))
}
func checkAuth(w http.ResponseWriter, r *http.Request){
	/*
type Auth struct {
	Id         int
	User_print string
}

	postされたuser_printとデータベースに存在するuser_printを比較する関数

	*/
	var auth []Auth
	err := r.ParseForm()
	whetherErr(err)	
	db := GetDBConn()
	defer db.Close()
	db.SingularTable(true)
	b := r.FormValue("data")
	flag := 0
	//fmt.Println(a)


	db.Where("id = ?", r.FormValue("id")).Find(&auth)
	for i:=0;i < len(auth);i++{
		if AtoString10(auth[i].User_print) == b{
			flag = 1
			break
		}
		//fmt.Println(AtoString10(auth[i].User_print) + "		" + b)
	}
	fmt.Fprintf(w, strconv.Itoa(flag))
}
func checkIp(w http.ResponseWriter, r *http.Request){
	//現在のipをハッシュにしてuserprintと比較する
	var auth []Auth
	err := r.ParseForm()
	whetherErr(err)	
	db := GetDBConn()
	defer db.Close()
	db.SingularTable(true)
	db.Where("id = ?", r.FormValue("id")).Find(&auth)
	flag := 0
	for i:=0;i < len(auth);i++{
	if to8hash(r.FormValue("ip")) == auth[i].User_print{
		flag = 1
		break
	}
}
	fmt.Fprintf(w, strconv.Itoa(flag))
}

func registration(w http.ResponseWriter, r *http.Request){
	/*
	Id int
	Name string
	Passward string
	Question string
	Answer string
	postされたこれらのデータを登録(passwardはハッシュに変換)

							   Yes
	dbにidが登録されているか ----+------->データを登録して1を返却する
							|  No
							+------->  処理を止めて0を返却する
	*/
	var users Users
	err := r.ParseForm()
	whetherErr(err)	
	db := GetDBConn()
	defer db.Close()
	d, _ := strconv.Atoi(r.FormValue("id"))
	db.Where("id = ?", d).First(&users)
	if users.Passward == ""{
	users.Id = d
	users.Name = r.FormValue("username")
	users.Passward = getSHA256(r.FormValue("passward"))
	//fmt.Print(users.Id)
	users.Question = r.FormValue("question")
	users.Answer = r.FormValue("answer")
	db.Create(&users)
	fmt.Fprintf(w, "1")
}else{
	fmt.Fprintf(w, "0")
}
}
func auth(w http.ResponseWriter, r *http.Request){
	//IDとハッシュ化したIPを紐図ける
	err := r.ParseForm()
	whetherErr(err)
	var auth Auth
	auth.User_print =  to8hash(r.FormValue("ip"))
	auth.Id, _ =  strconv.Atoi(r.FormValue("id"))
	db := GetDBConn()
	defer db.Close()
	db.SingularTable(true)
	db.Create(&auth)
	//fmt.Println(auth.User_print)
	//v, _ := strconv.ParseUint(auth.User_print, 16, 0)
	//fmt.Println(v)
	//b, _ := strconv.ParseInt(auth.User_print, 16, 64)
	//a := strconv.FormatInt(b, 10)
	//v := strconv.FormatInt(a, 10)
	//fmt.Println(a)
	fmt.Fprintf(w, AtoString10(auth.User_print))
}
func AtoString10(tar string) string{
	b, _ := strconv.ParseInt(tar, 16, 64)
	return strconv.FormatInt(b, 10)
}
func to8hash(ip string) string{
//8桁のハッシュ値に変換する
return getSHA256(ip)[0:8]
}

func receieveTest(w http.ResponseWriter, r *http.Request) {
	/*受け取るデータはカンマに既に分けられているデータ 例えば
	  {"data": "1, 2, sd", "url", 201803413}
	  のような形式でデータが送られてくる
	  またクエリにid=?&test_title=?  と、学籍番号を付与して送られてくる

	  データをまず取得
	  	↓
	  所定のurlにデータを書き込む(テスト結果の閲覧やMYSQLを使えない人を考慮してcsv形式で出力する)
	  csvは先頭に学籍番号を付与してデータをそのままcsvに出力
	*/
	/*___________________________________________  データの取得を行う

	  var error int
	  error = 1//err判定用

	      if err := r.ParseForm(); err != nil {
	          whetherErr(err, &error)
	      }
	  	var test Test
	      data :=  r.FormValue("data")
	      //fmt.Printf("%v\n", test)
	      var id string
	      id = r.URL.Query().Get("id")
	      var test_title string
	      test_title = r.URL.Query().Get("test_title")
	      //data = "1,2,3"          //for debug
	      test.Data = strings.Split(id + data, ",")//仕様通りにデータの結合を配列にして行う
	*/ //_____________________________________________
	var test Test
	cap := dataSet(w, r)
	id := cap[0]
	data := cap[1]
	test_title := cap[2]
	test.Data = strings.Split(id+data, ",")
	//fmt.Print(test.Data[1])
	f, err := os.OpenFile("../../mnt/data/"+test_title+"/result/"+id+".csv", os.O_WRONLY|os.O_CREATE, 0600)
	//f, err := os.OpenFile("../docker-laravel/infra/docker/python/data/"+test_title+"/result/"+id+".csv", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	whetherErr(err)
	f.Truncate(0) // ファイルの中身をからにする。
	//whetherErr(err, &error)
	//	 fw := os.NewWriter(f)
	for i := 0; i < len(test.Data); i++ {
		_, err := f.Write([]byte(test.Data[i] + ","))
		whetherErr(err)
	}
	fmt.Fprintf(w, "1")
}

func dataSet(w http.ResponseWriter, r *http.Request) []string {

	err := r.ParseForm()
	whetherErr(err)

	data := r.FormValue("data")
	//fmt.Printf("%v\n", test)
	var id string
	id = r.URL.Query().Get("id")
	var test_title string
	//var test Test
	test_title = r.URL.Query().Get("test_title")
	//data = "1,2,3"          //for debug
	//test.Data = strings.Split(id + data, ",")//仕様通りにデータの結合を配列にして行う
	cap := []string{id, data, test_title}
	return cap
}
func integration(w http.ResponseWriter, r *http.Request) { //保存されているファイルのデータと同じか確認し、同じでないなら再度receivetestを行う。
	/*
	   	var error int
	   	error = 1//err判定用

	       if err := r.ParseForm(); err != nil {
	           whetherErr(err, &error)
	       }
	       data :=  r.FormValue("data")
	       //fmt.Printf("%v\n", test)
	       var id string
	       id = r.URL.Query().Get("id")
	       var test_title string
	       test_title = r.URL.Query().Get("test_title")
	       //data = "1,2,3"          //for debug
	       test.Data = strings.Split(id + data, ",")//仕様通りにデータの結合を配列にして行う

	*/
	cap := dataSet(w, r)
	id := cap[0]
	data := cap[1]
	var test Test
	test_title := cap[2]
	//test.Data = cap[3]
	test.Data = strings.Split(id+data, ",")
	//fmt.Print(test_title)
	//f, err := os.Open("../python/data/" + test_title + "/result/" + id + ".csv")
	f, err := os.Open("../../mnt/data/" + test_title + "/result/" + id + ".csv")
	//f, err := os.Open("../docker-laravel/infra/docker/python/data/"+test_title+"/result/"+id+".csv")
	defer f.Close()
	whetherErr(err)
	var d string
	for i := 0; i < len(test.Data); i++ {
		d = d + test.Data[i] + ","
		whetherErr(err)
	}
	b, err := ioutil.ReadAll(f)
	/*
	   for i:=0; i < len(d);i++{
	   	fmt.Print(string(d[i]))
	   	fmt.Print("1      2")
	   	fmt.Print(string(b[i]))
	   if d[i] == b[i]{
	   }else{
	   	error = 0
	   }
	   }*/
	//fmt.Printf("結合データ%s、ファイルに保存されているデータ%s", d, string(b))
	if d == string(b) {
		fmt.Fprintf(w, "1")
	} else {
		fmt.Fprintf(w, "0")
	}

}

func addScore(w http.ResponseWriter, r *http.Request) {
	/*
	   postされたテスト名のテストすべてを採点する関数
	   採点したいテスト名をpost
	   	↓
	   受け取る
	   	↓
	   テスト名/result/にユーザーのテストが格納されており。それとanswer.csvを比較
	   正解ならpointに点数を追加する。
	   	↓
	   result.csvに[id, point]という形で書き込む
	   	↓
	   同時にテストの履歴をDBに書き込む
	   	↓
	   終了

	*/
	err := r.ParseForm()
	whetherErr(err)

	test_title := r.FormValue("test_title")
	//test_title := "1_1"
	answer_path := "../../mnt/data/" + test_title + "/answers.csv" //ここに答えと配点が格納されている
	result_path := "../../mnt/data/" + test_title + "/result.csv"  //ここに結果すべてが保存される
	answer, allocation := getCsv(answer_path)
	//fmt.Print(allocation)
	//result_csv, _ := getCsv(result_path)

	result, err := os.OpenFile(result_path, os.O_CREATE|os.O_WRONLY, 0666)
	whetherErr(err)
	defer result.Close()
	dir := "../../mnt/data/" + test_title + "/result/"
	dir_list := getAllFilenames(dir)
	//fmt.Print(dir_list)
	for i := 0; i < len(dir_list); i++ {
		//fmt.Print(dir_list[i])
		target, _ := getCsv(dir_list[i]) //target[0]には学籍番号が入っている
		//fmt.Printf("%v \n", target)
		point := scoringTest(target, answer, allocation)
		//fmt.Print(point)
		result.WriteString(target[0] + "," + strconv.Itoa(point) + "\n")
		addHistory(target[0], point, test_title)
	}
	ret := strconv.Itoa(len(dir_list))
	fmt.Fprintf(w, ret)

}
func addHistory(id string, point int, test_title string) {
	/*
	   idにtest_titleと点数を書き込む
	*/
	db := GetDBConn()
	defer db.Close()
	var histories History
	histories.Id = id
	histories.Test_title = test_title
	histories.Point = point
	db.Create(&histories)

}

func openCsv(csvpath string) []string {

	/*
	   csvのpathを受け取ってオープンする。また、中身を返す関数(ver2)
	*/
	f, err := os.Open(csvpath)
	if err != nil {
		whetherErr(err)
	}
	defer f.Close()
	reader := csv.NewReader(f)
	var line []string
	for {
		line, err = reader.Read()
		if err != nil {
			break
		}
	}
	return line

}

func scoringTest(target []string, answer []string, allocation []string) int {
	/*3つのcsvを受け取り、2つのcsv比較し一致しないなら何もしない、一致したならばその要素位置のallocationをpoint
	  に可算する。そしてそれを返す関数
	*/
	point := 0
	//fmt.Print(len(target))
	for i := 1; i < (len(target) - 1); i++ { //1つ目のデータは学籍番号など結果に影響しないものが入っているため読み飛ばす 一番最後の値が,であるため-1している
		//fmt.Printf("target:%s\nanswer:%s\nallocation:%s \n", target[i], answer[i], allocation[i])
		if target[i] == answer[i] {
			allocation_int, _ := strconv.Atoi(allocation[i])
			//fmt.Print(allocation_int)
			point += allocation_int
			//fmt.Printf("target:%s\nanswer:%s\nallocation:%s \n", target[i], answer[i], allocation[i])
		}
		//fmt.Printf("target:%s\nanswer:%s\nallocation:%s \n", target[i], answer[i], allocation[i])
	}
	return point
}

func getCsv(fname string) ([]string, []string) {
	//csvのファイル名を受け取りその中身を返す関数
	fmt.Printf(fname)
	f, err := os.Open(fname)
	whetherErr(err)

	r := csv.NewReader(f)
	content, err := r.Read()
	//[0]に生徒ならば結果が格納されている。answerファイルなら[0]に解答,[1]に配点が書かれている。
	whetherErr(err)

	content2, err := r.Read()
	if err != io.EOF {
		return content, content2
	} else {
		return content, nil
	}
}

func getAllFilenames(dir string) []string {
	//指定したディレクトリ内のresultディレクトリに格納されたファイル名をすべてスライスの中に入れて返す(csv限定)
	//dir = "../../mnt/data/" + testname + "/result"
	f, err := ioutil.ReadDir(dir)
	whetherErr(err)

	var paths []string
	for _, file := range f {
		if file.IsDir() {
			paths = append(paths, getAllFilenames(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths

}
func getTeacherComment(w http.ResponseWriter, r *http.Request) {
	/*
	idを逆順にして、6件データを取得する

	*/
	db := GetDBConn()
	defer db.Close()
	var comments []Comment
	db.Order("id desc").Limit(6).Find(&comments)
	a, _ := json.Marshal(comments)
	fmt.Fprint(w, string(a))

}

func insertTeacherComment(w http.ResponseWriter, r *http.Request) {
	//受け取ったデータに最大のid+1を付与して、返却する
	db := GetDBConn()
	defer db.Close()
	var comments Comment
	db.Last(&comments)
	comments.Id += 1
	err := r.ParseForm()
	whetherErr(err)
	comments.Contents = r.FormValue("content")
	db.Save(&comments)
	fmt.Fprintf(w, "1")
}

func updateHow_many(w http.ResponseWriter, r *http.Request) {
	//bbsのいくつグループを作るか の参照値を変更する。
	//また変更時にhow_manyとunique_numの中身のデータの全削除を行う
	var bbs Bb
	db := GetDBConn()
	defer db.Close()
	var unique_num Unique_num
	db.Find(&bbs)
	db.Delete(&bbs)
	db.SingularTable(true)
	var how_many How_many
	err := r.ParseForm()
	whetherErr(err)
	db.Delete(&how_many)
	db.Delete(&unique_num)
	how_many.Num_of_groups, _ = strconv.Atoi(r.FormValue("num"))
	db.Save(&how_many)
	fmt.Fprintf(w, "1")

}

func getAlldir(w http.ResponseWriter, r *http.Request) {
	//テスト結果が格納されているディレクトリの名前を返す関数
	dir := "../../mnt/data"
	f, err := ioutil.ReadDir(dir)
	whetherErr(err)
	test_list := []string{}
	for i := 0; i < len(f); i++ {
		if f[i].IsDir() == true {
			test_list = append(test_list, f[i].Name())
		}
	}

	a, _ := json.Marshal(test_list)
	fmt.Fprint(w, string(a))

}
func getHistories(w http.ResponseWriter, r *http.Request) {
	//idを受け取って、受け取ったidのテストの履歴を逆順に並べてjson順に並べて返却する。
	/*
	postされたidを取得
		|
		V
	idからデータを全取得する、インサートされた順にデータは並んでいるため、逆順にする
		|
		V
	jsonにして値を返却
	*/
	db := GetDBConn()
	defer db.Close()
	err := r.ParseForm()
	whetherErr(err)

	id := r.FormValue("id")
	var histories []History
	db.Where("id = ?", id).Find(&histories)
	var a []History
	for i := (len(histories) - 1); i > 0; i = i - 1 {
		a = append(a, histories[i])
	} //aを逆順ソート
	retJson, _ := json.Marshal(a)
	fmt.Fprintf(w, string(retJson))

}

func dec(w http.ResponseWriter, r *http.Request) { //クエリのハッシュをidに戻し所属するグループの割り出しを行う
	/*
	クエリよりidを取得し、そのidからunique_idを取得 
				|
				V
	ハッシュ値をパラメーターより取得し、グループとハッシュ値を返却する
	*/
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	db := GetDBConn()
	db.SingularTable(true) //データベースは自動的に複数形になるが、単数形のデータベースがあるため、無効化している
	defer db.Close()
	var auth Auth
	var how_many How_many
	var unique_num Unique_num
	db.First(&how_many) //grupの数を取得
	//fmt.Print(how_many.Num_of_groups)
	id := r.URL.Query().Get("id")
	db.Where("id = ?", id).First(&unique_num)                          //idが一致するデータを取得
	your_groups := (unique_num.Unique_id % how_many.Num_of_groups) + 1 //返し値の1つを取得。これは、ページにアクセスした人間のグループを返す
	//your_groups := 1
	hash, _ := strconv.Atoi(r.URL.Query().Get("hash")) //16進数にデコード
	dehash := fmt.Sprintf("%02x", hash)
	db.Where("user_print = ?", dehash).Where("id = ?", id).First(&auth)
	if auth.Id != unique_num.Id { //ユニークなIdが不正な値の時、0を返して不正な値が検出されたことを知らせる
		fmt.Fprintf(w, "0")
		//fmt.Print("error")
	} else {
		data := retData{strconv.Itoa(your_groups), strconv.Itoa(hash)} //stringじゃないとJsonにデコードできないため
		retJson, _ := json.Marshal(data)
		fmt.Fprintf(w, string(retJson))
		//print(retJson)
	}

}

func fetchAriticle(w http.ResponseWriter, r *http.Request) { //一覧データと取得するデータの数か記載
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	//var bb Bb
	var bbs []Bb
	db := GetDBConn()
	defer db.Close()
	get_above := r.URL.Query().Get("above") //取得してくるデータが上から何番目以上かデータを受け取る
	//get_above := 0								//テストのため0 (全件取得) 例えばもしこれが4ならば4以降のデータを取得してくる
	group := r.URL.Query().Get("group") //このページにアクセスしたユーザーがなんのグループに所属しているか取得する
	//group := 2								//テスト用
	//db.Find(&bbs)
	db.Where("group_id = ?", group).Offset(get_above).Find(&bbs)
	//fmt.Println(bbs)
	//*abode
	//print(len(bbs))
	//bbs, _ := json.Marshal(bbs)
	//fmt.Print(bbs)
	//var retData = &RetData{num: get_above, data: bbs}
	//fmt.Print(*retData)
	//retJson, _ := json.Marshal(*retData)
	//retJson, _ := json.Unmarshal(*retData)
	//fmt.Print(string(retJson))
	retJson, _ := json.Marshal(bbs)
	fmt.Fprintf(w, string(retJson))
}
func fetchAriticleAll(w http.ResponseWriter, r *http.Request) { //一覧データと取得するデータの数か記載
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	//var bb Bb
	var bbs []Bb
	db := GetDBConn()
	defer db.Close()
	get_above := r.URL.Query().Get("above") //取得してくるデータが上から何番目以上かデータを受け取る
	//get_above := 0								//テストのため0 (全件取得) 例えばもしこれが4ならば4以降のデータを取得してくる
	//group := 2								//テスト用
	//db.Find(&bbs)
	db.Offset(get_above).Find(&bbs)
	//fmt.Println(bbs)
	//*abode
	//print(len(bbs))
	//bbs, _ := json.Marshal(bbs)
	//fmt.Print(bbs)
	//var retData = &RetData{num: get_above, data: bbs}
	//fmt.Print(*retData)
	//retJson, _ := json.Marshal(*retData)
	//retJson, _ := json.Unmarshal(*retData)
	//fmt.Print(string(retJson))
	retJson, _ := json.Marshal(bbs)
	fmt.Fprintf(w, string(retJson))
}
func addAriticle(w http.ResponseWriter, r *http.Request) { //bbsにデータを追加
	db := GetDBConn()
	defer db.Close()
	//bbs := r.PostFormValue("data")

	var bbs Bb
	err := r.ParseForm()
	whetherErr(err)

	bbs.Id = r.FormValue("id")
	bbs.Group_id, _ = strconv.Atoi(r.FormValue("group_id"))
	bbs.Contents = r.FormValue("contents")

	db.Create(&bbs)
}

func GetDBConn() *gorm.DB {
	db, err := gorm.Open(GetDBConfig())
	//エラーがあるとreturn の2つ目にデータが返却されている
	whetherErr(err)

	db.LogMode(true)
	return db
}

func GetDBConfig() (string, string) {

	DBMS := "mysql"
	USER := "phper"
	PASS := "secret"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "laravel_local"
	OPTION := "charset=utf8&parseTime=True&loc=Local"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?" + OPTION

	return DBMS, CONNECT
}

func main() {
	handleRequests()
}
