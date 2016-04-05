--
-- Fichier généré par SQLiteStudio v3.0.7sur mar. mars 29 10:39:32 2016
--
-- Encodage texte utilisé: windows-1252
--
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- Table: Usagers
CREATE TABLE Usagers (id INT (11) NOT NULL, nom VARCHAR (35) NOT NULL, prenom VARCHAR (35) NOT NULL, email VARCHAR (70) NOT NULL, Auteur_fiche VARCHAR (35), Date VARCHAR (15) NOT NULL, NomTCCL VARCHAR (255), telephone_fixe VARCHAR (10) NOT NULL, Portable VARCHAR (10), CirconstancesTrauma VARCHAR (1024) NOT NULL, ConditionsPriseEnChargeMedicale VARCHAR (1024) NOT NULL, ConditionsPriseEnChargeSociale VARCHAR (1024), DemarchesAdministrativesJuridiquesEffectuees VARCHAR (1024), ProblemesRencontres VARCHAR (4096) NOT NULL, AttentesAFTC VARCHAR (1024) NOT NULL, Observations VARCHAR (1024) NOT NULL, PRIMARY KEY (id));

COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
