def es_bipartito(grafo):
    if not grafo:
        return True

    colores = {}
    origen = grafo.vertice_aleatorio()

    if not dfs(grafo, origen, 0, colores):
        return False

    return True


def dfs(grafo, v, color, colores):
    colores[v] = color

    for w in grafo.adyacentes(v):
        if w not in colores:
            if not dfs(grafo, w, 1 - color, colores):
                return False
        elif colores[w] == color:
            return False

    return True