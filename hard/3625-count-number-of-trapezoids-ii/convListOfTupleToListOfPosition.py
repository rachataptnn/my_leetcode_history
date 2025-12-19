def tuples_to_js_points(tuples):
    lines = ["const points = ["]
    for x, y in tuples:
        lines.append(f"  {{x: {x}, y: {y}}},")
    lines.append("];")
    return "\n".join(lines)


# Example usage
input_data = [[-33,-9],[30,-37],[-10,-9],[61,-9],[56,-67],[36,-9],[36,100],[36,96],[-32,84],[18,34],[-10,-82]]
print(tuples_to_js_points(input_data))
