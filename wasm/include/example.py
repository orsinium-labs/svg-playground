# The canvas on the right contains the SVG image from the `canvas` variable.
# Press the "draw" button to re-draw the image.
import svg
from random import choice

rows = 4
columns = 6
width = 60
height = 40
radius = 4
colors = [
    '#2ecc71',
    '#3498db',
    '#9b59b6',
    '#34495e',
    '#e67e22',
    '#e74c3c',
    '#7f8c8d',
]


def random_circles(x, y):
    yield svg.Circle(
        cx=x, cy=y, r=radius,
        fill=choice(colors),
    )
    yield svg.Circle(
        cx=x, cy=y, r=radius // 2,
        fill='none',
        stroke='white',
        stroke_width=1,
    )


circles = []
step_x = width // columns
step_y = height // rows
for x in range(step_x // 2, width, step_x):
    for y in range(step_y // 2, height, step_y):
        circles.extend(random_circles(x=x, y=y))


canvas = svg.SVG(
    viewBox=svg.ViewBoxSpec(0, 0, width, height),
    elements=circles,
)
