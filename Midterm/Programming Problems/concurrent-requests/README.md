# Create an http client worker

1. Create an http client that retrieves a list of numbers from the /api/numbers the result will be json array

2. Filter the numbers by the follloling rule:

    If the number is constructed using only one digit exp: [1111, 33, 9]
    **or** the sum of digits can be divided by 10 without remainder

3. Submit all "special" numbers by creating a post call to the following path /api/numbers/special. The payload should contain only **one number** converted to a string.

## Notes:
If the list of numbers (received /api/numbers) contains n number of "special" numbers, you should make n number of HTTP POST calls to the /api/numbers/special path. Try to call them in a separate goroutine
This is how you can make the post Request
```
req, err := http.NewRequest("POST", w.serverURL+"/api/numbers/special", strings.NewReader(strconv.Itoa(number)))
// do the request
```