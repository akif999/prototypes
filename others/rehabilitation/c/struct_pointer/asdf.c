#include <stdio.h>

typedef struct {
    int id;
    int value;
} Item;

int main(void) {
    Item item = {1, 42};
    Item *p = &item;

    printf("id = %d, value = %d\n", p->id, p->value);

    p->value = 100;
    printf("after change value = %d\n", item.value);

    return (0);
}
