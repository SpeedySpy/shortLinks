# projet golang 

**Omma Habiba BIPLOB**
**Faiza AKABLI**
**Biraveen SIVAHARAN**

**Installation**

go init
initiliser le projet: go mod init : go mod init api
go mod vendor
compiler le fichier main.go : go run main.go
fichier de test a tester : go run main_test.go

**packages**

context: Pour gérer les annulations de requête.
encoding/json: Pour encoder et décoder des données JSON.
log: Pour le journalisation.
net/http: Pour gérer les requêtes HTTP.
time: Pour travailler avec le temps.
github.com/google/uuid: Pour générer des identifiants uniques.
github.com/gorilla/mux: Pour le routage HTTP.
go.mongodb.org/mongo-driver/bson: Pour la manipulation de documents BSON.
go.mongodb.org/mongo-driver/mongo: Pour l'interaction avec MongoDB.
go.mongodb.org/mongo-driver/mongo/options: Pour spécifier les options lors de l'interaction avec MongoDB.

**structure URL**

Cette structure représente un document dans la collection MongoDB. Elle contient un identifiant unique (ID), une URL longue (LongUrl), une URL courte (ShortUrl), et une date d'expiration (ExpirationAt)

**Constantes**

mongoURI: L'URI de connexion à MongoDB.
dbName: Le nom de la base de données MongoDB.
collectionName: Le nom de la collection MongoDB.

**Variables globales**

client: Le client MongoDB.
collection: La collection MongoDB.

**Fonction init**

Initialise la connexion à MongoDB et affecte le client et la collection

**Fonction main**

Initialise un routeur Gorilla Mux pour gérer les requêtes HTTP.
Définit deux points de terminaison pour les routes /shorten et /get-long-url.
Lance le serveur HTTP sur le port 8080

**Fonction shortenUrl**

Extrait les données JSON de la requête HTTP pour obtenir l'URL longue.
Génère un identifiant unique pour l'URL courte.
Crée une URL courte avec un lien préfixé.
Insère l'URL dans la collection MongoDB avec une date d'expiration.
Renvoie l'URL courte en tant que réponse JSON.

**Fonction redirectToLongURL**

Extrait les données JSON de la requête HTTP pour obtenir l'URL courte.
Recherche l'URL correspondante dans la collection MongoDB.
Si l'URL est trouvée et n'a pas expiré, renvoie l'URL longue en tant que réponse JSON.
Sinon, renvoie une erreur correspondante.


