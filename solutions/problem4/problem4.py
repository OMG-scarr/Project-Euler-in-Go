def is_palindrome(n):
    s = str(n)
    return s == s[::-1]

# Example usage:
if __name__ == "__main__":
    num = int(input("Enter a number: "))
    print(is_palindrome(num))