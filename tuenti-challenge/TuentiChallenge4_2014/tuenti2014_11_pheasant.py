#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 11 - Pheasant

There is an optimization: Users are processed sorted by newest_time desc. We keep the tentative event result ordered
by time. When we have the list with the number of events requested, then as soon as we find a user whose newest_time
 is less than the lowest in the result, then we can stop safely.
We also use memoization to cache user events to optimize repetitive access to the same user
"""

from Crypto.Cipher import AES
from itertools import product
from string import ascii_letters
from collections import namedtuple
from operator import attrgetter
from functools import lru_cache
from sys import stdin


User = namedtuple('User', ['id', 'password', 'last_time'])
Event = namedtuple('Event', ['id', 'time'])
three_letter_prods = tuple(bytes([ord(a), ord(b), ord(c)]) for (a,b,c) in product(ascii_letters, ascii_letters, ascii_letters))


def main():
    for line in stdin:
        users = []
        friends = line.split(';')
        event_count = int(friends[0])
        for friend in friends[1:]:
            user_id, password = [x.strip() for x in friend.split(',')]
            users.append(User(user_id, password, last_time(user_id)))

        events = []
        for user in sorted(users, key=attrgetter('last_time'), reverse=True):
            if len(events) == event_count and user.last_time < events[-1].time:
                break; # impossible to get any event into the result from the rest of users
            events += user_feed_data(user.id, user.password)
            events = list(sorted(events, key=attrgetter('time'), reverse=True))[:event_count]

        print(' '.join([event.id for event in events]))


@lru_cache(maxsize=300000)
def user_feed_data(user_id, key_first_part):
    with open('./feeds/encrypted/{}/{}.feed'.format(user_id[-2:], user_id), mode='rb') as file:
        encrypted_content = file.read()
    for key in possible_keys(key_first_part):
        msg_bytes = AES.new(key, AES.MODE_ECB).decrypt(encrypted_content)
        try:
            msg = msg_bytes.decode(encoding='ascii')
            if msg.startswith(user_id):
                events = []
                try:
                    for line in msg.split('\n'):
                        _, time, id = line.split()
                        events.append(Event(id, time))
                except:
                    pass
                return events
        except UnicodeDecodeError:
            pass
    raise ValueError("info not found")


def last_time(user_id):
    with open('./feeds/last_times/{}/{}.timestamp'.format(user_id[-2:], user_id), encoding='ascii') as file:
        content = file.read()
    return content


def possible_keys(key_first_part):
    key_bytes = bytes(key_first_part, encoding='ascii')
    return (key_bytes+letters for letters in three_letter_prods)


if __name__ == '__main__':
    main()


