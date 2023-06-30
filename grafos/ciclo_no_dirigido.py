from tdas.grafo.grafo import Grafo
from tdas.cola.cola import Cola

def obtener_ciclo_dfs(grafo):
    visitados = {}
    padre = {}
    
    for v in grafo:
        if v not in visitados:
            ciclo = dfs_ciclo(grafo, v, visitados, padre)
            if ciclo is not None:
                return ciclo
    return None


def dfs_ciclo(grafo, v, visitados, padre):
    visitados[v] = True
    for w in grafo.adyacentes(v):
        if w in visitados:
        # Si w fue visitado y es padre de v, entonces es la arista de donde
        # vengo (no es ciclo).
        # Si no es su padre, esta arista (v, w) cierra un ciclo que empieza
        # en w.
            if w != padre[v]:
                return reconstruir_ciclo(padre, w, v)
        else:
            padre[w] = v
            ciclo = dfs_ciclo(grafo, w, visitados, padre)
            if ciclo is not None:
                return ciclo

    # Si llegamos hasta acá es porque no encontramos ningún ciclo.
    return None


def reconstruir_ciclo(padre, inicio, fin):
    v = fin
    camino = []
    while v != inicio:
        camino.append(v)
        v = padre[v]
    camino.append(inicio)
    return camino[::-1]




def obtener_ciclo_bfs(grafo):
    visitados = {}
    for v in grafo:
        if v not in visitados:
            ciclo = bfs_ciclo(grafo, v, visitados)
            if ciclo is not None:
                return ciclo
    return None


def bfs_ciclo(grafo, v, visitados):
    q = Cola()
    q.encolar(v)
    visitados[v] = True
    padre = {}  # Para poder reconstruir el ciclo
    orden = {}
    padre[v] = None
    orden[v] = 0

    while not q.esta_vacia():
        v = q.desencolar()
        for w in grafo.adyacentes(v):
            if w in visitados:
                # Si w fue visitado y es padre de v, entonces es la arista
                # de donde vengo (no es ciclo).
                # Si no es su padre, esta arista (v, w) cierra un ciclo que
                # empieza en w.
                if w != padre[v]:
                    return reconstruir_camino(padre, orden, w, v)
            else:
                q.encolar(w)
                visitados[w] = True
                padre[w] = v
                orden[w] = orden[v] + 1

    # Si llegamos hasta acá es porque no encontramos ningún ciclo.
    return None


def reconstruir_camino(padre, orden, v1, v2):
    ciclo = []
    if orden[v1] != orden[v2]: # no puede haber más que 1 de diferencia
        if orden[v1] > orden[v2]:
            ciclo.append(v1)
            v1 = padre[v1]
        else:
            ciclo.append(v2)
            v2 = padre[v2]
    while v1 != v2:
        ciclo.append(v1)
        ciclo.append(v2)
        v1 = padre[v1]
        v2 = padre[v2]
    ciclo.append(v1)
    return ciclo



def main():
    # Crear un grafo
    grafo = Grafo()

    # Agregar vértices y aristas
    grafo.agregar_vertice("A")
    grafo.agregar_vertice("B")
    grafo.agregar_vertice("C")
    grafo.agregar_vertice("D")
    grafo.agregar_arista("A", "B")
    grafo.agregar_arista("B", "C")
    grafo.agregar_arista("C", "D")
    grafo.agregar_arista("D", "A")

    # Llamar a la función para encontrar un ciclo
    ciclo = obtener_ciclo_dfs(grafo)

    # Imprimir el ciclo encontrado
    if ciclo:
        print("Ciclo encontrado:", ciclo)
        print(grafo)
    else:
        print("No se encontró ningún ciclo en el grafo.")


if __name__ == "__main__":
    main()