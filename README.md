# Prueba tecnica Zapping

## Utilizacion

Para utilizar la aplicacion se requiere docker 

Para crear las imagenes de docker se utiliza el comando
```
> docker-compose build
```
para activar el containe
```
> docker-compose up
```

## Aplicación

La aplicacion principal se aloja en localhost:8080

## Desarrollo de la aplicación

Para el desarrollo se asumio que:

- El microservicio solamente entrega el archivo con formato m3u8. Los archivos .ts estan siempre disponibles.

- Solo se requiere de tres vistas, login, registro y reproductor.

- En el enunciado se hablaba de microservicio de NodeJs, este se asumio que se referia a Golang.


En este caso se desarrollo lo que se entendio que se requeria de la aplicacion sin trabajar sobre los detalles de esta para evitar tiempo extra desarrollo. Ejemplo de estos detalles:

- Un front mas producido y responsivo.
- Mas restriccion al momento de crear una cuenta (limite de caracteres, solo caracteres y numeros, etc).
- Una view intermedia entre el login y el livestreaming.
- Mayor detalle de erorres para el usuario.
- Utilizar imagnes, un header y un footer.
- Temas de seguridad de datos.

La mayor parte del tiempo de desarrollo se dedico al estudio de crear aplicaciones web tanto en javascript como en Golang. Asi como montar la aplicacion utilizando docker

