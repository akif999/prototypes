#include <stdio.h>
#include <stdlib.h>

#define NOT_ERROR 0
#define ERROR 1

int main(void) {
    int n = 5;
    int *arr = malloc(n * sizeof(int));

    if (!arr) {
        perror("malloc");
        return (ERROR);
    }

    int i;
    for (i = 0; i < n; i++) {
        arr[i] = i * 10;
    }

    for (i = 0; i < n; i++) {
        printf("%d\n", arr[i]);
    }

    free(arr);

    return (NOT_ERROR);
}
