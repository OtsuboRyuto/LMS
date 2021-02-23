#!/usr/bin/python3.9
print('Content-Type: text/html\n')
"""
ユーザのidに固有のidを追加する。これは最も大きな固有idに1を足した値である。
"""
import MySQLdb, sys, json
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

cursor.execute("select max(unique_id) from unique_num")
data = cursor.fetchone()[0]
if data == None:
	cursor.execute("insert into unique_num (id, unique_id) values (%(id)s,  0)", receive)
else:
	input_data = {"id": receive["id"], "input": int(data) + 1}
	cursor.execute("insert into unique_num (id, unique_id) values (%(id)s, %(input)s)", input_data)

print(1)																


#1の時パスワードは正常, 0の時パスワードエラー, 2の時登録が重複している																	#connectionを保存する
connection.commit()
connection.close()	