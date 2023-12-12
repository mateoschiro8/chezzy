#include "tablero.h"

using namespace std;

char Tablero::quePiezaEs(int columna, int fila) {

    uint64_t mascara = 1ULL << 8 * (fila - 1) + (8 - columna);

    if (peonesBlancos & mascara)
        return 'P';
    else if (caballosBlancos & mascara)
        return 'N';
    else if (alfilesBlancos & mascara)
        return 'B';
    else if (torresBlancas & mascara)
        return 'R';
    else if (reinaBlanca & mascara)
        return 'Q';
    else if (reyBlanco & mascara)
        return 'K';

    // Piezas negras
    else if (peonesNegros & mascara)
        return 'p';
    else if (caballosNegros & mascara)
        return 'n';
    else if (alfilesNegros & mascara)
        return 'b';
    else if (torresNegras & mascara)
        return 'r';
    else if (reinaNegra & mascara)
        return 'q';
    else if (reyNegro & mascara)
        return 'k';

    else
        return '.';

}

bool Tablero::movValidoEnTablero(int columnaOrigen, int filaOrigen, int columnaDestino, int filaDestino) {

    // quePiezaEs(columnaOrigen, filaOrigen)
    


    return true;

}

/*

Que pieza es, y si ese movimiento es valido para ese tipo de pieza

Si el rey no está jaque, que ese movimiento no lo ponga en jaque
Si el rey está en jaque, que ese movimiento evite el jaque

Si no es un caballo, que no atravieza ninguna pieza en su camino

Que no se coma una pieza de un mismo color

Puede enrocarse?

*/