def es_conexo(grafo):
    if not grafo.obtener_vertices():
        return True

    visitados = set()
    origen = grafo.vertice_aleatorio()
    dfs(grafo, origen, visitados)

    return len(visitados) == len(grafo.obtener_vertices())


def dfs(grafo, v, visitados):
    visitados.add(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs(grafo, w, visitados)