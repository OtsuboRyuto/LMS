#!/usr/bin/python3.9
"""
ver 1.02 12/20
What's new
ディレクトリを作成しファイルをテストごとに分けて保存できるようになった.
整合性チェックに永遠と成功できない場合->強制的にssを取らせるシステムを入れてそれをもとに採点する。 →データ送信機構は未実装

Improvement for problem
テストごとにHTMLを変更できるように、出力先を変更できるようにする.
	->してもいいが、しない方が使いやすい気もする
	->結局テストは同じところでやるので、ページを変えるとユーザーも混乱しやすい。
	->あって損はないかもしれないが実装するメリットがあまりない
	->ssの禁止するシステムを実装


UIを整える -> Bootstrapを導入予定
なおデータは保持されるのでページをリロードして、それをスクリーンショットしてもらうというのでもいいのでは..?
もしくは、もう一度送りなおしてもらうなど。
きちんとデバッグ。 また、どれだけのサーバーがシステムに耐えられるか実験を行う。
"""
import sys
import csv
import os
import glob

"""
def scanf():		#入力時数字がどうかをチェックして数字ならint型にそうでないならstr
	which = ""
	which = input()
		if str.isdigit(which) == True:
			which = int(which)
	return which
"""
def usage():
	print("########################################")
	print("#Usage                                 #")
	print("#問題に画像を使いたい時は直接、HTMLとタグを入力#")
	print("#してください。注意タグですがプログラム上では「\"」を #")
	print("#扱うことができません。ですので直接\\\" と入力し #")
	print("#てください。                              #")
	print("#空白はすべて読み飛ばします。空白が必要な問題は #")
	print("#作らないでください。                         #")
	ptint("########################################")


def main():
	part = []		
	number_of_part = []						#大問用のリスト
	ques = []								#小問用のリスト
	answer_list = []						#回答のリスト
	answer = []								#答え
	print("テスト科目を入力してください。")
	sub = ""
	sub = input()
	print("テスト実施日を入力してください。")							
	buf = ""
	day = ""
	day = input()
	print("問題の数はいくつですか?")
	number_of_ques = int(input()) 				#問題数を格納
	dir = "../" + sub + "_" + day						#格納するディレクトリ 前の階層に戻ってsub+dayのディレクトリを作成
	for i in range(number_of_ques):
		print("大問を入力してください。ない場合は空白のままエンターを押してください。")
		part.append(input())
		#print("従属する小問の問題数を入力してください。ない場合は空白のままエンターを押してください。")
		#number_of_part.append(input())
		print("小問を入力してください。")
		ques.append(input())
		print("解答のリストを入力してください。(直接文字を打ち込ませる問題は入力しなくて大丈夫です。)")
		answer_list.append(input())
		print("答えを入力してください。")
		answer.append(input())
	MAKE_HTML(part, ques, answer_list, answer, number_of_ques, day, sub, dir)
	MAKE_ANSWER(answer, dir)
def MAKE_ANSWER(answer, dir):
	os.mkdir(dir)										#専用のディレクトリを生成
	os.mkdir(dir + "/result")						#結果格納用のディレクトリを生成
	f = open(dir + "/answers.csv", 'w')
	f.write("100,")									#学籍番号が最初に入るため適当な値を入れている
	for i in answer:
		f.write(i + ",")
	f.close()

def MAKE_HTML(part, ques, answer_list, answer, number_of_ques, day, sub, dir):
	f = open("../online_test_2.html", 'w')
	f.write("<!DOCTYPE html>\n")
	f.write("<!--pythonスクリプトによる自動的に作成したwebページ/-->\n")
	f.write("<html lang=\"ja\">\n")
	f.write("<head>\n")
	f.write("<meta charset=\"UTF-8\">\n")
	f.write("<title>オンライン試験</title>")
	f.write("<script src=\"https://cdnjs.cloudflare.com/ajax/libs/html2canvas/0.4.1/html2canvas.js\"></script>\n")
	f.write("<script src=\"https://code.jquery.com/jquery-3.5.1.min.js\"></script>\n")
	f.write("<script src=\"garlic.js\"></script>\n")
	f.write("<script src=\"online_test.js\"></script>\n")
	f.write("<style>header{width: 100%; padding: 15px 0; margin: 0 auto; text-align: center;}</style>")
	f.write("</head>\n")
	f.write("<header>このシステムはフルスクリーン限定です。また、スクリーンショットなどの不正行為は禁止されています</header>")
	f.write("<body style = \"background-image : url(graPink.gif);\" id = \"target\">")
	f.write("<h3 style=\"text-align: center\">" + sub + "</h3><br>")
	f.write("<h4 style = \"text-align: right\">"+ day + "</h2>")
	f.write("<form>\n")
	f.write("<div id = \"test_area\" style = \"visibility:visible;\">")
	f.write("<ol id = \"form_list\" data-persist = \"garlic\">\n")
	for i in range(number_of_ques):
		tag = 0
		length = len(answer)
		if part[i] != "":
			f.write(part[i] + "<br>\n")
		f.write(ques[i] + "<br>\n")
		if answer_list[i] == "":
			f.write("解答欄に直接入力してください。<br>")								
		else:
			f.write(answer_list[i] + "<br>\n")
			tag = 1
		if tag == 1:
			f.write("<li><input type=\"text\" maxlength = " + "1 " + "pattern= \'[1-4]\' name = \"app" + str(i+1) + "\" id = \"apply" + str(i+1) + "\" "+ " class = \"FORM\"></li><br>\n")
		else :
			f.write("<li><input type=\"text\" maxlength = " + str(len(answer[i])) + " name = \"app" + str(i+1) + "\" id = \"apply" + str(i+1) + "\" " + " class = \"FORM\"></li><br>\n") #解答欄に空白などが入り予期せぬ間違いを減らすため、文字の最高は答えと同じ

	f.write("</ol>")
	f.write("<input type=\"button\" onclick = \"return CHECK_FORM()\" value = \"提出\">\n")
	#f.write("<script type=\"text/javascript\" src=\"./footerFixed.js\">")
	#f.write("<input type = \"button\" download = \"result\">")
	#f.write("</script>")
	f.write("</div>")
	f.write("<script>")
	f.write("//var FormElement = document.getElementByClassName(\"FORM\");\n")
	f.write("//const number_of_ques = FormElement.childElementCount;\n")
	f.write("const number_of_ques = $('input').length - 1;\n")
	f.write("const dir = " + " \"" + dir + "\"\n") 
	f.write("FORBIT_PRT();")
	f.write("setInterval(INVISIBLE_IF_FULLSCREEN(), 100)");
	f.write("</script>\n")
	f.write("</body>\n")
	f.write("</html>\n")
	f.close()

def main2():
	"""
	テストを開始終了を宣言するプログラム
	"""
	print("テストを開始しますか？ yesなら1, noなら0")
	which = input()
	if str.isdigit(which) == True:
		which = int(which)
	while True:
		if which == 0 or which == 1:
			break
		print("1か0以外が入力されました。")
		print("テストを開始しますか？ yesなら1, noなら0")
		which = input()
		if str.isdigit(which) == True:
			which = int(which)
	if which == 0:
		print("プログラムを終了します。")
		sys.exit()
	f = open("../data/ref1.txt", 'w')
	f.write("true")
	f.close()
	print("テストを開始しました。")
	print("テストを終了するとき1を入力してください。")
	which = input()
	if str.isdigit(which) == True:
		which = int(which)
	while which != 1:
		print("1以外が入力されました。終了する際は１を入力してください")
		which = input()
		if str.isdigit(which) == True:
			which = int(which)
	print("プログラムを終了します。")
	f = open("../data/ref1.txt", 'w')					#ref1をfalseにして終了
	f.write("false")
	f.close()

def main3():
	#DIR = './result'									#ファイルの数を調べたいディレクトリ(ここでは出力データの格納先)
	#number_of_dir = (sum(os.path.isfile(os.path.join(DIR, name)) for name in os.listdir(DIR)))#ファイルの数を調べるプログラム
	#for i in range(number_of_dir):
	answers = []
	match = 0	
	print("採点したいファイルの格納先を入力してください。")			
	dir = input()							#この中に正解の数を代入
	with open("../" + dir + "/answers.csv", 'r') as f:
		answer = csv.reader(f)
		answers = [row for row in answer]				#この中に回答が入ってい
	for name in glob.glob("../" + dir + '/result/*'):					#ファイルの片っ端からnameに入れる
		with open(name) as a:
			result = csv.reader(a)
			results = [row for row in result]	#生徒の解答を代入
		for i in range(len(answers[0])):
			if answers[0][i] == results[0][i]:
				match += 1								#答えが同じならマッチの数をインクリメント
		f = open("../" + dir + "/result.csv", 'a')
		f.write(str(results[0][0]) + "," + str(match - 1) + "\n")
		match = 0
		f.close()
	print("採点が完了しました。")



which = 0
print("テストを作成するなら0を、テストを開始するなら1、採点する際は2を入力してください。")
which = input()
if str.isdigit(which) == True:
	which = int(which)
while True:
	if which == 0 or which == 1 or which == 2:
		break
	
	print(which)
	print("0, 1, 2以外が入力されました。")
	print("テストを作成するなら0を、テストを開始するなら1、採点する際は2を入力してください。")
	which = input()
	if str.isdigit(which) == True:
		which = int(which)
if which == 0:
	main()
elif which == 1:
	main2()
else:
	main3()