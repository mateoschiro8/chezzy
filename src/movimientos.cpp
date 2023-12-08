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

    // Saco la pieza del lugar de destino (en caso de captura)
    mascaraDestino = ~mascaraDestino;

    peonesBlancos   = peonesBlancos & mascaraDestino;
    caballosBlancos = caballosBlancos & mascaraDestino;
    alfilesBlancos  = alfilesBlancos & mascaraDestino;
    torresBlancas   = torresBlancas & mascaraDestino;
    reinaBlanca     = reinaBlanca & mascaraDestino;
    reyBlanco       = reyBlanco & mascaraDestino;

    peonesNegros   = peonesNegros & mascaraDestino;
    caballosNegros = caballosNegros & mascaraDestino;
    alfilesNegros  = alfilesNegros & mascaraDestino;
    torresNegras   = torresNegras & mascaraDestino;
    reinaNegra     = reinaNegra & mascaraDestino;
    reyNegro       = reyNegro & mascaraDestino;

    mascaraDestino = ~mascaraDestino;

    // Pongo la pieza en el lugar de destino
    if((peonesBlancos & mascaraOrigen) != 0) 
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
    mascaraOrigen = ~mascaraOrigen;
    
    peonesBlancos   = peonesBlancos & mascaraOrigen;
    caballosBlancos = caballosBlancos & mascaraOrigen;
    alfilesBlancos  = alfilesBlancos & mascaraOrigen;
    torresBlancas   = torresBlancas & mascaraOrigen;
    reinaBlanca     = reinaBlanca & mascaraOrigen;
    reyBlanco       = reyBlanco & mascaraOrigen;

    peonesNegros   = peonesNegros & mascaraOrigen;
    caballosNegros = caballosNegros & mascaraOrigen;
    alfilesNegros  = alfilesNegros & mascaraOrigen;
    torresNegras   = torresNegras & mascaraOrigen;
    reinaNegra     = reinaNegra & mascaraOrigen;
    reyNegro       = reyNegro & mascaraOrigen;

    /*
    if(usuarioEsBlanco) {

        // Si el usuario juega con blanco       

        if((peonesBlancos & mascaraOrigen) != 0) {
            // Borro la pieza del lugar original
            mascaraOrigen = ~mascaraOrigen;
            peonesBlancos = peonesBlancos & mascaraOrigen;

            // La pongo donde va
            peonesBlancos = peonesBlancos | mascaraDestino;
        }

        else if((caballosBlancos & mascaraOrigen) != 0) {
            // Borro la pieza del lugar original
            mascaraOrigen = ~mascaraOrigen;
            caballosBlancos = caballosBlancos & mascaraOrigen;

            // La pongo donde va
            caballosBlancos = caballosBlancos | mascaraDestino;
        }

        else if((alfilesBlancos & mascaraOrigen) != 0) {
            // Borro la pieza del lugar original
            mascaraOrigen = ~mascaraOrigen;
            alfilesBlancos = alfilesBlancos & mascaraOrigen;

            // La pongo donde va
            alfilesBlancos = alfilesBlancos | mascaraDestino;
        }

        else if((torresBlancas & mascaraOrigen) != 0) {
            // Borro la pieza del lugar original
            mascaraOrigen = ~mascaraOrigen;
            torresBlancas = torresBlancas & mascaraOrigen;

            // La pongo donde va
            torresBlancas = torresBlancas | mascaraDestino;
        }
            
        else if((reinaBlanca & mascaraOrigen) != 0) {
            // Borro la pieza del lugar original
            mascaraOrigen = ~mascaraOrigen;
            reinaBlanca = reinaBlanca & mascaraOrigen;

            // La pongo donde va
            reinaBlanca = reinaBlanca | mascaraDestino;
        }
            
        else if((reyBlanco & mascaraOrigen) != 0) {
            // Borro la pieza del lugar original
            mascaraOrigen = ~mascaraOrigen;
            reyBlanco = reyBlanco & mascaraOrigen;

            // La pongo donde va
            reyBlanco = reyBlanco | mascaraDestino;
        }
            
        // Saco la pieza del otro color que hay en su lugar
        mascaraDestino = ~mascaraDestino;
        peonesNegros   = peonesNegros & mascaraDestino;
        caballosNegros = caballosNegros & mascaraDestino;
        alfilesNegros  = alfilesNegros & mascaraDestino;
        torresNegras   = torresNegras & mascaraDestino;
        reinaNegra     = reinaNegra & mascaraDestino;
        reyNegro       = reyNegro & mascaraDestino;

    }
    else {

        // Si el usuario juega con negro

        if((peonesNegros & mascaraOrigen) != 0) {
            // Borro la pieza del lugar original
            mascaraOrigen = ~mascaraOrigen;
            peonesNegros = peonesNegros & mascaraOrigen;

            // La pongo donde va
            peonesNegros = peonesNegros | mascaraDestino;
        }

        else if((caballosNegros & mascaraOrigen) != 0) {
            // Borro la pieza del lugar original
            mascaraOrigen = ~mascaraOrigen;
            caballosNegros = caballosNegros & mascaraOrigen;

            // La pongo donde va
            caballosNegros = caballosNegros | mascaraDestino;
        }

        else if((alfilesNegros & mascaraOrigen) != 0) {
            // Borro la pieza del lugar original
            mascaraOrigen = ~mascaraOrigen;
            alfilesNegros = alfilesNegros & mascaraOrigen;

            // La pongo donde va
            alfilesNegros = alfilesNegros | mascaraDestino;
        }

        else if((torresNegras & mascaraOrigen) != 0) {
            // Borro la pieza del lugar original
            mascaraOrigen = ~mascaraOrigen;
            torresNegras = torresNegras & mascaraOrigen;

            // La pongo donde va
            torresNegras = torresNegras | mascaraDestino;
        }
            
        else if((reinaNegra & mascaraOrigen) != 0) {
            // Borro la pieza del lugar original
            mascaraOrigen = ~mascaraOrigen;
            reinaNegra = reinaNegra & mascaraOrigen;

            // La pongo donde va
            reinaNegra = reinaNegra | mascaraDestino;
        }
            
        else if((reyNegro & mascaraOrigen) != 0) {
            // Borro la pieza del lugar original
            mascaraOrigen = ~mascaraOrigen;
            reyNegro = reyNegro & mascaraOrigen;

            // La pongo donde va
            reyNegro = reyNegro | mascaraDestino;
        }
            
        // Saco la pieza del otro color que hay en su lugar
        mascaraDestino = ~mascaraDestino;
        peonesBlancos   = peonesBlancos & mascaraDestino;
        caballosBlancos = caballosBlancos & mascaraDestino;
        alfilesBlancos  = alfilesBlancos & mascaraDestino;
        torresBlancas   = torresBlancas & mascaraDestino;
        reinaBlanca     = reinaBlanca & mascaraDestino;
        reyBlanco       = reyBlanco & mascaraDestino;

    }
    */
}