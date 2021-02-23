#!/usr/bin/python3.9
print('Content-Type: text/html\n')
"""
データを送信
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


cursor.execute("insert into bbs(id, group_id, contents) values (%(id)s, %(group)s, %(contents)s)", receive)


print(1)

connection.commit()																				#connectionを保存する
connection.close()	