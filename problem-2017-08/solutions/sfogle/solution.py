from __future__ import print_function
import argparse
from collections import namedtuple


Location = namedtuple('Location', ['letter', 'x', 'y'])


def get_colored_letters(locations, term):
    result = ''.join([l.letter for l in locations]).find(term)
    if result == -1:
        return []
    return locations[result:result+len(term)]


def print_grid_with_colors(grid, locations):
    STARTC = '\033[92m'
    ENDC = '\033[0m'
    for row_idx, row in enumerate(grid):
        for col_idx, col in enumerate(row):
            if (col_idx, row_idx) in [(l.x, l.y) for l in locations]:
                print((STARTC + col + ENDC), end='')
            else:
                print(col, end='')
        print()


def get_horizontal_matches(grid, initial_col_idx, col_idx_valid, move_col_idx):
    matches = []
    row_idx = 0
    while row_idx < len(grid):
        col_idx = initial_col_idx
        location_row = []
        while(col_idx_valid(col_idx)):
            location = Location(grid[row_idx][col_idx], col_idx, row_idx)
            location_row.append(location)
            col_idx = move_col_idx(col_idx)
        matches.append(location_row)
        row_idx += 1
    return matches


def get_vertical_matches(grid, initial_row_idx, row_idx_valid, move_row_idx):
    matches = []
    col_idx = 0
    while col_idx < len(grid[0]):
        row_idx = initial_row_idx
        location_row = []
        while row_idx_valid(row_idx):
            location = Location(grid[row_idx][col_idx], col_idx, row_idx)
            location_row.append(location)
            row_idx = move_row_idx(row_idx)
        matches.append(location_row)
        col_idx += 1
    return matches


def get_diagonal_matches(grid):
    matches = []
    # forward diagonal
    row_idx = len(grid) - 1
    while row_idx >= 0:
        col_idx = 0
        inner_row_idx = row_idx
        location_row = []
        while col_idx < len(grid[row_idx]) and inner_row_idx < len(grid):
            location = Location(grid[inner_row_idx][col_idx], col_idx, inner_row_idx)
            location_row.append(location)
            col_idx += 1
            inner_row_idx += 1
        matches.append(location_row)
        row_idx -= 1

    col_idx = 0
    while col_idx < len(grid[0]):
        row_idx = 0
        inner_col_idx = col_idx
        location_row = []
        while row_idx < len(grid) and inner_col_idx < len(grid[row_idx]):
            location = Location(grid[row_idx][inner_col_idx], inner_col_idx, row_idx)
            location_row.append(location)
            row_idx += 1
            inner_col_idx += 1
        matches.append(location_row)
        col_idx += 1

    # reverse diagonal
    row_idx = len(grid) - 1
    while row_idx >= 0:
        col_idx = len(grid[row_idx]) - 1
        inner_row_idx = row_idx
        location_row = []
        while col_idx < len(grid[row_idx]) and inner_row_idx < len(grid):
            location = Location(grid[inner_row_idx][col_idx], col_idx, inner_row_idx)
            location_row.append(location)
            col_idx -= 1
            inner_row_idx += 1
        matches.append(location_row)
        row_idx -= 1

    col_idx = len(grid[0]) - 1
    while col_idx >= 0:
        row_idx = 0
        inner_col_idx = col_idx
        location_row = []
        while row_idx < len(grid) and inner_col_idx >= 0:
            location = Location(grid[row_idx][inner_col_idx], inner_col_idx, row_idx)
            location_row.append(location)
            row_idx += 1
            inner_col_idx -= 1
        matches.append(location_row)
        col_idx -= 1

    return matches


def get_opposite_diagonal_matches(grid):
    matches = []
    # opposite forward diagonal
    row_idx = 0
    while row_idx < len(grid):
        col_idx = 0
        inner_row_idx = row_idx
        location_row = []
        while col_idx < len(grid[row_idx]) and inner_row_idx >= 0:
            location = Location(grid[inner_row_idx][col_idx], col_idx, inner_row_idx)
            location_row.append(location)
            col_idx += 1
            inner_row_idx -= 1
        matches.append(location_row)
        row_idx += 1

    col_idx = 0
    while col_idx < len(grid[0]):
        row_idx = len(grid) - 1
        inner_col_idx = col_idx
        location_row = []
        while row_idx >= 0 and inner_col_idx < len(grid[row_idx]):
            location = Location(grid[row_idx][inner_col_idx], inner_col_idx, row_idx)
            location_row.append(location)
            row_idx -= 1
            inner_col_idx += 1
        matches.append(location_row)
        col_idx += 1

    # # reverse opposite diagonal
    col_idx = 0
    while col_idx < len(grid[0]):
        row_idx = len(grid) - 1
        inner_col_idx = col_idx
        location_row = []
        while inner_col_idx >= 0 and row_idx >= 0:
            location = Location(grid[row_idx][inner_col_idx], inner_col_idx, row_idx)
            location_row.append(location)
            inner_col_idx -= 1
            row_idx -= 1
        matches.append(location_row)
        col_idx += 1

    row_idx = len(grid) - 1
    while row_idx >= 0:
        col_idx = len(grid[row_idx]) - 1
        inner_row_idx = row_idx
        location_row = []
        while inner_row_idx >= 0 and col_idx >= 0:
            location = Location(grid[inner_row_idx][col_idx], col_idx, inner_row_idx)
            location_row.append(location)
            inner_row_idx -= 1
            col_idx -= 1

        matches.append(location_row)
        row_idx -= 1

    return matches


def get_all_locations_to_search(grid, rows, columns):
    increment = lambda x: x + 1
    decrement = lambda x: x - 1
    return (
        # horizontal forwards and backwards
        get_horizontal_matches(grid, 0, lambda col: col < columns, increment) +
        get_horizontal_matches(grid, columns-1, lambda col: col >= 0, decrement) +
        # vertical forwards and backwards
        get_vertical_matches(grid, 0, lambda row: row < rows, increment) +
        get_vertical_matches(grid, rows-1, lambda row: row >= 0, decrement) +
        get_diagonal_matches(grid) +
        get_opposite_diagonal_matches(grid)
    )


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('infile', type=argparse.FileType('r'))
    f = [l.strip() for l in parser.parse_args().infile.readlines()]
    rows = int(f[0][0])
    columns = int(f[0][2])
    search_terms = f[rows+1:]
    grid = f[1:rows+1]

    colored_letters = []

    for location in get_all_locations_to_search(grid, rows, columns):
        for term in search_terms:
            colored_letters.extend(get_colored_letters(location, term))

    print_grid_with_colors(f[1:rows+1], colored_letters)
