f x = x + 1
add x y = x + y

f2 1 = "1"
f2 a = "?"

fact n
    | n == 0    = 1
    | otherwise = n * fact (n - 1)

main = do
    let a = 1
        b = 2
        c = a + b
    print c
    print $ add 1 2
    print $ f 2

    if c == 3
        then print "true"
        else print "false"
    print $ f2 3
    print $ f2 1

    print $ fact 5

    print $ [1,2,3,4,5] !! 3
