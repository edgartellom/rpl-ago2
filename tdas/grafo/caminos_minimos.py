from tdas.cola.cola import Cola
from tdas.cola_prioridad.heap import Heap

def camino_minimo_bfs(grafo, origen):
    distancia, padre, visitado = {}, {}, {}
    
    for v in grafo:
        distancia[v] = float('inf')
    
    distancia[origen] = 0
    padre[origen] = None
    visitado[origen] = True
    q = Cola()
    q.encolar(origen)

    while not q.esta_vacia():
        v = q.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitado:
                distancia[w] += distancia[v] + 1
                padre[w] = v
                visitado[w] = True
                q.encolar(w)
    
    return padre, distancia

def camino_minimo_dijkstra(grafo, origen):
    distancia, padre = {}, {}

    for v in grafo:
        distancia[v] = float('inf')
    
    distancia[origen] = 0
    padre[origen] = None
    q = Heap()
    q.encolar((origen, 0))

    while not q.esta_vacia():
        v, _ = q.desencolar()
        # Si tuviera un destino definido:
        # if v == destino: return padre, distancia
        for w in grafo.adyacentes(v):
            if distancia[v] + grafo.peso_arista(v, w) < distancia[w]:
                distancia[w] = distancia[v] + grafo.peso_arista(v, w)
                padre[w] = v
                q.encolar((w, distancia[w]))
    
    return padre, distancia




