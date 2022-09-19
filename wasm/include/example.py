import svg


def heart_path():
    return [
        svg.M(10, 30),
        svg.Arc(20, 20, 0, False, True, x=50, y=30),
        svg.Arc(20, 20, 0, False, True, x=90, y=30),
        svg.Q(90, 60, 50, 90),
        svg.Q(10, 60, 10, 30),
        svg.Z(),
    ]


canvas = svg.SVG(
    viewBox=svg.ViewBoxSpec(-40, 0, 150, 100),
    elements=[
        svg.Path(
            d=heart_path(),
            fill="grey",
            transform=[
                svg.Rotate(-10, 50, 100),
                svg.Translate(-36, 45.5),
                svg.SkewX(40),
                svg.Scale(1, 0.5),
            ],
        ),
        svg.Path(
            d=heart_path(),
            fill="none",
            stroke="red",
        ),
    ],
)
canvas
