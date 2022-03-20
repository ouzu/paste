import { dictionary, locale, getLocaleFromNavigator, _ } from 'svelte-i18n';

function setupI18n({ withLocale: _locale } = { withLocale: getLocaleFromNavigator() }) {
    dictionary.set({
        en: {
            app: {
                title: "paste",
                aboutLink: "about",
                clientLink: "client",
            },
            home: {
                files: "Files:",
                drag: "Drag file here or",
                select: "click to select",
                drop: "Drop to upload",
            },
            about: {
                text: "Work in progress...",
            },
            client: {
                heading: "client",
                summary: "The commandline client can be used to store and retreive files.",
                usage: "basic usage",
                upload: "To upload a file to this server use:",
                download: "To download a file from this server use:",
                help: "To checkout more options use:",
                releases: "releases",
            },
            file: {
                save: "save",
                pdf: {
                    warning: "It appears your Web browser is not configured to display PDF files.",
                },
                zip: {
                    contents: "archive contents:",
                },
            },
            notfound: {
                title: "Not found",
                message: "This page does not exist.",
            },
        },
        de: {
            app: {
                title: "Paste",
                aboutLink: "Info",
                clientLink: "Client",
            },
            home: {
                files: "Dateien:",
                drag: "Ziehe eine Datei her oder",
                select: "klicke um sie manuell auszuwählen",
                drop: "Lasse los zum hochladen",
            },
            about: {
                text: "Diese Seite ist noch in Arbeit...",
            },
            client: {
                heading: "Client",
                summary: "Der Kommandozeilen-Client kann benutzt werden, um Dateien hoch- und runterzuladen.",
                usage: "Grundlegende Nutzung",
                upload: "Um eine Datei hochzuladen:",
                download: "Um eine Datei herunterzuladen:",
                help: "Um weitere Optionen anzuschauen:",
                releases: "Releases",
            },
            file: {
                save: "speichern",
                pdf: {
                    warning: "Dein Browser kann keine PDF Dateien darstellen.",
                },
                zip: {
                    contents: "Archivinhalte:",
                },
            },
            notfound: {
                title: "Nicht gefunden",
                message: "Diese Seite existiert nicht.",
            },
        },
    });

    locale.set(_locale);
}

export { _, setupI18n };