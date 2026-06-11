# Gator

Gator es una CLI escrita en Go para registrar usuarios, guardar feeds RSS, seguir feeds y consultar posts agregados desde una base de datos PostgreSQL.

## Requisitos

Para ejecutar este programa necesitas tener instalado:

- Go
- PostgreSQL

Tambien necesitas tener una base de datos PostgreSQL creada para Gator.

## Instalacion

Instala la CLI con `go install`:

```bash
go install github.com/GianImpedovo/aggregator@latest
```

Si Go no encuentra el comando `gator` despues de instalarlo, asegurate de que el directorio de binarios de Go este en tu `PATH`:

```bash
go env GOPATH
```

Normalmente el binario queda en:

```bash
$GOPATH/bin
```

## Configuracion

Crea un archivo llamado `.gatorconfig.json` en tu directorio home.

Ejemplo:

```json
{
  "db_url": "postgres://usuario:password@localhost:5432/aggregator?sslmode=disable",
  "current_user_name": ""
}
```

Cambia `usuario`, `password`, `localhost`, `5432` y `aggregator` segun tu configuracion local de PostgreSQL.

Antes de usar la aplicacion, crea las tablas ejecutando los archivos SQL de `sql/schema` en orden:

```bash
psql "postgres://usuario:password@localhost:5432/aggregator?sslmode=disable" -f sql/schema/001_users.sql
psql "postgres://usuario:password@localhost:5432/aggregator?sslmode=disable" -f sql/schema/002_feeds.sql
psql "postgres://usuario:password@localhost:5432/aggregator?sslmode=disable" -f sql/schema/003_feed_follows.sql
psql "postgres://usuario:password@localhost:5432/aggregator?sslmode=disable" -f sql/schema/004_feed_lastfetched.sql
psql "postgres://usuario:password@localhost:5432/aggregator?sslmode=disable" -f sql/schema/005_posts.sql
```

## Uso

Registra un usuario:

```bash
gator register gian
```

Inicia sesion con un usuario existente:

```bash
gator login gian
```

Lista usuarios:

```bash
gator users
```

Agrega un feed y siguelo automaticamente:

```bash
gator addfeed "Boot.dev Blog" "https://blog.boot.dev/index.xml"
```

Lista todos los feeds:

```bash
gator feeds
```

Sigue un feed existente por URL:

```bash
gator follow "https://blog.boot.dev/index.xml"
```

Lista los feeds que sigue el usuario actual:

```bash
gator following
```

Busca posts nuevos periodicamente:

```bash
gator agg 1m
```

Muestra posts del usuario actual:

```bash
gator browse 10
```

Deja de seguir un feed:

```bash
gator unfollow "https://blog.boot.dev/index.xml"
```

## Repositorio

El proyecto esta disponible en GitHub:

https://github.com/GianImpedovo/Bootdev-aggregator
