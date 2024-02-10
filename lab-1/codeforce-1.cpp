#include <iostream>
#include <cmath>

using namespace std;

// 1. Encuentra el cuadrado
void findSquare() {
    int n;
    cin >> n;
    cout << n * n << endl;
}

// 2. Hola mundo bilingüe
void helloWorld() {
    int n;
    cin >> n;
    for(int x = 0; x < n; x++) {
        cout << (x % 2 == 0 ? "Hola mundo" : "Hello world") << endl;
    }
}

// 3. Positivos y negativos
void posAndNeg() {
    int sumPos = 0;
    int sumNeg = 0;
    int n;
    cin >> n;
    for(int i = 0; i < n; i++) {
        int num;
        cin >> num;
        if (num > 0) {
            sumPos += num;
        } else {
            sumNeg += num;
        }
    }
    cout << "positivos " << sumPos << ", negativos " << sumNeg << endl;
}

// 4. Los 4 fantásticos
void fantasticFour() {
    int n;
    cin >> n;
    int d[4] = {2, 3, 5, 7};
    for(int i = 0; i < 4; i++) {
        if(n % d[i] == 0) {
            cout << "es multiplo de " << d[i] << endl;
            exit(0);
        }
    }
    cout << "no es multiplo de ninguno de los primeros cuatro primos" << endl;
}

// 5. Potencias menores
void powerLess() {
    int n, limit;
    cin >> n >> limit;
    for(int p = 1; p < limit; p++) {
        int r = pow(n, p);
        if(r > limit) {
            break;
        }
        cout << r << endl;
    }
}

// 6. Conjetura de Collatz
void collatzConjecture() {
    int n;
    cin >> n;
    while(n > 1) {
        cout << n << endl;
        n = (n % 2 == 0) ? n / 2 : (3 * n) + 1;
    }
    cout << n << endl;
}

int main() {
    return 0;
}