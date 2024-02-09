#include <stdio.h>

int main() {
    int n;
    scanf("%d", &n);
    printf(n * n);
    return 0;
}

#include <stdio.h>
int main() {
    int numIterations, x;
    scanf("%d", &numIterations);
    for (x = 0; x < numIterations; x++) {
        printf("%s\n", (x % 2 == 0) ? "Hola mundo" : "Hello world");
    }
    return 0;
}