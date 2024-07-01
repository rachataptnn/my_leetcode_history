//           0     1   2       3       4     5
//           0 1   0   0 1 2   0 1 2   0 1   0 1 2
// graph = [[2,5],[3],[0,4,5],[1,4,5],[2,3],[0,2,3]]

doSomeThing() {} -> 

"0:2"        1
"0:2:4"      2
"0:2:4:3"    3

this is the one key that i want?
"0:2:4:3:1"  4 -> most far in the shortest path 

"0:5"        1
"0:5:3"      2

this is the one key that i want?
the longest
"0:5:3:1"    3

Default State of 'alreadyFind' is alreadyFind[0] = true (โจทย์ให้ constraint มางี้)

1. find '0' in node get 2, 5
    got [
        "0:2": 1, 
        "0:5": 1,   -> nearZeroMap[2] = true, nearZeroMap[5] = true
    ]               -> alreadyFind[0] = true

    ### find 2
    2-a. find '2' in node get 0, 4, 5 
        but 0 is 'alreadyFind'  so 0 is not append
        but 5 is 'nearZero'     so 5 is not append
        got [
            "0:2:4": 2
        ]

    3-a. find '4' in node get 2, 3
        but 2 is alredy find
        got [
            "0:2:4:3": 3
        ]

    4-a. find '3' in node get 1, 4, 5
        but 4 is 'alreadyFind' so 4 is not append
        but 5 is 'alreadyFind' so 5 is not append
        got [
            "0:2:4:3:1": 4
        ]
    5-a. find '1' in node get 3
        but 3 is 'alreadyFind' so 1 is not append
        got [nil]
        ### this is 'end case' GO BACK~~

    ### find 5 (Reset all "alreadyFind")
    2-b. find '5' in node get 0, 2, 3
        but 0 is 'alreadyFind' so 0 is not append
        but 2 is 'nearZero' so 2 is not append
        got [
            "0:5:3": 2
        ]
    2-c. find '3' in node get 1, 4, 5
        but 5 is 'alreadyFind' so 5 is not append
        got [
            "0:5:3:4":
                find 4 in node get 2, 3
                but 2 is 'nearZeo' so 2 is not append
                but 3 is 'alreadyFind' so 3 is not append
            "0:5:3:1": 3
        ]

    2-d. find '1' in node get 3
        but 3 is 'alreadyFind' so 3 is not append
        got [nil]
        ### this is 'end case' GO BACK~~