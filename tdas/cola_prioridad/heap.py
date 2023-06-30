import heapq

PANIC_COLA_VACIA = "La cola está vacía"

class Heap:
    def __init__(self):
        self.heap = []

    def esta_vacia(self):
        return len(self.heap) == 0

    """Heap de máximos (para heap de mínimos pasar self.contador en vez de -self.contador)"""
    def encolar(self, elemento):
        heapq.heappush(-self.heap, elemento)

    def ver_max(self):
        if self.esta_vacia():
            raise Exception(PANIC_COLA_VACIA)
        return self.heap[0]

    def desencolar(self):
        if self.esta_vacia():
            raise Exception(PANIC_COLA_VACIA)
        return heapq.heappop(self.heap)

    def cantidad(self):
        return len(self.heap)