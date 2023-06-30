from collections import deque

PANIC_COLA_VACIA = "La cola esta vacia"

class Cola:
    def __init__(self):
        self.cola = deque()

    def esta_vacia(self):
        return len(self.cola) == 0

    def ver_primero(self):
        if self.esta_vacia():
            raise Exception(PANIC_COLA_VACIA)
        return self.cola[0]

    def encolar(self, dato):
        self.cola.append(dato)

    def desencolar(self):
        if self.esta_vacia():
            raise Exception(PANIC_COLA_VACIA)
        return self.cola.popleft()