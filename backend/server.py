from flask import Flask , request
from flask_cors import CORS, cross_origin
from ctypes import *
from numpy.ctypeslib import ndpointer
import json
import random

num2 = random.randrange(100000000, 10000000000)

app = Flask(__name__)
CORS(app, resources={r"/*": {"origins": "*"}})
app.config['CORS_HEADERS'] = 'Content-Type'

class GoSlice(Structure):
	_fields_ = [("data", POINTER(c_void_p)), 
            ("len", c_longlong), ("cap", c_longlong)]


@app.route('/eval', methods=['GET','POST'])
def parse_request():
	lib = cdll.LoadLibrary("main.so")
	# result = request.json
	# input = int(result['input'])

	lst = []
	i = 0
	while i < 5:
	    lst.append(random.randrange(100000000, 150000000))
	    i += 1
	numbers = GoSlice((c_void_p * 5)(lst[0], lst[1], lst[2], lst[3], lst[4]), 5, 5)
	lib.ExportedFunction.argtypes = [GoSlice]
	lib.ExportedFunction.restype = ndpointer(dtype = c_uint64, shape = (100,))
	ret = lib.ExportedFunction(numbers)

	newLst = list(filter(lambda a: a != 0, ret.tolist()))
	# print(newLst)
	dic = {}
	for i in lst:
	    if i not in dic:
	        dic[i] = 1

	final = []
	j = 0
	while j < len(newLst):
	    if newLst[j] in dic:
	        tmp = []
	        tmp.append(newLst[j])
	        j += 1
	        while j < len(newLst):
	            if newLst[j] not in dic:
	                tmp.append(newLst[j])
	                j += 1
	            else:
	                break
	        final.append(tmp)

	final_set = set(tuple(x) for x in final)
	final = [list(x) for x in final_set]

	# print(final)

	myJSON = '{"array": [[]]}'
	o = json.loads(myJSON)
	o["array"] = final
	newJSON = json.dumps(o)
	return newJSON	
	# return 'Received!'


if __name__ == "__main__":
    app.run(host='127.0.0.1', port=3000)


# run this go build -o main.so -buildmode=c-shared main.go && python3 server.py 