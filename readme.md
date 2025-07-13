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
   - Can contain subdirectories listed under the `subdirectories` key, and files listed under the `files` key.

2. **Files**:
   - Represented by the `kind: "file"` key.
   - Can contain columns, which are listed under the `columns` key.

3. **Columns**:
   - Contain specific translation mappings for file headers.

### Example of Nested Structure

```json
{
    "Service Name": {
        "translations": {
            "kind": "directory",
            "fr": "Nom du service"
        },
        "subdirectories": {
            "subdirectory": {
                "translations": {
                    "kind": "directory",
                    "fr": "Nom du sous-répertoire"
                },
                "files": {
                    "file.csv": {
                        "translations": {
                            "kind": "file",
                            "fr": "nom_fichier.csv",
                            "columns": {
                                "Column Header": {
                                    "fr": "En-tête de colonne"
                                }
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

The `.json` file uses the following structure:

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
        "subdirectories": {},
        "files": {}
    }
}
```

### Adding a New Directory in an Existing Product

To add a new directory, include a new key under the appropriate parent directory's `subdirectories` object. Use the following format:

```json
"directory_name": {
    "translations": {
        "kind": "directory",
        "fr": "nom du répertoire",
        "es": "nombre del directorio",
        "ru": "название каталога"
    },
    "subdirectories": {},
    "files": {}
}
```

### Adding a New File in an Existing Directory

To add a new file, include a new key under the appropriate directory's `subdirectories` object. Use the following format:

```json
"file_name.csv": {
    "translations": {
        "kind": "file",
        "fr": "nom_fichier.csv",
        "es": "nombre_archivo.csv",
        "ru": "имя_файла.csv",
        "columns": {
            "title": {
                "fr": "titre",
                "es": "título",
                "ru": "заголовок"
            },
            "visibility": {
                "fr": "visibilité",
                "es": "visibilidad",
                "ru": "видимость"
            }
        }
    }
}
```

## Guidelines

- Always include the `kind` key to specify whether the entry is a directory or a file.
- Use the `columns` object to define translations for file headers.
- Ensure all translations are based on the actual files delivered by the Takeout service.
- Follow the existing structure to maintain readability and organization.

## Example

### Adding the `playlists.csv` File in the `playlists` Directory Under the `YouTube and YouTube Music` Directory

According to the actual Takeout files, the directory and the file name aren't translated, but the columns are. The structure is as follows:

```json
{
    "YouTube and YouTube Music": {
        "translations": {
            "kind": "directory",
            "fr": "YouTube et YouTube Music",
            "es": "YouTube y YouTube Music",
            "ru": "YouTube и YouTube Music"
        },
        "subdirectories": {
            "playlists": {
                "translations": {
                    "kind": "directory",
                    "fr": "playlists"
                },
                "files": {
                    "playlists.csv": {
                        "translations": {
                            "kind": "file",
                            "columns": {
                                "Playlist ID": {
                                    "fr": "ID de la playlist",
                                    "es": "ID de la lista de reproducción",
                                    "ru": "ID плейлиста"
                                },
                                "Playlist Title (Original)": {
                                    "fr": "Titre (d'origine) de la playlist",
                                    "es": "Título (original) de la lista de reproducción",
                                    "ru": "Название (оригинал) плейлиста"
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
```

## Future Work

- Add a license.
- Expand the localization files to include more Google services.
- Provide packages for different programming languages to make it easier to use the localization files.





