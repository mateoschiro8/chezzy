#ifndef tablero_H_
#define tablero_H_

#include <iostream>

using namespace std;

class Tablero {

    public:

        // Constructor. Asigna las posiciones iniciales de las piezas 
        Tablero();

        // Imprime bienvenida. Pregunta que quiere jugar el usuario
        void imprimirBienvenida();

        // Imprime tablero, dependiendo qué color está jugando el usuario
        void imprimirTablero();

        // Obtiene el movimiento del usuario, y llama a mover()
        void obtenerMovimiento();

        // Dado un movimiento (sea del usuario o no), lo realiza en el tablero
        void mover(int columnaOrigen, int filaOrigen, int columnaDestino, int filaDestino);

        // Guardado y cargado de partida
        void guardarPartida();
        void cargarPartida();

        
        // Estado del juego
        int estado() { return estadoJuego; }




    private:

        // Piezas blancas
        uint64_t peonesBlancos;
        uint64_t caballosBlancos;
        uint64_t alfilesBlancos;
        uint64_t torresBlancas;
        uint64_t reinaBlanca;
        uint64_t reyBlanco;

        // Piezas negras
        uint64_t peonesNegros;
        uint64_t caballosNegros;
        uint64_t alfilesNegros;
        uint64_t torresNegras;
        uint64_t reinaNegra;
        uint64_t reyNegro;

        bool juegaUsuario;
        bool usuarioEsBlanco;

        int estadoJuego;  // 0 si el juego está en progreso, 1 si no quiere jugar más, 2 si ganó blanco, 3 si ganó negro

}; 


#endif