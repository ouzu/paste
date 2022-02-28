import { dictionary, locale, getLocaleFromNavigator, _ } from 'svelte-i18n';

function setupI18n({ withLocale: _locale } = { withLocale: getLocaleFromNavigator() }) {
    dictionary.set({
        en: {
            app: {
                title: "paste",
                aboutLink: "about",
            },
            home: {
                files: "Files:",
                drag: "Drag file here or",
                select: "click to select",
                drop: "Drop to upload",
            },
            about: {
                text: "...",
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
                title: "paste",
                aboutLink: "Info",
            },
            home: {
                files: "Dateien:",
                drag: "Ziehe eine Datei her oder",
                select: "klicke um sie manuell auszuwählen",
                drop: "Lasse los zum hochladen",
            },
            about: {
                text: "...",
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