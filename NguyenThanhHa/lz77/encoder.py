def element_in_array(check_elements, elements):
    i = 0
    offset = 0
    for element in elements:
        if(len(check_elements) <= offset):
            # All of the elements in check_elements are in elements
            return i - len(check_elements)
        
        if check_elements[offset] == element:
            offset += 1
        else:
            offset = 0
        
        i += 1
    return -1
        
    
text = "SAM SAM"

encoding = "utf-8"

text_bytes = text.encode(encoding)

search_buffer = [] # Array of integers, representing bytes
check_characters = [] # Array of integers, representing bytes

i = 0


for char in text_bytes:
    check_characters.append(char)
    index = element_in_array(check_characters, search_buffer) # The index where the characters appears in our search buffer
    
    if index == -1 or i == len(text_bytes) - 1:
        if len(check_characters) > 1:
            index = element_in_array([check_characters[0]], search_buffer)
            offset = i - index - len(check_characters) + 1 # Calculate the relative offset
            length = len(check_characters) # Set the length of the token (how many character it represents)
        
            print(f"<{offset},{length}>") # Build and print our token
        else:    
            print(bytes([char]).decode(encoding)) # Print the character
        
        check_characters = []
        
    search_buffer.append(char) # Add the character to our search buffer
    
    i += 1
