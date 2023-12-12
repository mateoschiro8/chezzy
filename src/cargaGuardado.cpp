#include "tablero.h"
#include <fstream>

using namespace std;    

void Tablero::guardarPartida() {

    ofstream guardarArchivo("partidaGuardada.txt");

    if (guardarArchivo.is_open()) {
        
        // guardarArchivo << "Información de asdfasdfla partida..." << endl;

        for(int i = 63; i >= 0; i--) {
            // Creo una máscara con un 1 en la posición del bit que quiero verificar
            uint64_t mascara = 1ULL << i;

            // Piezas blancas
            if (peonesBlancos & mascara)
                guardarArchivo << "P";
            else if (caballosBlancos & mascara)
                guardarArchivo << "N";
            else if (alfilesBlancos & mascara)
                guardarArchivo << "B";
            else if (torresBlancas & mascara)
                guardarArchivo << "R";
            else if (reinaBlanca & mascara)
                guardarArchivo << "Q";
            else if (reyBlanco & mascara)
                guardarArchivo << "K";

            // Piezas negras
            else if (peonesNegros & mascara)
                guardarArchivo << "p";
            else if (caballosNegros & mascara)
                guardarArchivo << "n";
            else if (alfilesNegros & mascara)
                guardarArchivo << "b";
            else if (torresNegras & mascara)
                guardarArchivo << "r";
            else if (reinaNegra & mascara)
                guardarArchivo << "q";
            else if (reyNegro & mascara)
                guardarArchivo << "k";

            else
                guardarArchivo << ".";

            // Cada 8 piezas, salto de linea
            if(i % 8 == 0) 
                guardarArchivo << endl;
        }
        guardarArchivo.close();

    } else {
        cerr << "No se pudo abrir el archivo para guardar la partida." << endl;
    }

}

void Tablero::cargarPartida() {

    ifstream cargarArchivo("partidaGuardada.txt");

    if (cargarArchivo.is_open()) {
        
        // Borra todas las piezas
        peonesBlancos   = 0;
        caballosBlancos = 0;
        alfilesBlancos  = 0;
        torresBlancas   = 0;
        reinaBlanca     = 0;
        reyBlanco       = 0;

        peonesNegros   = 0;
        caballosNegros = 0;
        alfilesNegros  = 0;
        torresNegras   = 0;
        reinaNegra     = 0;
        reyNegro       = 0;

        string fila;
        int pos = 63; 
        while(getline(cargarArchivo, fila)) {

            for(char pieza : fila) {

                // Creo una máscara con un 1 en la posición que estoy mirando
                uint64_t mascara = 1ULL << pos;
                pos--;

                switch(pieza) {



                    // Piezas blancas
                    case 'P':
                        peonesBlancos |= mascara;
                        break;
                    case 'N':
                        caballosBlancos |= mascara;
                        break;
                    case 'B':
                        alfilesBlancos |= mascara;
                        break;
                    case 'R':
                        torresBlancas |= mascara;
                        break;
                    case 'Q':
                        reinaBlanca |= mascara;
                        break;
                    case 'K':
                        reyBlanco |= mascara;
                        break;

                    // Piezas negras
                    case 'p':
                        peonesNegros |= mascara;
                        break;
                    case 'n':
                        caballosNegros |= mascara;
                        break;
                    case 'b':
                        alfilesNegros |= mascara;
                        break;
                    case 'r':
                        torresNegras |= mascara;
                        break;
                    case 'q':
                        reinaNegra |= mascara;
                        break;
                    case 'k':
                        reyNegro |= mascara;
                        break;
                    
                    case '.':
                        break;
                    
                }        
            }
        
        }

        cargarArchivo.close();

    } else {
        std::cerr << "No se pudo abrir el archivo para cargar la partida." << std::endl;
    }

}