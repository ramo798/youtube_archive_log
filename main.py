from chat_downloader import ChatDownloader
import time
import json

url = 'https://www.youtube.com/watch?v=wh0qATKFpOo'
chat = ChatDownloader().get_chat(url,end_time="00:00:10")


result = []

t1 = time.time()


for message in chat:
    print(message)
    result.append(message)


with open('data.json', 'w') as fp:
    json.dump(result, fp)

t2 = time.time()

elapsed_time = t2-t1
print(f"経過時間：{elapsed_time}")
