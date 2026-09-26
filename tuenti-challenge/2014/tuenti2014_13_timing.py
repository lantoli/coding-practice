#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 13 - Tuenti Timing Auth

With debug=1 we get timing for encryption algorithm time, it increases when we're close to the key,
the more matches in key positions the more time it takes.
"""

import http.client, urllib.parse

numbers = '01234567890abcdef'
STR_FOUND = 'Correct key'
STR_TIME = 'Total run: '
STR_TIME_END = '\\n-->'


def main():
    inp = input()
    key = ''
    while True:
        times = []
        for number in numbers:
            content = get_web_content(inp, key + number)
            if STR_FOUND in content:
                print(key + number)
                return
            else:
                pos = content.index(STR_TIME)
                pos_end = content.index(STR_TIME_END)
                if 0 <= pos < pos_end:
                    times.append(float(content[pos+len(STR_TIME):pos_end]))
                else:
                    raise ValueError('time not found')
        max_time = max(times)
        key += numbers[times.index(max_time)]


def get_web_content(inp, key):
    conn = http.client.HTTPConnection('54.83.207.90:4242')
    conn.request("POST", "/?debug=1", urllib.parse.urlencode({'input': inp, 'key': key}))
    response = conn.getresponse()
    data = response.read()
    conn.close()
    return str(data)


if __name__ == '__main__':
    main()
