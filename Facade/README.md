Obtener un api simple y entendible

Buffer: Representa un buffer de texto con un ancho y altura especificados. Almacena una matriz de rune que representa el contenido del buffer.

NewBuffer: Función constructora para crear un nuevo buffer con el ancho y altura especificados.

At: Método en el tipo Buffer que devuelve el rune en la posición dada en el buffer.

ViewPort: Representa una ventana de visualización en un buffer. Tiene un offset para desplazarse por el buffer.

GetCharacterAt: Método en el tipo ViewPort que devuelve el rune en la posición dada en el buffer asociado, considerando el desplazamiento.

Console: Representa una consola virtual que contiene buffers y viewports. Tiene una lista de buffers y viewports, así como un offset global.

NewConsole: Función constructora para crear una nueva consola virtual con un buffer inicial y un viewport asociado.