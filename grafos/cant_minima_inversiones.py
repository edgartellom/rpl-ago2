from tdas.grafo.grafo import Grafo
from tdas.cola_prioridad.heap import Heap
INF = float('inf')

def minimas_inversiones(grafo, s, t):
    grafo_pesado = Grafo()
    for v in grafo:
        if not v in grafo_pesado: grafo_pesado.agregar_vertice(v)
        for w in grafo.adyacentes(v):
            if not w in grafo_pesado: grafo_pesado.agregar_vertice(w)
            grafo_pesado.agregar_arista(v, w, 0)
            if not grafo_pesado.hay_arista(w, v): grafo_pesado.agregar_arista(w, v, 1)

    return camino_minimo(grafo_pesado, s, t)

def camino_minimo(grafo, s, t):
    distancia = {}
    for v in grafo:
        distancia[v] = INF
    distancia[s] = 0
    q = Heap()
    q.encolar((s, 0))

    while not q.esta_vacia():
        v, _ = q.desencolar()
        for w in grafo.adyacentes(v):
            if (distancia[v] + grafo.peso_arista(v, w) < distancia[w]):
                distancia[w] = distancia[v] + grafo.peso_arista(v, w)
                q.encolar((w, distancia[w]))
    return distancia[t]