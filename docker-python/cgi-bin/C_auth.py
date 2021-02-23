#!/usr/bin/python3.9
print('Content-Type: text/html\n')
"""
IDがハッシュ値と紐づけられているか確認する
"""
import MySQLdb, sys
import hashlib, json
a = sys.stdin.readline()
receive = json.loads(a)
connection = MySQLdb.connect(
    host='localhost',
    user='root',
    passwd='',
    db='onlinetest',
     charset='utf8'
    )
cursor = connection.cursor()
receive["data"] = str(hex(int(receive["data"])))[2:]#送られてくるデータはハッシュを10進数にエンコードしたものであるため、それを16進数に変換し、0xをとりstrへエンコード
#print(receive["data"])
cursor.execute("select exists(select user_print from auth where user_print = %s)", [receive["data"]])
#cursor.execute("insert into auth(id, user_print) values (%(id)s, %(ip)s)", receive)
if cursor.fetchone()[0] == 1:
	cursor.execute("select id from auth where user_print = %s", [str(receive["data"])])#ハッシュ値が存在するかチェック
	print(int(cursor.fetchone()[0]))
else:
	print(0)


#1の時パスワードは正常, 0の時パスワードエラー, 2の時登録が重複している
connection.commit()																				#connectionを保存する
connection.close()	