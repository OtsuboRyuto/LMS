#!/usr/bin/python3.9
print('Content-Type: text/html\n')
"""
データをmysqlに登録するプログラム
"""

import MySQLdb
import sys
import hashlib
import json
a = sys.stdin.readline()
receive = json.loads(a)
receive["passward"] = hashlib.sha256(str(receive["passward"]).encode()).hexdigest()
connection = MySQLdb.connect(
    host='localhost',
    user='root',
    passwd='',
    db='onlinetest',
    charset='utf8'
    )
cursor = connection.cursor()
b = cursor.execute("select exists(select id from users where id = %s)", [receive["id"]])
if cursor.fetchone()[0] == 0:
	cursor.execute("insert into users (id, name, passward, question, answer) values (%(id)s, %(username)s, %(passward)s, %(question)s, %(answer)s)", receive)
	print(1)
#print("insert into users (id, username, passward, question, answer) values (\"" + str(receive["id"]) + "\",\"" + str(receive["username"]) + "\",\"" + hashlib.sha256(str(receive["passward"]).encode()).hexdigest() + "\",\"" + str(receive["question"]) + "\",\"" + hashlib.sha256(str(receive["answer"]).encode()).hexdigest() + "\",)")
else:
	print(0)

connection.commit()																				#connectionを保存する
connection.close()																				#connectionをクローズ
