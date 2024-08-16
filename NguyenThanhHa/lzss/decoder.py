encoding = "utf-8"

inside_token = False
scanning_offset = True

length = [] # Length number encoded as bytes
offset = [] # Offset number encoded as bytes

def decode(text):
    text_bytes = text.encode(encoding) # The text encoded as bytes
    output = [] # The output characters
    
    for char in text_bytes:
        if char == "<".encode(encoding)[0]:
            inside_token = True # We're now inside a token
            scanning_offset = True # We're now looking for the length number
            # print("Found opening of a token")
        elif char == ",".encode(encoding)[0]:
            scanning_offset = False
        elif char == ">".encode(encoding)[0] and inside_token:
            inside_token = False # We're no longer inside a token
            
            # Convert length and offsets to an integer
            length_num = int(bytes(length).decode(encoding))
            offset_num = int(bytes(offset).decode(encoding))
            # print("Found closing of a token")
            
            print(f"Found token with length: {length_num}, offset: {offset_num}")
            
            # Reset length and offset
            length, offset = [], []
        elif inside_token:
            if scanning_offset:
                offset.append(char)
            else:
                length.append(char)
            
        output.append(char) # Add the character to our output
        
    return bytes(output)

if __name__ == "__main__":
    print(decode("supercalifragilisticexpialidocious <35,34>").decode(encoding))