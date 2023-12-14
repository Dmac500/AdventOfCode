import re
def pro1():
    file = open("day2pset.txt", 'r')
    array = []
    for x in file:
        x = x.replace("\n" , "")
        if x ==" ":
            continue
        else:
            array.append(x)

    games = {}
    
    for i in array:
        redCube = 12
        greenCube = 13
        blueCube = 14 
        i = i.split(":")

        games[i[0]] = i[1].split(";")
        
    allowedIdSums = 0
    void = 0 
    for i in games.keys():
        blue = True
        red = True
        green = True
        print(i)
        for t in games[i]:
            t = t.split(",")
            for q in t:
                #print(q)
                if "blue" in q:
                    q = q.replace("blue","") 
                    if int(q) > blueCube:
                        print(q)
                        blue = False

                if "red" in q:
                    q = q.replace("red","")
                    if int(q) > redCube:
                        print(q)
                        red = False  
                 
                if "green" in q:
                    q = q.replace("green","")
                    if int(q) > greenCube:
                        print(q)
                        green = False  
        if blue == True and red == True and green == True:
            gamenum = int(i.replace("Game ", ""))
            print("this has all sause "+ str(gamenum))
            allowedIdSums = allowedIdSums + gamenum
    print(allowedIdSums)
                 



def pro12():
    saveGame= []
    file = open("day2pset.txt", 'r')
    array = []
    for x in file:
        x = x.replace("\n" , "")
        if x ==" ":
            continue
        else:
            array.append(x)

    games = {}
    
    for i in array:
        redCube = 12
        greenCube = 13
        blueCube = 14 
        i = i.split(":")
        games[i[0]] = i[1].split(";")
    allowedIdSums = 0  
    for i in games.keys():
        blueNum=0
        redNum=0
        greenNum=0 
        print(i)
        for t in games[i]:
            t = t.split(",")
            for q in t:
                if "blue" in q:
                     nums = re.findall(r'\d+',q)
                     blueNum = blueNum + int(nums[0])
                if "red" in q:
                     nums = re.findall(r'\d+',q)
                     redNum = redNum + int(nums[0]) 
                if "green" in q:
                     nums = re.findall(r'\d+',q)
                     greenNum = greenNum + int(nums[0])
        print("red " + str(redNum) + "/ "+ str(redCube)) 
        print("green " + str(greenNum)+ "/ "+ str(greenCube))
        print("blue " + str(blueNum)+ "/ "+ str(blueCube))
        
        if blueNum <= blueCube and redNum <=redCube and greenNum <=greenCube:
            print(re.findall(r'\d+',i))
            saveGame.append(re.findall(r'\d+',i))
            allowedIdSums = allowedIdSums + int(re.findall(r'\d+',i)[0])

    print(allowedIdSums)
    print(saveGame)
    #102
                 
    
pro1()
def pro2():
    file = open("day2pset.txt", 'r')
    array = []
    for x in file:
        x = x.replace("\n" , "")
        if x ==" ":
            continue
        else:
            array.append(x)

    games = {}
    for i in array:
        i = i.split(":")
        games[i[0]] = i[1].split(";")

    addThisHomie = []

    for i in games.keys():
        blue = 0
        red = 0
        green = 0
        print(i)
        for t in games[i]:
            t = t.split(",")
            for q in t:
                if "blue" in q:
                    q = q.replace("blue","") 
                    if int(q) > blue:
                        blue = int(q)
                if "red" in q:
                    q = q.replace("red","")
                    if int(q) > red:
                        red = int(q)  
                if "green" in q:
                    q = q.replace("green","")
                    if int(q) > green:
                        green = int(q) 

        apple = red * blue * green # its 3am we have given up on any variable names 
        addThisHomie.append(apple)
        print(apple)
    total = 0
    for addThisNumber in addThisHomie:
        total = addThisNumber + total
    print("Total: "+ str(total))
    
#pro2()