PANIC_PILA_VACIA = "La pila está vacía"

class Pila:
    def __init__(self):
        self.pila = []

    def esta_vacia(self):
        return len(self.pila) == 0

    def ver_tope(self):
        if self.esta_vacia():
            raise Exception(PANIC_PILA_VACIA)
        return self.pila[-1]

    def apilar(self, elemento):
        self.pila.append(elemento)

    def desapilar(self):
        if self.esta_vacia():
            raise Exception(PANIC_PILA_VACIA)
        return self.pila.pop()

