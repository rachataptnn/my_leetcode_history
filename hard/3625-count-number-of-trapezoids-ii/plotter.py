import matplotlib.pyplot as plt

# Your array of points
# points = [
#     [2, 3.4],
#     [4.22, 5.5],
#     [1.5, 2.1],
#     [3.8, 0.9]
# ]


# points = [[0,0],[1,0],[0,1],[2,1]] # case2

# points = [[-32,12],[-32,-94],[-32,-15],[-30,88]] # 146/551

points = [[71,-89],[-75,-89],[-9,11],[-24,-89],[-51,-89],[-77,-89],[42,11]]# 418/551
# points = [[-33,-9],[30,-37],[-10,-9],[61,-9],[56,-67],[36,-9],[36,100],[36,96],[-32,84],[18,34],[-10,-82]]# 482/551




# Separate x and y coordinates
x = [p[0] for p in points]
y = [p[1] for p in points]

# Plot
plt.scatter(x, y)
plt.xlabel("X")
plt.ylabel("Y")
plt.title("Points in the Plane")
plt.grid(True)

plt.show()
