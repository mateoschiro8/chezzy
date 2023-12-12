// COMPILAR Y EJECUTAR DESDE AFUERA DE SRC CON    g++ src/*.cpp -o main   ->   ./main

#include <iostream>
#include "tablero.h"

using namespace std;

int main() {

    Tablero tab;

    tab.imprimirBienvenida();

    while(!tab.estado())
        tab.obtenerMovimiento();
        
    // tab.guardarPartida();
    // tab.cargarPartida();
    // tab.imprimirTablero();
    // cout << "Gracias por jugar!" << endl << endl;

}