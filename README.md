# SecureEntry - Gestion des Entrées et Sorties avec RFID


## Fonctionnalités


## Structure du Projet
```
GrandOral/
          ├── .env                  # Fichier de configuration des variables d'environnement
          ├── .gitignore            # Fichiers et dossiers à ignorer par Git
          ├── app.py                # Fichier principal de l'application Flask
          ├── createadmin.py        # Script pour créer un utilisateur administrateur
          ├── requirements.txt      # Liste des dépendances Python
          ├── scan_rfid.py          # Script pour gérer la lecture des badges RFID
          ├── lecteur_rfid/         # Dossier contenant les scripts liés au lecteur RFID
          │   └── lecteur_rfid.py   # Script pour interagir avec le matériel RFID
          ├── templates/            # Dossier contenant les fichiers HTML pour l'interface utilisateur
          │   ├── index.html        # Page d'accueil
          │   ├── rechercher.html   # Page pour rechercher un étudiant
          │   ├── registre.html     # Page pour afficher le registre des entrées/sorties
          │   ├── badges_inconnus.html # Page pour gérer les badges inconnus
          │   ├── simulateur.html   # Page pour simuler les actions RFID
          │   ├── login.html        # Page de connexion
          │   └── menu.html         # Menu principal
          ├── instance/             # Dossier contenant les fichiers spécifiques à l'instance
          │   ├── database.db       # Base de données SQLite
          │   └── COMMANDE.txt      # Fichier de commandes ou notes spécifiques
          └── README.md             # Documentation du projet
```
## Prérequis

- Go

## Installation

1. Clonez le dépôt :
    ```bash
    git clone https://github.com/Statistiques00/
    cd```

## Utilisation

1. Lancez l'application
``go run /src/main.go``
2. Accédez au terminal de commande pour accéder aux fonctionnalités

## Contribution
Les contributions sont les bienvenues ! Veuillez soumettre une pull request ou ouvrir une issue pour signaler un problème.

## Licence

Ce projet est sous licence MIT. Consultez le fichier ``LICENSE`` pour plus d'informations.