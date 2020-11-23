export const Translations = {
  en: {},
  'en-US': {},
  'de-DE': {
      form: {
          task: {
              task: {
                  wraps: {
                      header: {text: 'Wraps'},
                      secondary: {text: 'Test der wraps'},
                  },
              },
          },
      },
      task: {
          Task: {
              display_name: {
                  label: 'Aufgabe',
                  hint: 'Kurz und prägnant, so dass es andere auch verstehen können.',
              },
              note: {
                  label: 'Notizen',
                  hint: 'Hier darfst und sollst du ausführlicher sein.',
              },
              description: {
                  label: 'Beschreibung',
              },
              id: {
                  label: 'ID',
              },
          },
      },
      reload: 'Neu laden',
      refresh: 'Aktualiseren',
      create: 'Ertellen',
      cancel: 'Abbrechen',
      delete: 'Löschen',
      search: 'Suchen',
      lorem:
          'This text translates lorem to: At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. ',
      install: 'Installieren',
      spec: {
          typeproto: {
              form: {
                  header: {
                      text: '__proto',
                  },
                  secondary: {
                      text: '',
                  },
              },
          },
          fieldproto: {
              form: {
                  header: {text: '__proto'},
                  secondary: {
                      text:
                          'The smallest field number you can specify is 1, and the largest is 229 - 1, or 536,870,911. You also cannot use the numbers 19000 through 19999 (FieldDescriptor::kFirstReservedNumber through FieldDescriptor::kLastReservedNumber), as they are reserved for the Protocol Buffers implementation - the protocol buffer compiler will complain if you use one of these reserved numbers in your .proto. Similarly, you cannot use any previously reserved field numbers. ',
                  },
              },
          },
          field: {
              map: {
                  headertext: 'Fields',
              },
          },
      },
      furo: {
          fieldconstraint: {
              map: {
          headertext: 'Constraints',
        },
      },
    },
  },
  de: {
    save: 'Speichern',
    reload: 'Neu laden',
    cancel: 'Abbrechen',
    delete: 'Löschen',
    install: 'Installieren',
  },
  'it-CH': {},
  'fr-CH': {},
};

// wire the defaults
Translations['de-CH'] = Translations['de-DE'];
