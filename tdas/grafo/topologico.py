from ejemplos import correlativas
from pila import Pila
from cola import Cola


def grados_entrada(grafo):
    g_ent = {}
    for v in grafo:
        g_ent[v] = 0
    for v in grafo:
        for w in grafo.adyacentes(v):
            g_ent[w] += 1
    return g_ent
    # O(V + E)


def topologico_grados(grafo):
    g_ent = grados_entrada(grafo)
    q = Cola()
    for v in grafo:
        if g_ent[v] == 0:
            q.encolar(v)
    resultado = []
    while not q.esta_vacia():
        v = q.desencolar()
        resultado.append(v)
        for w in grafo.adyacentes(v):
            g_ent[w] -= 1
            if g_ent[w] == 0:
                q.encolar(w)
    return resultado


def _dfs(grafo, v, visitados, pila):
    visitados.add(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            _dfs(grafo, w, visitados, pila)
    pila.apilar(v)


def topologico_dfs(grafo):
    visitados = set()
    pila = Pila()
    for v in grafo:
        if v not in visitados:
            _dfs(grafo, v, visitados, pila)
    return pila_a_lista(pila)


def pila_a_lista(pila):
    lista = []
    while not pila.esta_vacia():
        lista.append(pila.desapilar())
    return lista


if __name__ == "__main__":
    print(topologico_dfs(correlativas))