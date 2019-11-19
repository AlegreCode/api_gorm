# api_gorm
### Example api of user, posts and comments with gorm framework.

##### Packages utilizados:
- **godotenv**: permite cargar variables de entorno desde un archivo .env. [Aquí](https://github.com/joho/godotenv)
- **gorm**: herramienta para acceso a datos sin escribir sentencias SQL. [Aquí](http://gorm.io/)

##### Config base de datos
El archivo .env-example renombrar a .env y completar las variables con tus credenciales de conexión.

##### Instalación de paquetes
Para evitar conflictos de versiones utilizamos el gestor de dependencias de go [dep](https://golang.github.io/dep/). Debes tener instalado esta herramienta (para ver como clic [aquí](https://golang.github.io/dep/docs/installation.html)), luego entrar a la raíz de tu proyecto y ejecutar el comando `dep ensure`. Una vez instalados todos los paquetes ya correr el proyecto.
