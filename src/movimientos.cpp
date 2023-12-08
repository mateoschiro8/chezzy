#include "tablero.h"
#include "validoEnTablero.h"

#include <string>

using namespace std;

// Si un movimiento es válido semánticamente
bool movValidoEnInput(int columna, int fila) {

    if(!(columna >= 1 && columna <= 8) || !(fila >= 1 && fila <= 8))
        return false;

    return true;
}

void Tablero::obtenerMovimiento() {

    string mov;
    bool valido = false;

    char columnaTmpOrigen, filaTmpOrigen, columnaTmpDestino, filaTmpDestino;
    int columnaOrigen, filaOrigen, columnaDestino, filaDestino;

    imprimirTablero();
    
    while(!valido) {

        cout << "Ingresar posición origen y destino (ColumnaFilaColumnaFila):" << endl << endl;

        cin >> mov;

        if(mov.size() != 4) {
            // imprimirTablero();
            cout << "Movimiento no válido, intente de nuevo" << endl << endl;
            continue;
        }

        columnaTmpOrigen = mov[0];
        filaTmpOrigen = mov[1];

        columnaTmpDestino = mov[2];
        filaTmpDestino = mov[3];

        imprimirTablero();
        
        // Transformar los caracteres de las filas y las columnas a los enteros correspondientes
        if (columnaTmpOrigen >= 'A' && columnaTmpOrigen <= 'Z') 
            columnaOrigen = columnaTmpOrigen - 'A' + 1;
        else
            columnaOrigen = columnaTmpOrigen - 'a' + 1;
            
        filaOrigen = filaTmpOrigen - '0';
        


        if (columnaTmpDestino >= 'A' && columnaTmpDestino <= 'Z') 
            columnaDestino = columnaTmpDestino - 'A' + 1;
        else
            columnaDestino = columnaTmpDestino - 'a' + 1;
            
        filaDestino = filaTmpDestino - '0';
        
        if(!movValidoEnInput(columnaOrigen, filaOrigen) || !movValidoEnInput(columnaDestino, filaDestino) ||
           !movValidoEnTablero(columnaOrigen, filaOrigen, columnaDestino, filaDestino)) {
            imprimirTablero();
            cout << "Movimiento no válido, intente de nuevo" << endl << endl;
            continue;
        }

        valido = true;
    }

    mover(columnaOrigen, filaOrigen, columnaDestino, filaDestino);
}

void Tablero::mover(int columnaOrigen, int filaOrigen, int columnaDestino, int filaDestino) {

    // Hay que ver qué pieza hay en la posición origen
    // Armo máscaras con un 1 en esa posición
    uint64_t mascaraOrigen = 1ULL << 8 * (filaOrigen - 1) + (8 - columnaOrigen);
    uint64_t mascaraDestino = 1ULL << 8 * (filaDestino - 1) + (8 - columnaDestino);

    uint64_t notMascaraOrigen = ~mascaraOrigen;
    uint64_t notMascaraDestino = ~mascaraDestino;




    // Saco la pieza del lugar de destino (en caso de captura)
    peonesBlancos   = peonesBlancos & notMascaraDestino;
    caballosBlancos = caballosBlancos & notMascaraDestino;
    alfilesBlancos  = alfilesBlancos & notMascaraDestino;
    torresBlancas   = torresBlancas & notMascaraDestino;
    reinaBlanca     = reinaBlanca & notMascaraDestino;
    reyBlanco       = reyBlanco & notMascaraDestino;

    peonesNegros   = peonesNegros & notMascaraDestino;
    caballosNegros = caballosNegros & notMascaraDestino;
    alfilesNegros  = alfilesNegros & notMascaraDestino;
    torresNegras   = torresNegras & notMascaraDestino;
    reinaNegra     = reinaNegra & notMascaraDestino;
    reyNegro       = reyNegro & notMascaraDestino;

    
    // Pongo la pieza en el lugar de destino
    if((peonesBlancos & mascaraOrigen)) 
        peonesBlancos = peonesBlancos | mascaraDestino;

    else if((caballosBlancos & mascaraOrigen) != 0)
        caballosBlancos = caballosBlancos | mascaraDestino;
    
    else if((alfilesBlancos & mascaraOrigen) != 0)
        alfilesBlancos = alfilesBlancos | mascaraDestino;
    
    else if((torresBlancas & mascaraOrigen) != 0)
        torresBlancas = torresBlancas | mascaraDestino;
                
    else if((reinaBlanca & mascaraOrigen) != 0)
        reinaBlanca = reinaBlanca | mascaraDestino;
            
    else if((reyBlanco & mascaraOrigen) != 0)
        reyBlanco = reyBlanco | mascaraDestino;
            
    else if((peonesNegros & mascaraOrigen) != 0) 
        peonesNegros = peonesNegros | mascaraDestino;
        
    else if((caballosNegros & mascaraOrigen) != 0)
        caballosNegros = caballosNegros | mascaraDestino;
    
    else if((alfilesNegros & mascaraOrigen) != 0)
        alfilesNegros = alfilesNegros | mascaraDestino;
    
    else if((torresNegras & mascaraOrigen) != 0)
        torresNegras = torresNegras | mascaraDestino;
        
    else if((reinaNegra & mascaraOrigen) != 0)
        reinaNegra = reinaNegra | mascaraDestino;    
    
    else if((reyNegro & mascaraOrigen) != 0)
        reyNegro = reyNegro | mascaraDestino;


    // Borro la pieza del lugar original
    peonesBlancos   = peonesBlancos & notMascaraOrigen;
    caballosBlancos = caballosBlancos & notMascaraOrigen;
    alfilesBlancos  = alfilesBlancos & notMascaraOrigen;
    torresBlancas   = torresBlancas & notMascaraOrigen;
    reinaBlanca     = reinaBlanca & notMascaraOrigen;
    reyBlanco       = reyBlanco & notMascaraOrigen;

    peonesNegros   = peonesNegros & notMascaraOrigen;
    caballosNegros = caballosNegros & notMascaraOrigen;
    alfilesNegros  = alfilesNegros & notMascaraOrigen;
    torresNegras   = torresNegras & notMascaraOrigen;
    reinaNegra     = reinaNegra & notMascaraOrigen;
    reyNegro       = reyNegro & notMascaraOrigen;

}