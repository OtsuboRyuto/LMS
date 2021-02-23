#!/usr/bin/python3.9
print('Content-Type: text/html\n')
"""
ハッシュ値と現在のipを比べそれが不正な値であるかチェックする関数
"""
import MySQLdb, sys, json, hash_m
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

#b = cursor.execute("select exists(select id from users where id = %s)", [receive["id"]])
#print(cursor.fetchone()[0])
receive["ip"] = hash_m.TO8HASH(receive["ip"])								   #現在のipをハッシュにして格納
cursor.execute("select user_print from auth where id = %(id)s", receive)	   #データベースに格納されている、ipハッシュ
data = cursor.fetchone()[0]
#print(receive["ip"])
if data == receive["ip"]:
	print(1)																   #値が一致
else:
	print(0)

#1の時パスワードは正常, 0の時パスワードエラー, 2の時登録が重複している																	#connectionを保存する
connection.close()	