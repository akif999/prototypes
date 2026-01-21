#include <stdio.h>

int main(void) {
    unsigned long i;
    for (i = 1; i <= 100; i++) {
        if ((i % 3 == 0) && (i % 5 == 0)) {
            printf("FizzBuzz\n", i);
        } else if (i % 3 == 0) {
            printf("Fizz\n", i);
        } else if (i % 5 == 0) {
            printf("Buzz\n", i);
        } else {
            printf("%u\n", i);
        }
    }
    return (0);
}
