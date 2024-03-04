#include <stdio.h>
#include <pthread.h>
#include <unistd.h>

void * funcion(void * arg) {
    printf("Hilo %d \n", (int) pthread_self());
    pthread_exit(0);
}

int main() {
    int j;
    pthread_t hilos[10];
    pthread_attr_t attr;

    pthread_attr_init(&attr);
    pthread_attr_setdetachstate(&attr, PTHREAD_CREATE_DETACHED);

    for (j = 0; j < 10; j++) {
        printf("Creando hilo %d\n", j);
        pthread_create(&hilos[j], &attr, funcion, NULL);
    }

    printf("Terminando el programa\n");
    return 0;
}