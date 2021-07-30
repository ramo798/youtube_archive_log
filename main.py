from chat_downloader import ChatDownloader
import time
import json


def main_logic(tid):

    url = 'https://www.youtube.com/watch?v='

    chat = ChatDownloader().get_chat(url+tid)

    result = []

    t1 = time.time()

    for message in chat:
        # print(message)
        result.append(message)

    with open(tid + '.json', 'w') as fp:
        json.dump(result, fp)

    t2 = time.time()

    elapsed_time = t2-t1
    print(f"経過時間：{elapsed_time}")


if __name__ == '__main__':
    target = ['bWOUe_X9Q8I', 'aRpwmqYwPs8', 'zg7AD3xg960', 'w9njHYokH6E']

    for a in target:
        main_logic(a)
