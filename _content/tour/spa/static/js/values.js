/* Copyright 2012 The Go Authors.   All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
'use strict';

angular.module('tour.values', []).

// List of modules with description and lessons in it.
value('tableOfContents', [{
    'id': 'welcome',
    'title': '',
    'description': `<p>Este material es para cualquier desarrollador de nivel intermedio que tenga algo de experiencia con otros lenguajes de programación y quiera aprender Go. Creemos que este manterial es perfecto para cualquiera que quiera comenzar a aprender Go o quiera una comprensión más profunda del lenguaje y sus aspectos internos.</p>
    <p><b>La primera vez que ejecutes un ejemplo puede ser que lleve su tiempo para ser completado. Ésto es porque el servicio del tour puede no estar ejecutándose. Por favor dale algo de tiempo para que termine.</b></p>
    `,
    'lessons': [
        'welcome',
    ]
}, {
    'id': 'language-mechanics',
    'title': 'Mecanismos del Lenguaje',
    'description': `
        <p>Este material cubre toda la sintáxis, idiomas, implementación y espeficación del lenguaje. Una vez que acabes con este material entenderás tanto los mecanismos del lenguaje como las simpatías mecánicas que tiene para el hardware y el sistema operativo.</p>
        `,
    'lessons': [
        'variables',
        'struct-types', 
        'pointers', 
        'constants',
        'functions',
        'arrays',
        'slices',
        'maps',
        'methods',
        'interfaces',
        'embedding',
        'exporting',
    ]
}, {
    'id': 'composition-interfaces',
    'title': 'Composición e Interfaces',
    'description': `
        <p>Este material cubre las cosas prácticas que necesitas saber sobre composición e interfaces.</p>
        `,
    'lessons': [
        'composition-grouping',
        'composition-assertions',
        'composition-pollution',
        'composition-mocking',
        'composition-decoupling',
        'error-handling',
    ]
},{
    'id': 'concurrency',
    'title': 'Concurrencia',
    'description': `
        <p>Este material cubre todos los aspectos sobre concurrencia del lenguaje. Una vez hayas terminado el material entenderás los mecanismos de concurrencia del lenguaje y las simpatías mecánicas relacionadas con la concurrencia que tiene Go para tanto el hardware como para el sistema operativo.</p>
       `,
    'lessons': [
        'goroutines',
        'data_race',
        'context',
        'channels',
    ]
},{
    'id': 'generics',
    'title': 'Generics',
    'description': `
        <p>Este material cubre todos los aspectos sobre generics del lenguaje. Generics consiste en proporcionar la habilidad de escribir funciones concretas polimórficas con el soporte de listas de parámetros de tipo.</p>
       `,
    'lessons': [
        'generics-basics',
        'generics-underlying-types',
        'generics-struct-types',
        'generics-behavior-constraints',
        'generics-type-constraints',
        'generics-multi-type-params',
        'generics-slice-constraints',
        'generics-channels',
        'generics-hash-table',
    ]
},{
    'id': 'algorithms',
    'title': 'Algoritmos',
    'description': `
        <p>Este material proporciona código Go implementando un conjunto común y divertido de algoritmos.</p>
       `,
    'lessons': [
        'algorithms-bits-seven',
        'algorithms-strings',
        'algorithms-numbers',
        'algorithms-slices',
        'algorithms-sorting',
        'algorithms-data',
        'algorithms-searches',
        'algorithms-fun',
    ]
}]).
// translation
value('translation', {
    'off': 'off',
    'on': 'on',
    'syntax': 'Resaltado de Sintáxis',
    'lineno': 'Números de Línea',
    'reset': 'Reiniciar Diapositiva',
    'format': 'Formatear Código Fuente',
    'kill': 'Terminar Programa',
    'run': 'Ejecutar',
    'compile': 'Compilar y Ejecutar',
    'more': 'Opciones',
    'toc': 'Tabla de Contenidos',
    'prev': 'Anterior',
    'next': 'Siguiente',
    'waiting': 'Esperando al servidor remoto...',
    'errcomm': 'Error comunicándose con el servidor remoto.',
    'submit-feedback': 'Enviar feedback sobre esta página',
    'search': 'Buscar contenido',

    // GitHub issue template: update repo and messaging when translating.
    'github-repo': 'github.com/ardanlabs/gotour',
    'issue-title': 'tour: [REPLACE WITH SHORT DESCRIPTION]',
    'issue-message': 'Change the title above to describe your issue and add your feedback here, including code if necessary',
    'context': 'Contexto',
}).

// Config for codemirror plugin
value('ui.config', {
    codemirror: {
        mode: 'text/x-go',
        matchBrackets: true,
        lineNumbers: true,
        autofocus: true,
        indentWithTabs: true,
        indentUnit: 4,
        tabSize: 4,
        lineWrapping: true,
        extraKeys: {
            'Shift-Enter': function() {
                $('#run').click();
            },
            'Ctrl-Enter': function() {
                $('#format').click();
            },
            'PageDown': function() {
                return false;
            },
            'PageUp': function() {
                return false;
            },
        },
        // TODO: is there a better way to do this?
        // AngularJS values can't depend on factories.
        onChange: function() {
            if (window.codeChanged !== null) window.codeChanged();
        }
    }
}).

// mapping from the old paths (#42) to the new organization.
// The values have been generated with the map.sh script in the tools directory.
value('mapping', {
    '#1': '/variables/1', // Variables
});
