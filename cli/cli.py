import requests
import sys

if __name__ == '__main__':
	if sys.argv[1] == 'task':
		if sys.argv[2] == 'push':
			r = requests.post('http://localhost:8080/tasks', data={'image' : sys.argv[3]})
		elif sys.argv[2] == 'get':
			r = requests.get('http://localhost:8080/tasks/' + sys.argv[3])
		elif sys.argv[2] == 'del':
			r = requests.delete('http://localhost:8080/tasks/' + sys.argv[3])
		elif sys.argv[2] == 'trigger':
			r = requests.post('http://localhost:8080/trigger?task=' + sys.argv[3])
		print(r.status_code)
		print(r.text)
	elif sys.argv[1] == 'log':
		#for i in range(100):
		#	print(i)
			r = requests.get('http://localhost:8080/logs/' + sys.argv[2])
			print(r.status_code)
			print(r.text)
	else:
		print('Unrecognized command ' + sys.argv[1])
