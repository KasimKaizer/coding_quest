def heatshield(data: str) -> int:
    """Return the number of points not covered by rectangles.

    Day 7.

    This uses a linesweep approach. By considering the sorted y-values of tiles,
    we can compute at which row another tile comes into play or is removed from play.
    Tiles can be "added" or "removed" in a presorted order.

    If no new tile applies to any given row, the row has the same exposure as
    the previous row.

    If we know which tiles are in play for a row, we can sort tiles by x-value and
    walk the row by jumping from row-start to tile-start to tile-end.
    """
    squares = set()
    for line in data.splitlines():
        xstart, ystart, width, height = [int(i) for i in line.split()]
        # Covert width/height to end column/row.
        squares.add((xstart, ystart, xstart + width, ystart + height))

    # Different grids for different inputs.
    if len(squares) == 3:
        xmax, ymax = 10, 10
    elif len(squares) == 20:
        xmax, ymax = 100, 100
    else:
        xmax, ymax = 20000, 100000

    # Sort tile starts by the start row.
    to_add = sorted(squares, key=lambda s: s[1], reverse=True)
    # Sort tile ends by the end row.
    to_remove = sorted(squares, key=lambda s: s[3], reverse=True)
    # Track which tiles are currently in play.
    in_use: set[tuple[int, int, int, int]] = set()

    total = 0
    exposed = 0
    changed = True
    for y in range(ymax):
        # Remove tiles which have gone out of play at this row.
        while to_remove and to_remove[-1][3] == y:
            changed = True
            in_use.remove(to_remove.pop())

        # Add tiles which go into play at this row.
        while to_add and to_add[-1][1] == y:
            changed = True
            in_use.add(to_add.pop())

        # If the tiles changed, update the exposure.
        if changed:
            changed = False
            exposed = 0
            cur = 0
            for square in sorted(in_use, key=lambda square: square[0]):
                xstart, _, xend, _ = square
                if cur < xstart:
                    exposed += xstart - cur
                cur = max(cur, xend)
            if cur < xmax:
                exposed += xmax - cur

        total += exposed

    return total


print(heatshield(open("real_test.txt", "r").read()))
