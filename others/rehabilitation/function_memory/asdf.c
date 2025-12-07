#include <stdio.h>

int add(int a, int b) {
    return (a + b);
}

int main(void) {
    int (*func)(int, int);

    func = add;

    int result = func(10, 20);
    printf("result = %d\n", result);

    return(0);
}
