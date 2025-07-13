# Google Takeout Localization

## Table of Contents
- [The Problem with Google Takeout Files](#the-problem-with-google-takeout-files)
- [Project Aims](#this-project-aims)
- [Repository Structure](#repository-structure)
- [How to Contribute](#how-to-contribute)
- [Steps to Add a Translation](#steps-to-add-a-translation)
- [Example](#example)
- [Future Work](#future-work)

## The Problem with Google Takeout Files

Google Takeout is a service that allows users to export their data from various Google services. However, the exported structure depends partially on the user account's language settings. Some files or directories are named differently depending on the language, which can make it difficult to manage and process the data consistently.

## This Project Aims

The idea is to collect the different localizations used by Google and present them in a machine-readable format. This way, tool developers can use this data to create scripts or applications that can handle Google Takeout files regardless of the user's language settings.

## Repository Structure

The folder `localizations` contains the localization files in JSON format. Each file represents a Google service and contains the different localizations for the files and directories used by that service. Entries are hierarchical, with nested structures as follows:

1. **Directories**:
   - Represented by the `kind: "directory"` key.
   - Can contain subdirectories listed under the `directories` key, and files listed under the `files` key.

2. **Files**:
   - Represented by the `kind: "file"` key.
   - Can contain columns, which are listed under the `columns` key.

3. **Columns**:
   - Contain specific translation mappings for file headers.

### Example of Nested Structure

```json
{
    "Service Name": {
        "kind": "directory",
        "localizations": {
            "en": "Service Name",
            "fr": "Nom du service"
        },
        "directories": {
            "Subdirectory": {
                "kind": "directory",
                "localizations": {
                    "en": "Subdirectory",
                    "fr": "Nom du sous-répertoire"
                },
                "files": {
                    "file.csv": {
                        "kind": "file",
                        "localizations": {
                            "en": "file.csv",
                            "fr": "fichier.csv"
                        },
                        "columns": {
                            "Column Header": {
                                "en": "Column Header",
                                "fr": "En-tête de colonne"
                            }
                        }
                    }
                }
            }
        }
    }
}
```

## How to Contribute

**Note:** Provide translations based on the exact file names delivered by the Takeout service.

### Adding a New Product

Create a new file in the `localizations` directory with the name of the product (e.g., `ProductName.json`). Use the following format:

```json
{
    "Product Name": {
        "translations": {
            "kind": "directory",
            "fr": "Nom du produit",
            "es": "Nombre del producto",
            "ru": "Название продукта"
        },
        "directories": {},
        "files": {}
    }
}
```

### Adding a localization for an Existing Product 

Locate the element in the JSON file and add a new key under the `localization` object for the desired language. Use the following format:

```json
    "localizations": {
        "fr": "fichier.csv",
        "es": "archivo.csv",
        "de": "datei.csv"
    }
```


## Guidelines

- fork the repository and create a new branch for your changes.
- Follow the existing structure for directories and files.
- Always include the `kind` key to specify whether the entry is a directory or a file.
- Use the `columns` object to define translations for file headers.
- Ensure all translations are based on the actual files delivered by the Takeout service.
- submit a pull request with a clear description of your changes.


## Future Work

- Add a license.
- Expand the localization files to include more Google services.
- Provide packages for different programming languages to make it easier to use the localization files.





