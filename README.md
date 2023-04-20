# GROUPIE-TRACKER

- [GROUPIE-TRACKER](#groupie-tracker)
  - [Présentation du projet](#présentation-du-projet)
  - [Installation et utilisation](#installation-et-utilisation)
  - [Fonctionnalités](#fonctionnalités)

## Présentation du projet

Groupie-tracker est un site pour les fans de musique, permettant de voir toutes les informations sur un groupe de musique comme l'année de création ou même les artistes dont le groupe est composé et tout ça de façon très simple et ludique.

## Installation et utilisation

L'installation est simple, il suffit de cloner le répertoire git en faisant la commande :
```bash
https://github.com/Archilive/Groupie-Tracker.git
```
Et d'ensuite aller dans le fichier `GROUPIE-TRACKER` et faire :
```bash
go run serveur/serveur.go
```
Une fois ces commandes faites, vous pourrez immédiatement ouvrir votre navigateur et accéder au site en tapant [`http://localhost:8080`](http://localhost:8080). Vous allez arriver sur la page d'accueil ou vous trouverez tous les artistes affichés avec quelques informations, sur la droite, vous trouverez les filtres, la barre de recherche et la façon dont vous voulez que ce soit trié. Vous pouvez à tout moment cliquer sur le groupe qui vous intéresse pour avoir plus d'information.

## Fonctionnalités

- Pouvoir rechercher un groupe avec la barre de recherche
- Pouvoir trier les groupes dans un ordre précis que l'on définit
- Pouvoir filtrer les groupes avec le nombre d'artistes que l'on cherche dans le groupe
- Pouvoir filtrer les groupes selon leur date de création en choisissant le minimum comme le maximum
- Pouvoir filtrer les groupes selon leur date de premier album en choisissant le minimum comme le maximum