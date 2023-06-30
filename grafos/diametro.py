from tdas.cola.cola import Cola

def diametro(grafo):
    if not grafo:
        return 0

    max_distancia = 0

    for v in grafo:
        distancias = bfs(grafo, v)
        max_distancia = max(max_distancia, max(distancias.values()))

    return max_distancia


def bfs(grafo, origen):
    distancias = {origen: 0}
    cola = Cola()
    cola.encolar(origen)

    while not cola.esta_vacia():
        v = cola.desencolar()

        for w in grafo.adyacentes(v):
            if w not in distancias:
                distancias[w] = distancias[v] + 1
                cola.encolar(w)

    return distancias