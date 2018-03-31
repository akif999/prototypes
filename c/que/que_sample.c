#include <stdio.h>

#define BUFFER_SIZE 5

typedef struct
{
    unsigned short buf[BUFFER_SIZE];
    unsigned short  *head;
    unsigned short  *tail;
} ring_buffer;

void init_ring_buffer(ring_buffer *r_buf);
void dump_buffer(ring_buffer r_buf);
void enqueue(ring_buffer *r_buf, unsigned short elem);
unsigned short dequeue(ring_buffer *r_buf);

int main (void)
{
    ring_buffer r_buf;

    init_ring_buffer(&r_buf);
    enqueue(&r_buf, 0x1111);
    enqueue(&r_buf, 0x2222);
    enqueue(&r_buf, 0x3333);
    enqueue(&r_buf, 0x4444);
    enqueue(&r_buf, 0x5555);
    enqueue(&r_buf, 0x6666);
    enqueue(&r_buf, 0x7777);
    dump_buffer(r_buf);
    printf("deque  : %04X\n", dequeue(&r_buf));
    printf("deque  : %04X\n", dequeue(&r_buf));
    printf("deque  : %04X\n", dequeue(&r_buf));
    dump_buffer(r_buf);
}

void init_ring_buffer(ring_buffer *r_buf)
{
    int i;
    for (i = 0; i < sizeof(r_buf->buf) / sizeof(unsigned short); i++)
    {
        r_buf->buf[i] = 0x0000;
    }
    r_buf->head = &r_buf->buf[0];
    r_buf->tail = &r_buf->buf[0];
}

void enqueue(ring_buffer *r_buf, unsigned short elem)
{
    *r_buf->tail = elem;
    if (r_buf->tail == &r_buf->buf[BUFFER_SIZE - 1])
    {
        r_buf->tail = &r_buf->buf[0];
    } else {
        r_buf->tail++;
    }
}

unsigned short dequeue(ring_buffer *r_buf)
{
    unsigned short ret;

    ret = *r_buf->head;
    *r_buf->head = 0x0000;
    r_buf->head++;
    return ret;
}

void dump_buffer(ring_buffer r_buf)
{
    int i;
    printf("buffer : ");
    for (i = 0; i < sizeof(r_buf.buf) / sizeof(unsigned short); i++)
    {
        printf("%04X ", r_buf.buf[i]);
    }
    printf("\n");
}
