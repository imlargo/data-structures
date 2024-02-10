

// 1. Encuentra el cuadrado
#include <stdio.h>
int main() {
	int n;
	scanf("%d", &n);
	printf(n * n);
	return 0;
}

// 2. Hola mundo bilingüe
#include <stdio.h>
int main() {
	int numIterations, x;
	scanf("%d", &numIterations);
	for (x = 0; x < numIterations; x++)
	{
		printf("%s\n", (x % 2 == 0) ? "Hola mundo" : "Hello world");
	}
	return 0;
}

// 3. Positivos y negativos
#include <stdio.h>
int main() {
	int sumPos = 0;
	int sumNeg = 0;
	int n, input;
	scanf("%d", &n);

	for (int i = 0; i < n; i++) {
		scanf("%d", &input);
		if (input > 0) {
			sumPos += input;
		} else {
			sumNeg += input;
		}
	}
	printf("positivos %d, negativos %d", sumPos, sumNeg);

	return 0;
}

// 4. Los 4 fantásticos
#include <stdio.h>
int main() {

	int n;
	scanf("%d", &n);

	if (n % 2 == 0) {
		printf("es multiplo de 2");
		return 0;
	} else if (n % 3 == 0) {
		printf("es multiplo de 3");
		return 0;
	} else if (n % 5 == 0) {
		printf("es multiplo de 5");
		return 0;
	} else if (n % 7 == 0) {
		printf("es multiplo de 7");
		return 0;
	} else {
		printf("no es multiplo de ninguno de los primeros cuatro primos");
		return 0;
	}

	return 0;
}

// 5. Potencias menores
#include <stdio.h>
#include <math.h>
int main() {

	double n, limit;
	scanf("%d", &n);
	scanf("%d", &limit);

	for (double p = 1; p < limit; p++) {
		
		double r = pow(n, p);
		if (r > limit) {
			return 0;
		} else {
			printf("%d\n", (int)r);
		}
	}
	
	return 0;
}

// 6. Conjetura de Collatz
#include <stdio.h>
int main() {
	return 0;
}