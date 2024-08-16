text = "Order: pizza, fries, milkshake"
codes = {"pizza": "1", "fries": "2", "milkshake": "3"}

encoded = text

for value, code in codes.items():
    encoded = encoded.replace(value, code)
    
print(encoded)