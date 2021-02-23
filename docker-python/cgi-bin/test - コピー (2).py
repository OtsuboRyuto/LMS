#!/usr/bin/python3.9
print('Content-Type: text/html\n')
"""
IDとハッシュ化したIPを紐づける
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
receive["ip"] = hash_m.TO8HASH(receive["ip"])
cursor.execute("insert into auth(id, user_print) values (%(id)s, %(ip)s)", receive)
#print(receive["ip"])
receive["ip"] = "0x" + receive["ip"]#16進数と認識させるためフラグを追加
print(int(receive["ip"], 16))#strだとデータを送れないため、int, 10進数に値を変換してデータを渡す

connection.commit()																				#connectionを保存する
connection.close()	