# Heuristics-TSP
# TSP Solver - Threshold Accepting (Go)

Implementación en Go de **Aceptación por Umbrales** (*Threshold Accepting*) aplicada a la resolución del **Problema del Agente Viajero (TSP)** en su variante de Camino Hamiltoniano. Desarrollado como proyecto para el seminario de Heurísticas de Optimización Combinatiora.

## Estructura del Proyecto

```text
.
├── cli/                 # Manejo de argumentos y lectura de entrada
├── data/                # Base de datos SQLite (tsp.db)
├── models/              # Manejadores de base de datos y consultas SQL
├── simulated_annealing/ # Framework genérico de la metaheurística (Interfaz Solution)
├── tsp/                 # Modelado del problema TSP y matriz de adyacencia
├── tsp_solver/          # Implementación concreta del TSP y mutación in-place (Rollback)
├── reporte/             # Reporte en LaTeX, gráficas e imágenes de los resultados
└── main.go              # Punto de entrada de la aplicación
```

## Parámetros de Configuración

Los parámetros de configuración de la heurística (tamaño de lote, épsilon, factor de enfriamiento, etc.) se encuentran definidos directamente en el archivo `tsp_solver/solver.go`.

## Prerrequisitos

*   **Go** (versión 1.18 o superior recomendada).
*   Un archivo de instancia con los IDs de las ciudades separados por comas (ej. `cosas.txt`).

## Dependencias

Descargar las dependencias del módulo ejecutando:

```bash
go mod tidy
```

A continuación, compila el programa para generar el binario ejecutable:

## Compilación

Para compilar el programa y generar el binario ejecutable, ejecuta el siguiente comando en la raíz del proyecto:

```bash
go build -o new main.go
```

## Ejecución

El programa lee la instancia del TSP a través de la **entrada estándar (stdin)** y acepta una bandera opcional para la semilla pseudoaleatoria (`-seed`). 

Puedes ejecutarlo utilizando `cat` de la siguiente manera:

```bash
cat cosas.txt | ./new -seed 614
```

### Opciones disponibles:
*   `-seed <entero>`: Especifica la semilla pseudoaleatoria para la ejecución (por defecto toma el valor 4 si no se proporciona).

---
*Desarrollado por Rodrigo Zaldivar Alanis — Facultad de Ciencias, UNAM.*