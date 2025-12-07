#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int id;
    int value;
} Item;

int main(void) {
    int n = 3;
    Item *items = malloc(sizeof(Item) * n);

    if (!items) {
        return (1);
    }

    for (int i = 0; i < n; i++) {
        items[i].id = i;
        items[i].value = i * 10;
    }

    for (int i = 0; i < n; i++) {
        printf("id=%d value=%d\n", items[i].id, items[i].value);
    }

    free(items);

    return (0);
}
