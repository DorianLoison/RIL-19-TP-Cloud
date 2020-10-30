# RIL-19-TP-Cloud
Code Golang pour TP déplouement Cloud - Dorian Loison & Nicolas Kamphaus

# Etapes

## 1.Téléchargement du code
Réussite
## 2.Ouverture du code dans votre IDE
Réussite
## 3.Rédaction des tests unitaires : https://golang.org/pkg/testing/
Fait
## 4.Test unitaire pour les routes POST/GET/PUT/DELETE
Fait
## 5.Mise en place d’une solution MongoDB (la moins couteuse et la plus simple en gestion)
Fait => MongoDb Atlas
## 6.Mise en place d’un système CI/CD
Fait => Travis CI
## 7.Automatisation des exécutions des tests
Fait
## 8.Création d ‘un dépôt de code
Fait sur github (ici même)
## 9.Plug de la stack CI/CD etG Github
Fait
## 10.Création d’un compte client dans un cloud provider
Fait => AWS
## 11.Déploiement d’un espace (VM free-tier type linux)
Fait => LightSail
## 12.Configuration des accès à l’espace entre l’outil de CI/CD et la VM
En cours
## 13.Si test ok déploiement dans la VM
A faire
## 14.Réaliser plusieurs pull request en créant des nouvelles routes
A faire

# Critères d’acceptation

## 1.Jusification du CI/CD?
Travis CI, outil de CI/CD gratuit pour les projets open source.

Créé pour s'intégrer à GitHub, que nous utilisons.

## 2.Coût mensuelle < 100 $?
DB => MongoDb Atlas, cluster type M0 (0 USD/mois)

Application en go => VM LightSail AWS (3,50 USD/mois)

CI/CD => Travis CI (projet open source 0 USD/mois)

Total : 3,50 USD/mois

## 3.Dispo Europe/Asie?
Actuellement non. Juste une VM dans un center en France.

Pour augmenter la disponibilité et reduire la latence, ajouter des VM dans les différents secteurs.

Utiliser notre outil de CI/CD pour diffuser les maj partout.

La solution MongoDb retenue dispose d'une accessibilité mondiale.

## 4.Solution mongoDB simple, peu cher et flexible?
MongoDb Atlas a été retenue.

Simplicité de connexion cloud + diffusion data a traver le monde grace a plus de 50 cluster dans le monde.

De base, elle dispose d'un offre free tier pour nos besoin.

Autrement, on paye à l'utilisation, ce qui fait que ça revient à moins cher en faible activité.

## 5.Limitations application et infrastructure cible?
Infra pratiquement gratuite avec faibles resources

## 6.Votre local (PC) est un cloud privé (?)