def obtener_orden(grafo):
    orden = []
    visitados = set()

    for v in grafo:
        if v not in visitados:
            dfs(grafo, v, visitados, orden)

    return invertir_lista(orden)


def dfs(grafo, v, visitados, orden):
    visitados.add(v)
    
    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs(grafo, w, visitados, orden)

    orden.append(v)


def invertir_lista(lista):
    return list(reversed(lista))