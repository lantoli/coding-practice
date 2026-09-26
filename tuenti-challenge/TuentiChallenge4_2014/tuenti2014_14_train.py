#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 14 - Train Empire

In any child state any train can move to any allowed station (if enough fuel and the station is connected)
and can take any of the wagons in the station (if there is any) or it can leave without any wagon.
Wagons in their destinations do not need to be moved.
We prune similar visited states to reduce search space (if we go to a state with same trains and wagon locations
as in a previous state we can ignore it because we're doing breadth-first search so we know previous state
fuel consumptions are better for those locations).
"""

from collections import namedtuple, deque
from math import sqrt


Coord = namedtuple('Coord', ['x', 'y'])
StationInfo = namedtuple('StationInfo', ['coord', 'destination', 'value'])
TrainInfo = namedtuple('TrainInfo', ['station_index', 'fuel'])
State = namedtuple('State', ['trains_info', 'wagons_station_index'])


def solve():
    stations, routes, starting_fuel = get_ints()
    station_lines = [get_strs() for _ in range(stations)]
    route_lines = [get_strs() for _ in range(routes)]

    stations_index = { info[0]: idx for idx, info in enumerate(station_lines) }

    stations_info = [StationInfo(coord= Coord(*map(int,station_line[1].split(','))),
                                 destination= stations_index[station_line[2]],
                                 value= int(station_line[3]))
                     for station_line in station_lines]

    starting_stations = [stations_index[route_line[0]] for route_line in route_lines]

    connections = [dict() for _ in range(routes)]
    for idx, route_line in enumerate(route_lines):
        for route in route_line[1:]:
            a, b = [stations_index[part] for part in route.split('-')]
            if a not in connections[idx]:
                connections[idx][a] = set()
            if b not in connections[idx]:
                connections[idx][b] = set()
            connections[idx][a].add(b)
            connections[idx][b].add(a)

    distances = [[0] * stations for _ in range(stations)]
    for idx_a in range(stations):
        a = stations_info[idx_a].coord
        for idx_b in range(idx_a+1, stations):
            b = stations_info[idx_b].coord
            distances[idx_a][idx_b] = distances[idx_b][idx_a] = \
                sqrt((b.y - a.y) * (b.y - a.y) + (b.x - a.x) * (b.x - a.x))

    ini_trains_info = [TrainInfo(station_index=starting_stations[route], fuel=starting_fuel) for route in range(routes)]
    ini_wagons_station_index = [idx for idx in range(stations)]
    ini = State(trains_info= ini_trains_info, wagons_station_index= ini_wagons_station_index)
    visited = set([visited_info(ini)])
    bfs = deque([ini])  # breadth-first search
    max_value = 0

    while bfs:
        state = bfs.popleft()
        wagons_station_index = state.wagons_station_index
        wagons_in_destination = [index
            for index in range(stations)
            if wagons_station_index[index] == stations_info[index].destination]
        wagons_not_in_destination = list(set(range(stations)) - set(wagons_in_destination))
        value = sum([stations_info[index].value for index in wagons_in_destination])
        max_value = max(max_value, value)

        if wagons_not_in_destination:  # only search if some wagon is not in destination yet
            for idx, train_info in enumerate(state.trains_info):
                for target_station_index in connections[idx][train_info.station_index]:
                    fuel_required = distances[train_info.station_index][target_station_index]
                    if train_info.fuel >= fuel_required:
                            new_fuel = train_info.fuel - fuel_required
                            new_trains_info = list(state.trains_info)
                            new_trains_info[idx] = TrainInfo(station_index= target_station_index, fuel= new_fuel)
                            no_wagons = State(trains_info= new_trains_info, wagons_station_index= wagons_station_index)
                            info = visited_info(no_wagons)
                            if info not in visited:
                                visited.add(info)
                                bfs.append(no_wagons)
                            for wagon_index in wagons_not_in_destination:
                                if wagons_station_index[wagon_index] == train_info.station_index:
                                    new_wagons_station_index = list(wagons_station_index)
                                    new_wagons_station_index[wagon_index] = target_station_index
                                    new_state = State(trains_info= new_trains_info,
                                                      wagons_station_index= new_wagons_station_index)
                                    info = visited_info(new_state)
                                    if info not in visited:
                                        visited.add(info)
                                        bfs.append(new_state)
    return max_value


def visited_info(state):
    trains_index = [info.station_index for info in state.trains_info]
    wagons_index = state.wagons_station_index
    return tuple(trains_index + wagons_index)


def get_line(): return input()
def get_int(): return int(get_line())
def get_strs(): return [x for x in get_line().split()]  # strings separated by blanks
def get_ints(): return [int(x) for x in get_line().split(',')] # ints separated by comma


if __name__ == '__main__':
    for case in range(get_int()):
        print(solve())
