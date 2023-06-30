def grados(grafo):
    # devolver un diccionario string -> int
    grados = {}
    for v in grafo:
        grados[v] = len(grafo.adyacentes(v))
    return grados

def grados_entrada(grafo):
    # devolver un diccionario string -> int
    entrada = {}
    for v in grafo:
        for w in grafo.adyacentes(v):
            entrada[w] = entrada.get(w, 0) + 1

    return entrada

def grados_salida(grafo):
    # devolver un diccionario string -> int
    salida = {}
    for v in grafo:
        salida[v] = len(grafo.adyacentes(v))
    return salida