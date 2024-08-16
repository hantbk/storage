codes = {"pizza": "1", "fries": "2", "milkshake": "3"}
text = "Order: 1, 2, 3"

decoded = text

for value, code in codes.items():
    decoded = decoded.replace(code, value)

print(decoded)