
char quePiezaEs(int columna, int fila) {

    uint64_t mascaraOrigen = 1ULL << 8 * (fila - 1) + (8 - columna);

    

}
// Si un movimiento es válido en el tablero
bool movValidoEnTablero(int columnaOrigen, int filaOrigen, int columnaDestino, int filaDestino) {

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