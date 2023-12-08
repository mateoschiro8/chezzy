#include "tablero.h"

using namespace std;

void limpiarPantalla() {
    // Para limpiar la terminal antes de cada impresión
    #ifdef __linux__
        system("clear");
    #elif _WIN64
        system("cls");
    #elif _WIN32
        system("cls");
    #endif
}


void Tablero::imprimirBienvenida() {

    limpiarPantalla();

    cout << endl << "Bienvenido!" << endl << endl;

    char input;
    bool valido = false;

    while(!valido) {
        cout << "Querés empezar una nueva partida (N), cargar la guardada (L), o salir (Q)?" << endl << endl;
        cin >> input;
        if(input == 'N' || input == 'n') {

            limpiarPantalla();

            cout << "Nueva partida" << endl << endl;
            cout << "Queres jugar con blanco (B) o con negro (N)?" << endl << endl;

            cin >> input;

            while(input != 'B' && input != 'b' && input != 'N' && input != 'n') {
                limpiarPantalla();
                cout << "No conozco ese color" << endl;
                cout << "Queres jugar con blanco (B) o con negro (N)?" << endl << endl;
                cin >> input;
            }

            if(input == 'B' || input == 'b') {
                estadoJuego = 0;
                usuarioEsBlanco = true;

            } else if(input == 'N' || input == 'n') {
                estadoJuego = 0;
                usuarioEsBlanco = false;
            } 
            valido = true;

        } else if(input == 'L' || input == 'l') {
            limpiarPantalla();
            cout << "Cargar partida" << endl;
            cout << "Todavía no sabemos como guardarlas pero algún día vamos a saber" << endl;
            cargarPartida();
            valido = true;
        }

        else if(input == 'Q' || input == 'q') {
            estadoJuego = 1;
            valido = true;
        }

        else {
            limpiarPantalla();
            cout << "Incorrecto" << endl;
        }
    }

}
 
void Tablero::imprimirTablero() {
    
    limpiarPantalla();
    cout << endl;

    if(usuarioEsBlanco) {

        // Si el usuario juega con blanco
        
        cout << " A B C D E F G H" << endl << endl;
        int fila = 8;
        for(int i = 63; i >= 0; i--) {
            // Creo una máscara con un 1 en la posición del bit que quiero verificar
            uint64_t mascara = 1ULL << i;

            // Piezas blancas
            if (peonesBlancos & mascara)
                cout << " P";
            else if (caballosBlancos & mascara)
                cout << " N";
            else if (alfilesBlancos & mascara)
                cout << " B";
            else if (torresBlancas & mascara)
                cout << " R";
            else if (reinaBlanca & mascara)
                cout << " Q";
            else if (reyBlanco & mascara)
                cout << " K";

            // Piezas negras
            else if (peonesNegros & mascara)
                cout << " p";
            else if (caballosNegros & mascara)
                cout << " n";
            else if (alfilesNegros & mascara)
                cout << " b";
            else if (torresNegras & mascara)
                cout << " r";
            else if (reinaNegra & mascara)
                cout << " q";
            else if (reyNegro & mascara)
                cout << " k";

            else
                cout << " .";

            // Cada 8 piezas, salto de linea
            if(i % 8 == 0) {
                cout << "   " << fila << endl;
                fila--;
            }
        }
    }
    else {

        // Si el usuario juega con negro
        
        cout << " H G F E D C B A" << endl << endl;
        int fila = 1;
        for(int i = 63; i >= 0; i--) {
            // Creo una máscara con un 1 en la posición del bit que quiero verificar
            uint64_t mascara = 1ULL << 63 - i;
         
            // Piezas blancas
            if (peonesBlancos & mascara)
                cout << " P";
            else if (caballosBlancos & mascara)
                cout << " N";
            else if (alfilesBlancos & mascara)
                cout << " B";
            else if (torresBlancas & mascara)
                cout << " R";
            else if (reinaBlanca & mascara)
                cout << " Q";
            else if (reyBlanco & mascara)
                cout << " K";

            // Piezas negras
            else if (peonesNegros & mascara)
                cout << " p";
            else if (caballosNegros & mascara)
                cout << " n";
            else if (alfilesNegros & mascara)
                cout << " b";
            else if (torresNegras & mascara)
                cout << " r";
            else if (reinaNegra & mascara)
                cout << " q";
            else if (reyNegro & mascara)
                cout << " k";

            else
                cout << " .";

            // Cada 8 piezas, salto de linea
            if(i % 8 == 0) {
                cout << "   " << fila << endl;
                fila++;
            }
        }
    }
    cout << endl << endl;
}