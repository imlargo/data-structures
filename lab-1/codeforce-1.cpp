#include <iostream>
using namespace std;

// 1. Encuentra el cuadrado
int main() {
    int n;
    cin >> n;
    cout << n * n;
    return 0;
}

// 2. Hola mundo bilingüe
int main() {

    int n;
    cin >> n;

    for (int i = 0; i < n; i++) {
        cout << ((i % 2 == 0) ? "Hola mundo" : "Hello world") << "\n";
    }

    return 0;
}

// 3. Positivos y negativos
int main() {
    int sumPos = 0;
    int sumNeg = 0;

    int n;
    cin >> n;

    for (int i = 0; i < n; i++) {
        int num;
        cin >> num;
        if (num > 0) {
            sumPos += num;
        } else {
            sumNeg += num;
        }
    }

    cout << "positivos " << sumPos << ", negativos " << sumNeg;

    return 0;
}
// 4. Los 4 fantásticos
int main() {
    int n;
    cin >> n;

    int divs[] = {2, 3, 5, 7};

    for (int d : divs) {
        if (n % d == 0) {
            cout << "es multiplo de " << d;
            return 0;
        }
    }

    cout << "no es multiplo de ninguno de los primeros cuatro primos";
    return 0;
}

// 5. Potencias menores
#include <cmath>

int main() {
    
    double n, limit;
    cin >> n;
    cin >> limit;

    for (int p = 1; p < limit; p++) {
        double r = pow(n, p);
        if (r > limit) {
            return 0;
        }
        cout << r << "\n";
    }

    return 0;
}

// 6. Conjetura de Collatz
int main() {
    double n;
    cin >> n;

    while (n > 1) {
        cout << (int)n;
        n = ((int)n % 2 == 0) ? n / 2.0  : (3.0 * n) + 1.0;
    }
    cout << (int)n << "\n";

    return 0;
}
