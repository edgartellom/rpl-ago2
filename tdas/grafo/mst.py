from tdas.cola_prioridad.heap import Heap
from grafo import Grafo

def mst_prim(grafo):
    v = grafo.vertice_aleatorio()
    visitados = set()
    visitados.add(v)
    q = Heap()

    for w in grafo.adyancentes(v):
        q.encolar(((v, w), grafo.peso_arista(v, w)))
    
    arbol = Grafo(dirigido=False, vertices = grafo.obtener_vertices())

    while not q.esta_vacia():
        (v, w), peso = q.desencolar()
        if w in visitados:
            continue

        arbol.agregar_arista(v, w, peso)
        visitados.add(w)

        for x in grafo.adyacentes(w):
            if x not in visitados:
                q. encolar(((w, x), grafo.peso_arista(w, x)))

    return arbol