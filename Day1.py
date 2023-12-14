import re
def prob1():
    # array = {"1abc2"
    # ,"pqr3stu8vwx"
    # ,"a1b2c3d4e5f"
    # ,"treb7uchet"}
    # print(array)
    array = []
    file = open("day1p1values.txt", 'r')

    for x in file:
        if x ==" ":
            continue
        else:
            array.append(x)
            print(x)


    total = 0
    a = 1
    for i in array:
        nums = re.findall("([0-9])",i)
        print(nums)
        if len(nums) == 0:
            continue
        elif len(nums) == 1:
            x = int(nums[0]) * 11
            print( str(a)+ ": "+ str(x))
            total = total + x
        elif str(nums).isdigit():
            first = str(nums[0])
            last = str(nums[len(nums)-1])
            t = first + last
            x = int(t)
            print(str(a)+ ": "+ str(x))
            total = total + x
        else:
            first = str(nums[0])
            last = str(nums[len(nums)-1])
            t = first + last
            x = int(t)
            print(str(a)+ ": "+ str(x))
            total = total + x
        a = a + 1 

    print(total)
    return
def problem2():
    print("its go time")
    file = open("day1p1values.txt", 'r')
    array = []
    for x in file:
        x = x.replace("\n" , "")
        if x ==" ":
            continue
        else:
            array.append(x)
            
    print(array)
    total = 0
    corNums = {
        "one":"1",
        "two":"2",
        "three":"3",
        "four":"4",
        "five":"5",
        "six":"6",
        "seven":"7",
        "eight":"8",
        "nine":"9"
    }
    for i in array:
        print(i)
        t = True
        while t :
            low = -1
            ar = {}
           
            if "one" in i:
                index = i.find("one")
                ar["one"] = index 
                print(str(index) + ": 1")
            if "two" in i:
                index = i.find("two")
                ar["two"] = index 
                print(str(index) + ": 2")
            if "three" in i:
                index = i.find("three")
                ar["three"] = index 
                print(str(index) + ": 3")
            if "four" in i:
                index = i.find("four")
                ar["four"] = index 
                print(str(index) + ": 4")
            if "five" in i:
                index = i.find("five")
                ar["five"] = index 
                print(str(index) + ": 5")
            if "six" in i:
                index = i.find("six")
                ar["six"] = index 
                print(str(index) + ": 6")
            if "seven" in i:
                index = i.find("seven")
                ar["seven"] = index 
                print(str(index) + ": 7")
            if "eight" in i:
                index = i.find("eight")
                ar["eight"] = index 
                print(str(index) + ": 8")
            if "nine" in i:
                index = i.find("nine")
                ar["nine"] = index 
                print(str(index) + ": 9")
            print(ar)
            if not ar:
                t = False
            else: 
                dic = dict(sorted(ar.items(), key=lambda item: item[1]))
                u = list(dic.keys())[0]
                print(ar[u])
                i = i.replace(u,u[0]+corNums[u]+u[len(u)-1],ar[u]+1)# ENGINEERING 

        print(i)

        nums = re.findall("([0-9])",i)
        print(nums)
        if len(nums) == 0:
            continue
        elif len(nums) == 1:
            x = int(nums[0]) * 11
            print(x)
            total = total + x
        elif str(nums).isdigit():
            first = str(nums[0])
            last = str(nums[len(nums)-1])
            t = first + last
            x = int(t)
            print(x)
            total = total + x
        else:
            first = str(nums[0])
            last = str(nums[len(nums)-1])
            t = first + last
            x = int(t)
            print(x)
            total = total + x

    print(str(total))
    #53845


    return

prob1()
problem2()