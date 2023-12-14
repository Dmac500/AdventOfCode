def fib():
    array = [0,1]
    for i in range(0 , 100):
        t = array[i] + array[i+1]
        array.append(t)
    print(array)

def fibRecursive(n):
    if n > 100:
        return
     
    

def fibDP(n, memo={}):
    if n in memo:
        return memo[n]
    if n <= 1:
        return n
    else:
        result = fib(n-1, memo) + fib(n-2, memo)
        memo[n] = result
        return result




fib()