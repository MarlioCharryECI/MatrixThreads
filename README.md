# Simulación de Agentes en Go

### Marlio Jose Charry Espitia / 2026-1

Una simulación concurrente donde Neo navega por un mundo en cuadrícula, esquivando agentes mientras intenta alcanzar teléfonos.

## 🚀 Características

- Simulación basada en cuadrícula 2D
- Múltiples agentes que persiguen a Neo
- Neo con inteligencia para esquivar agentes y alcanzar teléfonos
- Movimiento concurrente de todas las entidades
- Parámetros configurables (tamaño del mapa, número de agentes, teléfonos y muros)
- Visualización en tiempo real en la terminal

## 🛠️ Requisitos Previos

- Go 1.16 o superior

## ⚙️ Instalación

1. Clona el repositorio:
   ```bash
   git clone https://github.com/MarlioCharryECI/MatrixThreads.git
   cd MatrixThreads
   ```

## 🔄 Lógica de Concurrencia

El juego utiliza goroutines para manejar el movimiento concurrente de Neo y los agentes. Aquí está cómo funciona:

### Estructura de Goroutines

1. **Hilo Principal**
   - Se encarga de la visualización del juego en tiempo real
   - Actualiza la pantalla cada 300ms
   - Espera señales de finalización del juego

2. **Goroutine de Neo**
   - Controla el movimiento autónomo de Neo
   - Se ejecuta cada 700ms
   - Implementa la lógica de persecución de teléfonos y evasión de agentes
   - Notifica cuando Neo gana (consigue todos los teléfonos)

3. **Goroutines de Agentes**
   - Cada agente tiene su propia goroutine
   - Se ejecutan cada 900ms
   - Persiguen a Neo usando un algoritmo de búsqueda de ruta
   - Notifican cuando atrapan a Neo

### Sincronización

- **Mutex**: Se utiliza para proteger el acceso concurrente al estado del mundo (World)
- **Canales**:
   - `done`: Canal de string para notificar el fin del juego y su resultado
   - Cada goroutine escribe en este canal cuando se cumple una condición de victoria/derrota

### Flujo del Juego

1. Se inician todas las goroutines (Neo + Agentes)
2. Cada entidad se mueve de forma independiente según su temporizador
3. El hilo principal actualiza la pantalla periódicamente
4. Cuando ocurre una condición de fin de juego:
   - La goroutine correspondiente envía un mensaje al canal `done`
   - El hilo principal recibe el mensaje y muestra el resultado
   - Todas las goroutines terminan cuando el programa finaliza

### Condiciones de Victoria/Derrota

- **Victoria**: Neo recoge todos los teléfonos en el mapa
- **Derrota**: Cualquier agente atrapa a Neo
## 🎮 Cómo Usar

Ejecuta la simulación:
```bash
go run .
```

### Configuración

Puedes modificar los parámetros de la simulación en `main.go` cambiando la estructura `config`:

```go
var config = Config{
    Rows:      10,     // Número de filas en la cuadrícula
    Cols:      10,     // Número de columnas en la cuadrícula
    NumAgents: 2,      // Número de agentes en la simulación
    NumPhones: 1,      // Número de teléfonos que Neo debe alcanzar
    NumWalls:  2,      // Número de muros en la simulación
}
```

## 🎯 Reglas del Juego

- `N` - Neo (el jugador)
- `A` - Agentes (enemigos que persiguen a Neo)
- `P` - Teléfonos (objetivos que Neo debe alcanzar)
- `#` - Muros (obstáculos)
- `.` - Espacio vacío

Neo gana al alcanzar todos los teléfonos mientras esquiva a los agentes.