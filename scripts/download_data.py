import urllib.request
import datetime

with urllib.request.urlopen('http://slova.cetba.eu/generate.php?number=3500') as response:
    res = response.read().decode('utf-8')
    data = res.split('|')
    for i in range(len(data)):
        data[i] = data[i].strip()
    data = ';'.join(data).lower()
    file = open('words_' + datetime.datetime.now().strftime("%Y%m%d") + '.txt', 'w')
    file.write(data)
    file.close()