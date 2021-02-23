#!/usr/bin/python3.9
print('Content-Type: text/html\n')
"""
mysqlからデータを取り出してログインを行うシステム。
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
if cursor.execute("select exists(select id from users where id = %s);", [receive["id"]]) == 1:
	id = cursor.execute("select id from users where id = %s", [receive["id"]])
	passward = cursor.execute("select passward from users where id = %s", [receive["id"]])
	passward = cursor.fetchall()


#print(passward[0][0], hashlib.sha256(str(receive["passward"]).encode()).hexdigest())
if len(passward) == 1:
	if passward[0][0] == hashlib.sha256(str(receive["passward"]).encode()).hexdigest():
		print("1")
	else:
		print("0")
else:
	print("2")
#1の時パスワードは正常, 0の時パスワードエラー, 2の時登録が重複している
connection.commit()																				#connectionを保存する
connection.close()	