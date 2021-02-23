#!/usr/bin/python3.9
print('Content-Type: text/html\n')
"""
データをmysqlに登録するプログラム
"""
import MySQLdb, sys
import hashlib
a = sys.stdin.readline()
#a = open("data.json")
receive = json.loads(a)

connection = MySQLdb.connect(
    host='localhost',
    user='root',
    passwd='',
    db='onlinetest')
cursor = connection.cursor()

cusor.execute("insert into users values " + receive["id"] + ',' + hashlib.sha256(receive["passward"].encode()).hexdigest() + "," + receive["qusetion"] + ',' + hashlib.sha256(receive["answer"].encode()).hexdigest() + ')')


connection.commit()																				#connectionを保存する

connection.close()																				#connectionをクローズ