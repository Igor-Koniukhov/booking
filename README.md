# Bookings and Reservations
##### To run the project:

1. Ensure you are in the project's root directory.
2. Copy the environment variables:

```shell
cp .env.example .env
```
3. For migrations install CLI Soda (Pop) - [docs instalation](https://gobuffalo.io/documentation/database/soda/), run:

```shell
cp database.yml.example database.yml
```
- then write down your credentials in database.yml, run:

```bash
soda migrate
```
- for creation new migration file, run:

```bash
soda generate fizz [createYorTableNameTable]
```
- for deletion, run:
```bash
soda migrate down
```
4. To start the project ensure you are in the project's root, run:
```bash
./run.sh
```
This is the repository for my bookings and reservations project

- Built in Go version 1.16
- Uses the [chi router](https://github.com/go-chi/chi)
- Uses [alex edwards scs session management](https://github.com/alexedwards/scs)
- Uses [nosurf](https://github.com/justinas/nosurf)
- Uses [govalidator](https://github.com/asaskevich/govalidator)
# DB
- Uses [gobuffalo](https://github.com/gobuffalo/pop/)

     Buffalo [documentation](https://gobuffalo.io/en/docs/overview)
     For postgres [github.com/jacks/pgconn](https://github.com/jackc/pgconn)
     

# JS
- Uses [Vanilla JS Datepicker](https://mymth.github.io/vanillajs-datepicker/#/)

### Useful resources
- Parsing and [formating date/time in Go](https://www.pauladamsmith.com/blog/2011/05/go_time.html)