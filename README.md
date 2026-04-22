# ARI Project

This project contains two Go programs:
- One encodes data into a custom file format  
- The other decodes it back into readable text  

It uses a CSV file to store a conversion table, which makes the encoding system fully customizable.
The data is processed through a simple compression step, making the output non-readable without the decoder.

### toAri (encoder)

- Reads input text
- Splits it into words
- Stores unique words once
- Replaces words with their index
- Converts characters using the CSV mapping
- Writes everything into a `.ari` file

Output is saved as: output.ari

### fromAri (decoder)

- Reads the `.ari` file
- Rebuilds the list of words
- Uses indexes to restore sentence order
- Converts values back into characters using the CSV mapping
- Prints the original text
