#!/usr/bin/python3.9
print('Content-Type: text/html\n')
"""
ハッシュ値を取得して検証するプログラム
"""

import sys
import hashlib
#初期パス QwErTyUiOp
#ハッシュ値
hash_value = "9d387a6f1bbe43d28fd234d5940c4e9d8396ac85d5218113f9249c08cd8b6e61"
passward = hashlib.sha256(sys.stdin.readline().replace('\n', '').encode()).hexdigest()
if passward == hash_value:
	print(1)
else:
	print(0)
