def pro1():
    file = open("exm3.txt", 'r')
    array = []
    for x in file:
        x = x.replace("\n" , "")
        if x ==" ":
            continue
        else:
            array.append(x)
    print(array)

    currline = 0
    dic = {}
    for i in array:
        # first we must find digit
        splitInToDict(i)

        # second we must find symbol 
        # if digit is true then add else skip 
        print("apples")



        
 
pro1()