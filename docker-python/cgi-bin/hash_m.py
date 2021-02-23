import hashlib
def TO8HASH(address):
	"""
	受け取った値を8桁の16進数に変換
	"""
	dummy = hashlib.sha256(str(address).encode()).hexdigest() #16進数
	address = ""
	for i in range(8):
		address = address + dummy[i]
	return address