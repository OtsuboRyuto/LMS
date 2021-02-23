#!/usr/bin/python3.9
print('Content-Type: text/html\n')
import sys
import json
#receive = "data=%2C2019-3413%2C1%2C2%2C%2C"
#データを受け取ってcsvに変更するプログラム
#receive = sys.stdin.readline()
a = sys.stdin.readline()
#a = open("data.json")
receive = json.loads(a)
def RELIEF_EXTRA(target):
	"""
	データの中の余計な文字列を除外
	"""

	data = ""
	tag = 0
	for i in range(len(target)):
		if target[i] == '=':
			tag = 1
		if tag == 1:
			data = data + target[i]
	return data

def OUT_LIST(target):
	"""
	データの区切り文字を取って、csvに変換
	"""

	data = target.split("%2C")
	return data

def main(target):
	data = target["data"]	
	ID_path = target["url"]
	target = str(data[1:])
	data = target.split(',')
	ID = data[0]
	tag = 0;
	path = ID_path + "/result/" + ID
	f = open(path + ".csv", 'w')#学籍番号.txtでファイルを保存
	f.write(target + ",")#データを保存
	f.close()


#RELIEF_EXTRA(receive)
main(receive)