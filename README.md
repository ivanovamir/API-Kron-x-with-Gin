# API Kron-x

---

![Logo](https://github.com/ivanovamir/API-Kron-x-with-Gin/blob/main/python_backend/media/data/photos/categories/™аЃ≠•™б_дЃвЃ_3pv38Ez.png)


**Kron-x** - this is a commercial project that we developed with our team for a commercial site selling auto parts

## Technologies used in this project:
- Python (3.10.3)
  - Pillow (9.1.1)
  - gunicorn (20.1.0)
  - django-cors-headers (3.13.0)
  - psycopg2-binary (2.9.3)
  - django-admin-interface (0.20.0)
  - attr
  - django-object-actions (4.0.0)
  - requests (2.28.1)
  - crispy-bootstrap5 (0.3.1)
  - Django (4.0.5)
- Go (1.19.1)
  - Gin (1.18.1)
  - joho / godotenv
  - go-gorm / gorm
  - spf13 / viper
  - golang-jwt / jwt
  - and other ...
- JavaScript
- HTML
- CSS
- Nginx
- PostgreSQL (14.5)


## Project objectives:
+ Create a fast and highly loaded **`API`** for an online store selling car parts
+ Develop a pleasant design and competent **`UX / UI`**
+ Convenient admin panel system
+ Correct database architecture
+ Extensibility for future project updates
+ Authentication system using **`JWT`** tokens
+ Registration and authentication system, with a personal user account
+ Informing about the placed order by **`SMS`** and e-mail using **`sms-hosting`**


## Installation

**For security reasons, all files that can be compromised have been moved to **`.gitignore`** and used through the virtual environment. Therefore, you need to enter your own data and change the code a little.**

**At first you need to cteate database and register user in `PostgreSQL`:**
```sql
CREATE DATABASE DATABAS_NAME;
CREATE USER user_name WITH PASSWORD 'PASSWORD';
```
**Then you need connect `go` and `python` to this db:**

**In golang_backend/config/db.go:**
```go
db, err := gorm.Open(postgres.Open("postgres://user_name:password@host/port")
```
**In python_backend/main_settings/settings.py:**
```python
DATABASES = {
   'default': {
       'ENGINE': 'django.db.backends.postgresql',
       'NAME': 'db_name',
       'USER': 'user_name',
       'PASSWORD': 'password',
       'HOST': 'host',
       'PORT': 'port',
   }
}
```

**In `token.go` you need to define youre `signed key` for Jwt token generation**
```go
const (
    SignedKey = "YOURE_SIGNED_KEY"
)
```

**And replace this:** `os.Getenv("Signed_Key")` **To this:** `SignedKey`

**Also need to create your sms hosting account, open `api gateway`, and change http requests and create email and get `applications password` to send emails from this mail**

**Also you need to replace** `os.Getnev("Key_Name")` **from .env to yore data's in this files:**
**`golang_backend\controller\emails.go`** and **`golang_backend\controller\emailchecker.go`**- change **`Email`** and **`App password`**

**If you want to run a separate server for posting posts, as we did, then if necessary, you can go into the folder `static_app` and change the settings you need**

---

**To run the Go api run this:**
```go
go mod build
go run cmd/main.go
```
**or**
```go
go mod build
go guild cmd/main.go
```
**also you can use makefile commands**
```
make build
make run
```


**To run the Admin panel go to `cd python_backend` and run this commands:**
```python
pip install -r req.txt
python manage.py makemigrations
python manage.py migrate
python manage.py runserver
```
___

## To contact with me:

### [Site link](https://kron-x.ru/)
### [My Telegram](https://t.me/amirich18)